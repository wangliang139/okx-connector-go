package okx_connector

import (
	"context"
	"log"
	"os"
	"testing"
	"time"
)

func Test_kline(t *testing.T) {
	if os.Getenv("OKX_RUN_WS_TESTS") == "" {
		t.Skip("set OKX_RUN_WS_TESTS=1 to run websocket integration tests")
	}
	client := NewWsStreamClient()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	doneCh, stopCh, err := client.WsKlineServe(ctx, []string{"SOL-USDT"}, "candle1s", func(event *WsKlineEvent) {
		log.Printf("%+v", event)
	}, func(err error) {
		log.Printf("%+v", err)
	})
	if err != nil {
		t.Fatal(err)
	}
	select {
	case <-doneCh:
		return
	case <-time.After(2 * time.Second):
		close(stopCh)
		return
	case <-stopCh:
		return
	}
}

func Test_depth(t *testing.T) {
	if os.Getenv("OKX_RUN_WS_TESTS") == "" {
		t.Skip("set OKX_RUN_WS_TESTS=1 to run websocket integration tests")
	}
	client := NewWsStreamClient()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	doneCh, stopCh, err := client.WsDepthServe(ctx, []string{"SOL-USDT"}, "books", func(event *WsDepthEvent) {
		log.Printf("%d, %d", event.Data[0].PrevSeqId, event.Data[0].SeqId)
	}, func(err error) {
		log.Printf("%+v", err)
	})
	if err != nil {
		t.Fatal(err)
	}
	select {
	case <-doneCh:
		return
	case <-time.After(2 * time.Second):
		close(stopCh)
		return
	case <-stopCh:
		return
	}
}

func Test_trade(t *testing.T) {
	if os.Getenv("OKX_RUN_WS_TESTS") == "" {
		t.Skip("set OKX_RUN_WS_TESTS=1 to run websocket integration tests")
	}
	client := NewWsStreamClient()
	client.Debug = true
	client.Keepalive = true
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	doneCh, stopCh, err := client.WsTradeServe(ctx, []string{"SOL-USDT"}, func(event *WsTradeEvent) {
		log.Printf("%+v", event)
	}, func(err error) {
		log.Printf("%+v", err)
	})
	if err != nil {
		t.Fatal(err)
	}
	select {
	case <-doneCh:
		return
	case <-time.After(2 * time.Second):
		close(stopCh)
		return
	case <-stopCh:
		return
	}
}

func Test_ticker(t *testing.T) {
	if os.Getenv("OKX_RUN_WS_TESTS") == "" {
		t.Skip("set OKX_RUN_WS_TESTS=1 to run websocket integration tests")
	}
	client := NewWsStreamClient()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	doneCh, stopCh, err := client.WsTickerServe(ctx, []string{"SOL-USDT-SWAP"}, func(event *WsTickerEvent) {
		log.Printf("%+v", event)
	}, func(err error) {
		log.Printf("%+v", err)
	})
	if err != nil {
		t.Fatal(err)
	}
	select {
	case <-doneCh:
		return
	case <-time.After(2 * time.Second):
		close(stopCh)
		return
	case <-stopCh:
		return
	}
}
