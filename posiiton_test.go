package okx_connector

import (
	"context"
	"log"
	"testing"
)

func Test_AccountPositions(t *testing.T) {
	client := newClient()
	client.Debug = true
	response, err := client.NewPositionsService().Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}
