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

func Test_Orders7DHistory(t *testing.T) {
	client := newClient()
	client.Debug = true
	response, err := client.NewOrders7DHistoryService().InstType("SWAP").Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_OrdersHistory3MService(t *testing.T) {
	client := newClient()
	client.Debug = true
	response, err := client.NewOrdersHistory3MService().InstType("SWAP").Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_OpenAlgoOrders(t *testing.T) {
	client := newClient()
	client.Debug = true
	response, err := client.NewOpenAlgoOrdersService().InstType("SWAP").Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}
