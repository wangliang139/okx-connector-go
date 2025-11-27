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
