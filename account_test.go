package okx_connector

import (
	"context"
	"log"
	"os"
	"testing"
	"time"
)

func Test_AccountBalance(t *testing.T) {
	requireIntegrationTests(t)
	client := newClient()
	client.Debug = true
	response, err := client.NewAccountBalanceService().Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_AccountConfig(t *testing.T) {
	requireIntegrationTests(t)
	client := newClient()
	client.Debug = true
	response, err := client.NewAccountConfigService().Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_AccountLeverageInfo(t *testing.T) {
	requireIntegrationTests(t)
	client := newClient()
	client.Debug = true
	response, err := client.NewGetLeverageInfoService().MgnMode("cross").InstId("ETH-USDT").Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_AccountTradeFee(t *testing.T) {
	requireIntegrationTests(t)
	client := newClient()
	client.Debug = true
	response, err := client.NewAccountTradeFeeService().InstType("SWAP").InstFamily("ETH-USDT").Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_WsAccountServe(t *testing.T) {
	requireIntegrationTests(t)
	Apikey := os.Getenv("OKX_API_KEY")
	Secretkey := os.Getenv("OKX_SECRET_KEY")
	Passphrase := os.Getenv("OKX_PASSPHRASE")
	if Apikey == "" || Secretkey == "" || Passphrase == "" {
		t.Skip("OKX credentials not set; skipping websocket integration test")
	}
	client := NewWsStreamClient(WithWsAPIAuth(Apikey, Secretkey, Passphrase))
	client.Debug = true
	handler := &TestWsUserDataHandler{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	doneCh, stopCh, err := client.WsUserDataServe(ctx, handler, func(err error) {
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
	case <-time.After(2 * time.Second):
		// Stop the reader loop and exit the test.
		close(stopCh)
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

func Test_SetLeverage(t *testing.T) {
	requireIntegrationTests(t)
	client := newClient()
	client.Debug = true
	response, err := client.NewSetLeverageService().InstId("ETH-USDT").MgnMode("cross").Lever("10").Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}
