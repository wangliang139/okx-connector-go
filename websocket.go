package okx_connector

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

type WebsocketStreamClient struct {
	Endpoint string
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
		Endpoint: url,
	}
}

func newWsConfig(endpoint string) *WsConfig {
	return &WsConfig{
		Endpoint: endpoint,
	}
}

type SubscribeChannel struct {
	Channel string `json:"channel"`
	InstId  string `json:"instId"`
}

type SubscribeRequest struct {
	Op   string `json:"op"`
	Args any    `json:"args"`
}

type SubscribeResponse struct {
	Event  string          `json:"event"`
	Code   *string         `json:"code"`
	Msg    *string         `json:"msg"`
	ConnId string          `json:"connId"`
	Arg    json.RawMessage `json:"arg"`
}

var wsServe = func(cfg *WsConfig, channels []*SubscribeChannel, handler WsHandler, errHandler ErrHandler) (doneCh, stopCh chan struct{}, err error) {
	if len(channels) == 0 {
		return nil, nil, fmt.Errorf("channels is empty")
	}
	Dialer := websocket.Dialer{
		Proxy:             http.ProxyFromEnvironment,
		HandshakeTimeout:  45 * time.Second,
		EnableCompression: false,
	}
	headers := http.Header{}
	headers.Add("User-Agent", fmt.Sprintf("%s/%s", Name, Version))
	c, _, err := Dialer.Dial(cfg.Endpoint, headers)
	if err != nil {
		return nil, nil, err
	}

	// subscribe
	op := &SubscribeRequest{
		Op:   "subscribe",
		Args: channels,
	}
	msg, err := json.Marshal(op)
	if err != nil {
		return nil, nil, err
	}
	err = c.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		return nil, nil, err
	}
	_, message, err := c.ReadMessage()
	if err != nil {
		return nil, nil, err
	}
	log.Printf("Subscribe response: %s\n", message)
	var response SubscribeResponse
	if err := json.Unmarshal(message, &response); err != nil {
		return nil, nil, err
	}
	if response.Code != nil {
		return nil, nil, &ApiError{Code: *response.Code, Message: *response.Msg}
	}

	c.SetReadLimit(655350)
	doneCh = make(chan struct{})
	stopCh = make(chan struct{})
	go func() {
		// This function will exit either on error from
		// websocket.Conn.ReadMessage or when the stopC channel is
		// closed by the client.
		defer close(doneCh)
		if WebsocketKeepalive {
			keepAlive(c, WebsocketTimeout)
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
			_, message, err := c.ReadMessage()
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
