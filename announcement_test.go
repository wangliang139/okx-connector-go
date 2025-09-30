package okx_connector

import (
	"context"
	"log"
	"testing"
)

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
	response, err := client.NewAnnouncementService().AnnType("announcements-new-listings").Page(1).Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", response)
}