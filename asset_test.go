package okx_connector

import (
	"context"
	"log"
	"testing"
)

func Test_AssetCurrencies(t *testing.T) {
	client := newClient()
	client.Debug = true
	response, err := client.NewFundingAssetCurrenciesService().Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_FundingAssetBalances(t *testing.T) {
	client := newClient()
	client.Debug = true
	response, err := client.NewFundingAssetBalancesService().Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_FundingAssetValuation(t *testing.T) {
	client := newClient()
	client.Debug = true
	response, err := client.NewFundingAssetValuationService().Ccy("USDT").Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}