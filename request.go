package okx_connector

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type secType int

const (
	secTypeNone secType = iota
	secTypeAPIKey
	secTypeSigned // if the 'timestamp' parameter is required
)

type params map[string]interface{}

// request define an API request
type request struct {
	method     string
	endpoint   string
	query      url.Values
	form       map[string]any
	// bodyObj is an optional raw JSON body payload.
	// When set, it will be JSON-marshaled and used as request body.
	// This is useful for endpoints requiring an array body, e.g. batch APIs.
	bodyObj    any
	recvWindow int64
	secType    secType
	header     http.Header
	body       io.Reader
	fullURL    string
}

// addParam add param with key/value to query string
func (r *request) addParam(key string, value interface{}) *request {
	if r.query == nil {
		r.query = url.Values{}
	}
	r.query.Add(key, fmt.Sprintf("%v", value))
	return r
}

// setParam set param with key/value to query string
func (r *request) setParam(key string, value interface{}) *request {
	if r.query == nil {
		r.query = url.Values{}
	}
	r.query.Set(key, fmt.Sprintf("%v", value))
	return r
}

// setParams set params with key/values to query string
func (r *request) setParams(m params) *request {
	for k, v := range m {
		r.setParam(k, v)
	}
	return r
}

func (r *request) validate() (err error) {
	if r.query == nil {
		r.query = url.Values{}
	}
	return nil
}

// Append `WithRecvWindow(insert_recvwindow)` to request to modify the default recvWindow value
func WithRecvWindow(recvWindow int64) RequestOption {
	return func(r *request) {
		r.recvWindow = recvWindow
	}
}

// RequestOption define option type for request
type RequestOption func(*request)
