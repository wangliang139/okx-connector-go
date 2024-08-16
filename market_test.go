package okx_connector

import (
	"context"
	"log"
	"testing"
)

const (
	Apikey     = "60532a79-4aad-40d8-a0ab-89025c54a781"
	Secretkey  = "D774B3D7CADCD4FDDF715336042AE96C"
	Passphrase = "Jiv@Hceq&zZJ6c"
)

func Test_SymbolInfo(t *testing.T) {
	client := NewClient(Apikey, Secretkey, Passphrase)
	client.Debug = true
	response, err := client.NewSymbolInfoService().InstType("SPOT").InstId("BTC-BRL").Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_MarketCandles(t *testing.T) {
	client := NewClient(Apikey, Secretkey, Passphrase)
	client.Debug = true
	response, err := client.NewMarketKlinesHisService().InstId("BTC-BRL").Bar("1s").Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}
