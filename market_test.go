package okx_connector

import (
	"context"
	"log"
	"os"
	"testing"
)

func newClient() *Client {
	Apikey := os.Getenv("OKX_API_KEY")
	Secretkey := os.Getenv("OKX_SECRET_KEY")
	Passphrase := os.Getenv("OKX_PASSPHRASE")
	return NewClient(Apikey, Secretkey, Passphrase)
}

func Test_SymbolInfo(t *testing.T) {
	client := newClient()
	client.Debug = true
	response, err := client.NewSymbolInfoService().InstType("SWAP").InstId("ETH-USDT-SWAP").Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_MarketCandles(t *testing.T) {
	client := newClient()
	client.Debug = true
	response, err := client.NewMarketKlinesHisService().InstId("G-USDT").Bar("1s").Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_SymbolQuotationService(t *testing.T) {
	client := newClient()
	client.Debug = true
	response, err := client.NewSymbolQuotationService().InstId("G-USDT").Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_MarketDepthService(t *testing.T) {
	client := newClient()
	client.Debug = true
	response, err := client.NewMarketDepthService().InstId("G-USDT").Size(100).Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_MarketDepthFullService(t *testing.T) {
	client := newClient()
	client.Debug = true
	response, err := client.NewMarketDepthFullService().InstId("G-USDT").Size(1000).Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_AnnouncementType(t *testing.T) {
	client := newClient()
	client.Debug = true
	response, err := client.NewAnnouncementTypeService().Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_Announcement(t *testing.T) {
	client := newClient()
	client.Debug = true
	response, err := client.NewAnnouncementService().Page(1).Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}
