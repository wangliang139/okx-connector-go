package okx_connector

import (
	"context"
	"log"
	"testing"
)

func Test_kline(t *testing.T) {
	client := NewWsStreamClient("", "", "", "")
	_, stopCh, err := client.WsKlineServe(context.Background(), []string{"SOL-USDT"}, "candle1s", func(event *WsKlineEvent) {
		log.Printf("%+v", event)
	}, func(err error) {
		log.Printf("%+v", err)
	})
	if err != nil {
		t.Fatal(err)
	}
	select {
	case <-stopCh:
		return
	}
}

func Test_depth(t *testing.T) {
	client := NewWsStreamClient("", "", "", "")
	_, stopCh, err := client.WsDepthServe(context.Background(), []string{"SOL-USDT"}, "books", func(event *WsDepthEvent) {
		log.Printf("%d, %d", event.Data[0].PrevSeqId, event.Data[0].SeqId)
	}, func(err error) {
		log.Printf("%+v", err)
	})
	if err != nil {
		t.Fatal(err)
	}
	select {
	case <-stopCh:
		return
	}
}

func Test_trade(t *testing.T) {
	client := NewWsStreamClient("", "", "", "")
	client.Debug = true
	client.Keepalive = true
	_, stopCh, err := client.WsTradeServe(context.Background(), []string{"SOL-USDT"}, func(event *WsTradeEvent) {
		log.Printf("%+v", event)
	}, func(err error) {
		log.Printf("%+v", err)
	})
	if err != nil {
		t.Fatal(err)
	}
	select {
	case <-stopCh:
		return
	}
}

func Test_ticker(t *testing.T) {
	client := NewWsStreamClient("", "", "", "")
	_, stopCh, err := client.WsTickerServe(context.Background(), []string{"SOL-USDT-SWAP"}, func(event *WsTickerEvent) {
		log.Printf("%+v", event)
	}, func(err error) {
		log.Printf("%+v", err)
	})
	if err != nil {
		t.Fatal(err)
	}
	select {
	case <-stopCh:
		return
	}
}
