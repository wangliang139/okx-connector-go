package okx_connector

import (
	"context"
	"log"
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
	response, err := client.NewAccountLeverageInfoService().InstId("ETH-USDT-SWAP").MgnMode("cross").Do(context.Background())
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
