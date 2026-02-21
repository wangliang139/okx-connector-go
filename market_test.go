package okx_connector

import (
	"context"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

func newClient() *Client {
	Apikey := os.Getenv("OKX_API_KEY")
	Secretkey := os.Getenv("OKX_SECRET_KEY")
	Passphrase := os.Getenv("OKX_PASSPHRASE")
	httpClient := &http.Client{Timeout: 10 * time.Second}
	return NewClient(
		WithApiAPIAuth(Apikey, Secretkey, Passphrase),
		WithApiHTTPClient(httpClient),
	)
}

func Test_SymbolInfo(t *testing.T) {
	requireIntegrationTests(t)
	client := newClient()
	client.Debug = true
	response, err := client.NewSymbolInfoService().InstType("SWAP").InstId("ETH-USDT-SWAP").Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_MarketCandles(t *testing.T) {
	requireIntegrationTests(t)
	client := newClient()
	client.Debug = true
	response, err := client.NewMarketKlinesHisService().InstId("G-USDT").Bar("1s").Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_SymbolQuotationService(t *testing.T) {
	requireIntegrationTests(t)
	client := newClient()
	client.Debug = true
	response, err := client.NewSymbolQuotationService().InstId("G-USDT").Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_MarketDepthService(t *testing.T) {
	requireIntegrationTests(t)
	client := newClient()
	client.Debug = true
	response, err := client.NewMarketDepthService().InstId("G-USDT").Size(100).Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_MarketDepthFullService(t *testing.T) {
	requireIntegrationTests(t)
	client := newClient()
	client.Debug = true
	response, err := client.NewMarketDepthFullService().InstId("G-USDT").Size(1000).Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_AnnouncementType(t *testing.T) {
	requireIntegrationTests(t)
	client := newClient()
	client.Debug = true
	response, err := client.NewAnnouncementTypeService().Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_Announcement(t *testing.T) {
	requireIntegrationTests(t)
	client := newClient()
	client.Debug = true
	response, err := client.NewAnnouncementService().Page(1).Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_MarketTrades(t *testing.T) {
	requireIntegrationTests(t)
	client := newClient()
	client.Debug = true
	response, err := client.NewMarketTradesService().InstId("G-USDT").Limit(100).Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_MarketTickers(t *testing.T) {
	requireIntegrationTests(t)
	client := newClient()
	client.Debug = true
	response, err := client.NewMarketTickersService().InstType("SWAP").Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_PositionTiers(t *testing.T) {
	requireIntegrationTests(t)
	client := newClient()
	client.Debug = true
	response, err := client.NewPositionTiersService().TdMode("cross").InstType("SWAP").InstFamily("ETH-USDT").Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_MarkPrice(t *testing.T) {
	requireIntegrationTests(t)
	client := newClient()
	client.Debug = true
	response, err := client.NewMarkPriceService().InstType("SWAP").InstId("ETH-USDT-SWAP").Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_FundingRate(t *testing.T) {
	requireIntegrationTests(t)
	client := newClient()
	client.Debug = true
	response, err := client.NewFundingRateService().InstId("ETH-USDT-SWAP").Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_FundingRateHistory(t *testing.T) {
	// requireIntegrationTests(t)
	client := newClient()
	client.Debug = true
	response, err := client.NewFundingRateHistoryService().InstId("ETH-USDT-SWAP").Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}
