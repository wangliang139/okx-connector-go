package okx_connector

import (
	"context"
	"log"
	"testing"
)

func Test_OpenOrders(t *testing.T) {
	client := newClient()
	client.Debug = true
	response, err := client.NewOpenOrdersService().Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_Order(t *testing.T) {
	client := newClient()
	client.Debug = true
	response, err := client.NewOrderService().InstId("ETH-USDT-SWAP").Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}
