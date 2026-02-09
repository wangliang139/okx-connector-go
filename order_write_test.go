package okx_connector

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestPlaceOrderService_BuildsBodyAndParsesResponse(t *testing.T) {
	c := NewClient(WithApiAPIAuth("k", "s", "p"))
	c.do = func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodPost {
			t.Fatalf("method = %q, want %q", req.Method, http.MethodPost)
		}
		if req.URL.Path != "/api/v5/trade/order" {
			t.Fatalf("path = %q, want %q", req.URL.Path, "/api/v5/trade/order")
		}
		if got := req.Header.Get("OK-ACCESS-KEY"); got != "k" {
			t.Fatalf("OK-ACCESS-KEY = %q, want %q", got, "k")
		}
		if got := req.Header.Get("OK-ACCESS-PASSPHRASE"); got != "p" {
			t.Fatalf("OK-ACCESS-PASSPHRASE = %q, want %q", got, "p")
		}
		if got := req.Header.Get("OK-ACCESS-TIMESTAMP"); got == "" {
			t.Fatalf("OK-ACCESS-TIMESTAMP is empty")
		}
		if got := req.Header.Get("OK-ACCESS-SIGN"); got == "" {
			t.Fatalf("OK-ACCESS-SIGN is empty")
		}

		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			t.Fatalf("read body: %v", err)
		}
		var gotBody map[string]any
		if err := json.Unmarshal(bodyBytes, &gotBody); err != nil {
			t.Fatalf("unmarshal body: %v, body=%s", err, string(bodyBytes))
		}

		want := map[string]any{
			"instId":  "BTC-USDT",
			"tdMode":  "cash",
			"side":    "buy",
			"ordType": "limit",
			"px":      "1000",
			"sz":      "0.01",
			"tag":     "t1",
		}
		for k, v := range want {
			if gotBody[k] != v {
				t.Fatalf("body[%q] = %#v, want %#v (full=%v)", k, gotBody[k], v, gotBody)
			}
		}

		resp := `{"code":"0","msg":"","data":[{"clOrdId":"","ordId":"123","tag":"t1","sCode":"0","sMsg":""}]}`
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewBufferString(resp)),
		}, nil
	}

	res, err := c.NewPlaceOrderService().
		InstId("BTC-USDT").
		TdMode("cash").
		Side("buy").
		OrdType("limit").
		Px("1000").
		Sz("0.01").
		Tag("t1").
		Do(context.Background())
	if err != nil {
		t.Fatalf("Do error: %v", err)
	}
	if len(res) != 1 {
		t.Fatalf("len(res) = %d, want %d", len(res), 1)
	}
	if res[0].OrdId != "123" || res[0].SCode != "0" {
		t.Fatalf("res[0] = %+v, want ordId=123 sCode=0", res[0])
	}
}

func TestCancelBatchOrdersService_SendsArrayBody(t *testing.T) {
	c := NewClient(WithApiAPIAuth("k", "s", "p"))
	c.do = func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodPost {
			t.Fatalf("method = %q, want %q", req.Method, http.MethodPost)
		}
		if req.URL.Path != "/api/v5/trade/cancel-batch-orders" {
			t.Fatalf("path = %q, want %q", req.URL.Path, "/api/v5/trade/cancel-batch-orders")
		}
		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			t.Fatalf("read body: %v", err)
		}
		var gotBody []map[string]any
		if err := json.Unmarshal(bodyBytes, &gotBody); err != nil {
			t.Fatalf("unmarshal body: %v, body=%s", err, string(bodyBytes))
		}
		if len(gotBody) != 2 {
			t.Fatalf("len(body) = %d, want %d, body=%v", len(gotBody), 2, gotBody)
		}
		if gotBody[0]["instId"] != "BTC-USDT" || gotBody[0]["ordId"] != "1" {
			t.Fatalf("body[0] = %v, want instId=BTC-USDT ordId=1", gotBody[0])
		}
		if gotBody[1]["instId"] != "ETH-USDT" || gotBody[1]["clOrdId"] != "c2" {
			t.Fatalf("body[1] = %v, want instId=ETH-USDT clOrdId=c2", gotBody[1])
		}

		resp := `{"code":"0","msg":"","data":[{"clOrdId":"","ordId":"1","sCode":"0","sMsg":""},{"clOrdId":"c2","ordId":"","sCode":"0","sMsg":""}]}`
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewBufferString(resp)),
		}, nil
	}

	res, err := c.NewCancelBatchOrdersService().
		Orders([]CancelOrderRequest{
			{InstId: "BTC-USDT", OrdId: "1"},
			{InstId: "ETH-USDT", ClOrdId: "c2"},
		}).
		Do(context.Background())
	if err != nil {
		t.Fatalf("Do error: %v", err)
	}
	if len(res) != 2 {
		t.Fatalf("len(res) = %d, want %d", len(res), 2)
	}
}

func TestCancelAlgoOrderService_SendsArrayBody(t *testing.T) {
	c := NewClient(WithApiAPIAuth("k", "s", "p"))
	c.do = func(req *http.Request) (*http.Response, error) {
		if req.URL.Path != "/api/v5/trade/cancel-algo-order" {
			t.Fatalf("path = %q, want %q", req.URL.Path, "/api/v5/trade/cancel-algo-order")
		}
		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			t.Fatalf("read body: %v", err)
		}
		var gotBody []map[string]any
		if err := json.Unmarshal(bodyBytes, &gotBody); err != nil {
			t.Fatalf("unmarshal body: %v, body=%s", err, string(bodyBytes))
		}
		if len(gotBody) != 1 || gotBody[0]["algoId"] != "a1" {
			t.Fatalf("body = %v, want [{algoId:a1}]", gotBody)
		}

		resp := `{"code":"0","msg":"","data":[{"algoId":"a1","algoClOrdId":"","sCode":"0","sMsg":""}]}`
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewBufferString(resp)),
		}, nil
	}

	res, err := c.NewCancelAlgoOrderService().
		Orders([]CancelAlgoOrderRequest{{AlgoId: "a1"}}).
		Do(context.Background())
	if err != nil {
		t.Fatalf("Do error: %v", err)
	}
	if len(res) != 1 || res[0].AlgoId != "a1" {
		t.Fatalf("res = %v, want algoId=a1", res)
	}
}

