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
