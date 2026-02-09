package okx_connector

import (
	"context"
	"log"
	"os"
	"testing"
)

func Test_OpenOrders(t *testing.T) {
	requireIntegrationTests(t)
	client := newClient()
	client.Debug = true
	response, err := client.NewOpenOrdersService().Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_Order(t *testing.T) {
	requireIntegrationTests(t)
	client := newClient()
	client.Debug = true
	instId := os.Getenv("OKX_TEST_INST_ID")
	ordId := os.Getenv("OKX_TEST_ORD_ID")
	clOrdId := os.Getenv("OKX_TEST_CL_ORD_ID")
	if instId == "" || (ordId == "" && clOrdId == "") {
		t.Skip("set OKX_TEST_INST_ID and one of OKX_TEST_ORD_ID/OKX_TEST_CL_ORD_ID to run this test")
	}
	svc := client.NewOrderService().InstId(instId)
	if ordId != "" {
		svc.OrdId(ordId)
	}
	if clOrdId != "" {
		svc.ClOrdId(clOrdId)
	}
	response, err := svc.Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_Orders7DHistory(t *testing.T) {
	requireIntegrationTests(t)
	client := newClient()
	client.Debug = true
	response, err := client.NewOrders7DHistoryService().InstType("SWAP").Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_OrdersHistory3MService(t *testing.T) {
	requireIntegrationTests(t)
	client := newClient()
	client.Debug = true
	response, err := client.NewOrdersHistory3MService().InstType("SWAP").Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}

func Test_OpenAlgoOrders(t *testing.T) {
	requireIntegrationTests(t)
	client := newClient()
	client.Debug = true
	// ordType is required by OKX for this endpoint.
	response, err := client.NewOpenAlgoOrdersService().InstType("SWAP").OrdType("trigger").Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}
