package okx_connector

import (
	"log"
	"testing"
)

func Test_kline(t *testing.T) {
	client := NewWsPublicStreamClient()
	_, stopCh, err := client.WsKlineServe([]string{"G-USDT"}, "candle1s", func(event *WsKlineEvent) {
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
	client := NewWsPublicStreamClient()
	_, stopCh, err := client.WsDepthServe([]string{"G-USDT"}, "books", func(event *WsDepthEvent) {
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
	WebsocketKeepalive = true
	client := NewWsPublicStreamClient()
	client.Debug = true
	_, stopCh, err := client.WsTradeServe([]string{"G-USDT"}, "trades", func(event *WsTradeEvent) {
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
