package okx_connector

import (
	"context"
	"log"
	"os"
	"testing"
)

func Test_AccountBalance(t *testing.T) {
	client := newClient()
	client.Debug = true
	response, err := client.NewAccountBalanceService().Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_AccountConfig(t *testing.T) {
	client := newClient()
	client.Debug = true
	response, err := client.NewAccountConfigService().Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_AccountLeverageInfo(t *testing.T) {
	client := newClient()
	client.Debug = true
	response, err := client.NewAccountLeverageInfoService().MgnMode("cross").Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_AccountTradeFee(t *testing.T) {
	client := newClient()
	client.Debug = true
	response, err := client.NewAccountTradeFeeService().InstType("SWAP").Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_WsAccountServe(t *testing.T) {
	Apikey := os.Getenv("OKX_API_KEY")
	Secretkey := os.Getenv("OKX_SECRET_KEY")
	Passphrase := os.Getenv("OKX_PASSPHRASE")
	client := NewWsStreamClient(WithWsAPIAuth(Apikey, Secretkey, Passphrase))
	client.Debug = true
	handler := &TestWsUserDataHandler{}
	doneCh, stopCh, err := client.WsUserDataServe(context.Background(), handler, func(err error) {
		log.Printf("%+v", err)
	})
	if err != nil {
		t.Fatal(err)
	}
	select {
	case <-doneCh:
		return
	case <-stopCh:
		return
	}
}

type TestWsUserDataHandler struct{}

func (h *TestWsUserDataHandler) HandleAccountEvent(event *WsAccountEvent) {
	// log.Printf("%+v", event)
}

func (h *TestWsUserDataHandler) HandlePositionEvent(event *WsPositionEvent) {
	// log.Printf("%+v", event)
}

func (h *TestWsUserDataHandler) HandleOrderEvent(event *WsOrderEvent) {
	// log.Printf("%+v", event)
}

func (h *TestWsUserDataHandler) HandleBalanceAndPositionEvent(event *WsBalanceAndPositionEvent) {
	// log.Printf("%+v", event)
}
