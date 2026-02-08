package okx_connector

import "testing"

func TestClient_IsTestNet_AddsSimulatedTradingHeader(t *testing.T) {
	c := NewClient(WithApiIsTestNet(true))
	r := &request{
		method:   "GET",
		endpoint: "/api/v5/public/time",
		secType:  secTypeNone,
	}
	if err := c.parseRequest(r); err != nil {
		t.Fatalf("parseRequest error: %v", err)
	}
	if got := r.header.Get("x-simulated-trading"); got != "1" {
		t.Fatalf("x-simulated-trading header = %q, want %q", got, "1")
	}
}

func TestClient_NotTestNet_DoesNotAddSimulatedTradingHeader(t *testing.T) {
	c := NewClient(WithApiIsTestNet(false))
	r := &request{
		method:   "GET",
		endpoint: "/api/v5/public/time",
		secType:  secTypeNone,
	}
	if err := c.parseRequest(r); err != nil {
		t.Fatalf("parseRequest error: %v", err)
	}
	if got := r.header.Get("x-simulated-trading"); got != "" {
		t.Fatalf("x-simulated-trading header = %q, want empty", got)
	}
}

func TestWebsocketStreamClient_WithWsIsTestNet_SetsDemoEndpoint(t *testing.T) {
	ws := NewWsStreamClient(WithWsIsTestNet(true))
	if ws.BaseURL != "wss://wspap.okx.com:8443" {
		t.Fatalf("BaseURL = %q, want %q", ws.BaseURL, "wss://wspap.okx.com:8443")
	}
}

func TestWebsocketStreamClient_WithWsIsTestNetFalse_UsesProdEndpoint(t *testing.T) {
	ws := NewWsStreamClient(WithWsIsTestNet(false))
	if ws.BaseURL != "wss://ws.okx.com:8443" {
		t.Fatalf("BaseURL = %q, want %q", ws.BaseURL, "wss://ws.okx.com:8443")
	}
}
