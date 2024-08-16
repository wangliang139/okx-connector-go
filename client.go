package okx_connector

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	jsoniter "github.com/json-iterator/go"
)

// Client define API client
type Client struct {
	APIKey     string
	SecretKey  string
	Passphrase string
	BaseURL    string
	HTTPClient *http.Client
	Debug      bool
	Logger     *log.Logger
	TimeOffset int64
	do         doFunc
}

type ApiResponse struct {
	Code string          `json:"code"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data"`
}

type doFunc func(req *http.Request) (*http.Response, error)

// Globals
const (
	recvWindowKey = "recvWindow"
)

func currentTime() string {
	return time.Now().Format(time.RFC3339)
}

func currentTimestamp() int64 {
	return FormatTimestamp(time.Now())
}

// FormatTimestamp formats a time into Unix timestamp in milliseconds, as requested by Binance.
func FormatTimestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func (c *Client) debug(format string, v ...interface{}) {
	if c.Debug {
		c.Logger.Printf(format, v...)
	}
}

// Create client function for initialising new Binance client
func NewClient(apiKey, secretKey, passphrase string, baseURL ...string) *Client {
	url := "https://www.okx.com"

	if len(baseURL) > 0 {
		for _, u := range baseURL {
			if len(u) > 0 {
				url = u
				break
			}
		}
	}

	return &Client{
		APIKey:     apiKey,
		SecretKey:  secretKey,
		Passphrase: passphrase,
		BaseURL:    url,
		HTTPClient: http.DefaultClient,
		Logger:     log.New(os.Stderr, Name, log.LstdFlags),
	}
}

func (c *Client) parseRequest(r *request, opts ...RequestOption) (err error) {
	// set request options from user
	for _, opt := range opts {
		opt(r)
	}
	err = r.validate()
	if err != nil {
		return err
	}

	ctime := currentTime()

	fullURL := fmt.Sprintf("%s%s", c.BaseURL, r.endpoint)
	if r.recvWindow > 0 {
		r.setParam(recvWindowKey, r.recvWindow)
	}
	queryString := r.query.Encode()
	body := &bytes.Buffer{}
	bodyString := r.form.Encode()
	header := http.Header{}
	if r.header != nil {
		header = r.header.Clone()
	}
	header.Set("User-Agent", fmt.Sprintf("%s/%s", Name, Version))
	header.Set("Content-Type", "application/json")
	header.Set("OK-ACCESS-TIMESTAMP", ctime)
	if bodyString != "" {
		body = bytes.NewBufferString(bodyString)
	}
	if r.secType == secTypeAPIKey || r.secType == secTypeSigned {
		header.Set("OK-ACCESS-PASSPHRASE", c.Passphrase)
		header.Set("OK-ACCESS-KEY", c.APIKey)
	}
	if r.secType == secTypeSigned {
		raw := fmt.Sprintf("%s%s%s", ctime, r.method, r.endpoint)
		if len(queryString) > 0 {
			raw = fmt.Sprintf("%s?%s", raw, queryString)
		}
		c.debug("before sign: %s", raw)
		mac := hmac.New(sha256.New, []byte(c.SecretKey))
		_, err = mac.Write([]byte(raw))
		if err != nil {
			return err
		}
		sign := base64.StdEncoding.EncodeToString(mac.Sum(nil))
		header.Set("OK-ACCESS-SIGN", sign)
	}
	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}
	c.debug("full url: %s, body: %s", fullURL, bodyString)
	r.fullURL = fullURL
	r.header = header
	r.body = body
	return nil
}

func (c *Client) callAPI(ctx context.Context, r *request, opts ...RequestOption) (data []byte, err error) {
	err = c.parseRequest(r, opts...)
	if err != nil {
		return []byte{}, err
	}
	req, err := http.NewRequest(r.method, r.fullURL, r.body)
	if err != nil {
		return []byte{}, err
	}
	req = req.WithContext(ctx)
	req.Header = r.header
	c.debug("request: %#v", req)
	f := c.do
	if f == nil {
		f = c.HTTPClient.Do
	}
	res, err := f(req)
	if err != nil {
		return []byte{}, err
	}
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	defer func() {
		cerr := res.Body.Close()
		// Only overwrite the retured error if the original error was nil and an
		// error occurred while closing the body.
		if err == nil && cerr != nil {
			err = cerr
		}
	}()
	c.debug("response: %#v", res)
	c.debug("response body: %s", string(data))
	c.debug("response status code: %d", res.StatusCode)

	response := new(ApiResponse)
	e := jsoniter.Unmarshal(data, response)
	if e != nil {
		c.debug("failed to unmarshal json: %s", e)
		return nil, e
	}
	if response.Code != "0" {
		return nil, &ApiError{Code: response.Code, Message: response.Msg}
	}
	return response.Data, nil
}

// Market Endpoints:
func (c *Client) NewPingService() *PingService {
	return &PingService{c: c}
}

func (c *Client) NewSystemTimeService() *SystemTimeService {
	return &SystemTimeService{c: c}
}

func (c *Client) NewSystemStatusService() *SystemStatusService {
	return &SystemStatusService{c: c}
}

func (c *Client) NewSymbolInfoService() *SymbolInfoService {
	return &SymbolInfoService{c: c}
}

func (c *Client) NewMarketCandlesService() *MarketCandlesService {
	return &MarketCandlesService{c: c}
}
