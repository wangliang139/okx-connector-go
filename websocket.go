package okx_connector

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
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
	Endpoint        string
	Debug           bool
	Logger          *log.Logger
	Timeout         time.Duration // Timeout for ping/pong messages
	Keepalive       bool          // Enable ping/pong keepalive
}

type WebsocketStreamConn struct {
	Conn   *websocket.Conn
	Client *WebsocketStreamClient
}

func (c *WebsocketStreamClient) debug(format string, v ...interface{}) {
	if c.Debug {
		c.Logger.Printf(format, v...)
	}
}

func (c *WebsocketStreamClient) dail(ctx context.Context) (*WebsocketStreamConn, error) {
	Dialer := websocket.Dialer{
		Proxy:             http.ProxyFromEnvironment,
		HandshakeTimeout:  45 * time.Second,
		EnableCompression: false,
	}
	headers := http.Header{}
	headers.Add("User-Agent", fmt.Sprintf("%s/%s", Name, Version))
	conn, _, err := Dialer.DialContext(ctx, c.Endpoint, headers)
	if err != nil {
		return nil, err
	}
	conn.SetReadLimit(655350)
	return &WebsocketStreamConn{
		Conn:   conn,
		Client: c,
	}, nil
}

func (c *WebsocketStreamConn) subscribe(channels []SubOpArg) error {
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
	err = c.Conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		return err
	}
	_, message, err := c.Conn.ReadMessage()
	if err != nil {
		return err
	}
	c.Client.debug("Subscribe response: %s\n", message)
	var response SubscribeResponse
	if err := json.Unmarshal(message, &response); err != nil {
		return err
	}
	if response.Code != nil {
		return &ApiError{Code: *response.Code, Message: *response.Msg}
	}
	return nil
}

func (c *WebsocketStreamConn) serve(handler WsHandler, errHandler ErrHandler) (doneCh, stopCh chan struct{}, err error) {
	doneCh = make(chan struct{})
	stopCh = make(chan struct{})
	go func() {
		// This function will exit either on error from
		// websocket.Conn.ReadMessage or when the stopC channel is
		// closed by the client.
		defer close(doneCh)

		lastResponse := time.Now()
		if c.Client.Keepalive {
			timeout := c.Client.Timeout
			if timeout == 0 {
				timeout = time.Second * 10 // 默认值
			}
			go func(conn *websocket.Conn) {
				ticker := time.NewTicker(timeout)
				defer ticker.Stop()
				for {
					select {
					case <-doneCh:
						return
					case <-ticker.C:
						if time.Since(lastResponse) >= timeout {
							c.Client.debug("Send ping message")
							err := conn.WriteMessage(websocket.TextMessage, []byte("ping"))
							if err != nil {
								return
							}
						}
					}
				}
			}(c.Conn)
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
			_, message, err := c.Conn.ReadMessage()
			if err != nil {
				if !silent {
					errHandler(err)
				}
				return
			}
			lastResponse = time.Now()
			if string(message) == "pong" {
				c.Client.debug("Receive pong message")
				continue
			}
			handler(message)
		}
	}()
	return
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
		Endpoint:  url + "/ws/v5/public",
		Logger:    log.New(os.Stderr, Name, log.LstdFlags),
		Timeout:   time.Second * 10,
		Keepalive: true,
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
		Endpoint:  url + "/ws/v5/private",
		Logger:    log.New(os.Stderr, Name, log.LstdFlags),
		Timeout:   time.Second * 10,
		Keepalive: true,
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
		Endpoint:  url + "/ws/v5/business",
		Logger:    log.New(os.Stderr, Name, log.LstdFlags),
		Timeout:   time.Second * 10,
		Keepalive: true,
	}
}
