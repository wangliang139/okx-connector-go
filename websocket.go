package okx_connector

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

var (
	// WebsocketTimeout is an interval for sending ping/pong messages if WebsocketKeepalive is enabled
	WebsocketTimeout = time.Second * 60
	// WebsocketKeepalive enables sending ping/pong messages to check the connection stability
	WebsocketKeepalive = false
)

// WsHandler handle raw websocket message
type WsHandler func(message []byte)

// ErrHandler handles errors
type ErrHandler func(err error)

// WsConfig webservice configuration
type WsConfig struct {
	Endpoint string
}

type SubscribeRequest struct {
	Op   string     `json:"op,omitempty"`
	Args []SubOpArg `json:"args,omitempty"`
}

type SubscribeResponse struct {
	Event  string          `json:"event,omitempty"`
	Code   *string         `json:"code,omitempty"`
	Msg    *string         `json:"msg,omitempty"`
	ConnId string          `json:"connId,omitempty"`
	Arg    json.RawMessage `json:"arg,omitempty"`
}

type SubOpArg struct {
	Channel *string `json:"channel,omitempty"`
	InstId  *string `json:"instId,omitempty"`
}

type WebsocketStreamClient struct {
	Endpoint string
	Debug    bool
	Logger   *log.Logger
	conn     *websocket.Conn
}

func (c *WebsocketStreamClient) debug(format string, v ...interface{}) {
	if c.Debug {
		c.Logger.Printf(format, v...)
	}
}

func (client *WebsocketStreamClient) dail() error {
	Dialer := websocket.Dialer{
		Proxy:             http.ProxyFromEnvironment,
		HandshakeTimeout:  45 * time.Second,
		EnableCompression: false,
	}
	headers := http.Header{}
	headers.Add("User-Agent", fmt.Sprintf("%s/%s", Name, Version))
	conn, _, err := Dialer.Dial(client.Endpoint, headers)
	if err != nil {
		return err
	}
	client.conn = conn
	client.conn.SetReadLimit(655350)
	return nil
}

func (client *WebsocketStreamClient) subscribe(channels []SubOpArg) error {
	if len(channels) == 0 {
		return fmt.Errorf("channels is empty")
	}
	op := &SubscribeRequest{
		Op:   "subscribe",
		Args: channels,
	}
	msg, err := json.Marshal(op)
	if err != nil {
		return err
	}
	err = client.conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		return err
	}
	_, message, err := client.conn.ReadMessage()
	if err != nil {
		return err
	}
	client.debug("Subscribe response: %s\n", message)
	var response SubscribeResponse
	if err := json.Unmarshal(message, &response); err != nil {
		return err
	}
	if response.Code != nil {
		return &ApiError{Code: *response.Code, Message: *response.Msg}
	}
	return nil
}

func (client *WebsocketStreamClient) serve(handler WsHandler, errHandler ErrHandler) (doneCh, stopCh chan struct{}, err error) {
	doneCh = make(chan struct{})
	stopCh = make(chan struct{})
	go func() {
		// This function will exit either on error from
		// websocket.Conn.ReadMessage or when the stopC channel is
		// closed by the client.
		defer close(doneCh)
		if WebsocketKeepalive {
			keepAlive(client.conn, WebsocketTimeout)
		}

		// Wait for the stopC channel to be closed.  We do that in a
		// separate goroutine because ReadMessage is a blocking
		// operation.
		silent := false
		go func() {
			select {
			case <-stopCh:
				silent = true
			case <-doneCh:
			}
		}()
		for {
			_, message, err := client.conn.ReadMessage()
			if err != nil {
				if !silent {
					errHandler(err)
				}
				return
			}
			handler(message)
		}
	}()
	return
}

func keepAlive(c *websocket.Conn, timeout time.Duration) {
	ticker := time.NewTicker(timeout)

	lastResponse := time.Now()
	c.SetPongHandler(func(msg string) error {
		lastResponse = time.Now()
		return nil
	})

	go func() {
		defer ticker.Stop()
		for {
			deadline := time.Now().Add(10 * time.Second)
			err := c.WriteControl(websocket.PingMessage, []byte{}, deadline)
			if err != nil {
				return
			}
			<-ticker.C
			if time.Since(lastResponse) > timeout {
				return
			}
		}
	}()
}

func NewWsPublicStreamClient(baseURL ...string) *WebsocketStreamClient {
	// Set default base URL to production WS URL
	url := "wss://ws.okx.com:8443"
	if len(baseURL) > 0 {
		for _, u := range baseURL {
			if len(u) > 0 {
				url = u
				break
			}
		}
	}
	return &WebsocketStreamClient{
		Endpoint: url + "/ws/v5/public",
		Logger:   log.New(os.Stderr, Name, log.LstdFlags),
	}
}

func NewWsPrivateStreamClient(baseURL ...string) *WebsocketStreamClient {
	// Set default base URL to production WS URL
	url := "wss://ws.okx.com:8443"
	if len(baseURL) > 0 {
		for _, u := range baseURL {
			if len(u) > 0 {
				url = u
				break
			}
		}
	}
	return &WebsocketStreamClient{
		Endpoint: url + "/ws/v5/private",
		Logger:   log.New(os.Stderr, Name, log.LstdFlags),
	}
}

func NewWsBusinessStreamClient(baseURL ...string) *WebsocketStreamClient {
	// Set default base URL to production WS URL
	url := "wss://ws.okx.com:8443"
	if len(baseURL) > 0 {
		for _, u := range baseURL {
			if len(u) > 0 {
				url = u
				break
			}
		}
	}
	return &WebsocketStreamClient{
		Endpoint: url + "/ws/v5/business",
		Logger:   log.New(os.Stderr, Name, log.LstdFlags),
	}
}
