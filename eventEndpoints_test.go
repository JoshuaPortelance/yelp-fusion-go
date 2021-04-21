package yelp

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestNewEventSearchAndLookup(t *testing.T) {
	client := YelpClient{
		key: os.Getenv("YELP_API_KEY"),
	}
	eventsReq := client.NewEventSearch()
	events, eventsErr := eventsReq.Get()
	if eventsErr != nil {
		t.Fatalf(`Unexpected error: %q`, eventsErr)
	}

	// Doing a detailed lookup using the found Id.
	eventReq := client.NewEventLookup(events.Events[0].Id)
	_, eventErr := eventReq.Get()
	if eventErr != nil {
		t.Fatalf(`Unexpected error: %q`, eventErr)
	}
}

func TestNewEventSearch(t *testing.T) {
	client := YelpClient{
		key: os.Getenv("YELP_API_KEY"),
	}
	eventsReq := client.NewEventSearch()
	_, err := eventsReq.Get()
	if err != nil {
		t.Fatalf(`Unexpected error: %q`, err)
	}
}

func TestNewEventSearchResponse(t *testing.T) {
	client := YelpClient{
		key: os.Getenv("YELP_API_KEY"),
	}
	eventsReq := client.NewEventSearch()
	response, err := eventsReq.GetResponse()
	if err != nil {
		body, _ := io.ReadAll(response.Body)
		fmt.Println(string(body))
		t.Fatalf(`Unexpected error: %q`, err)
	}
	defer response.Body.Close()
	if response.StatusCode == 0 {
		t.Fatalf(`Invalid status: %d`, response.StatusCode)
	}
	if response.StatusCode != 200 {
		t.Fatalf(`Unexpected status: %d`, response.StatusCode)
	}
}

func TestNewFeaturedEventGet(t *testing.T) {
	client := YelpClient{
		key: os.Getenv("YELP_API_KEY"),
	}
	eventReq := client.NewFeaturedEvent()
	eventReq.AddParameter("location", "Vancouver, Canada")
	_, err := eventReq.Get()
	if err != nil {
		t.Fatalf(`Unexpected error: %q`, err)
	}
}

func TestNewFeaturedEventGetResponse(t *testing.T) {
	client := YelpClient{
		key: os.Getenv("YELP_API_KEY"),
	}
	eventReq := client.NewFeaturedEvent()
	eventReq.AddParameter("location", "Vancouver, Canada")
	response, err := eventReq.GetResponse()
	if err != nil {
		body, _ := io.ReadAll(response.Body)
		fmt.Println(string(body))
		t.Fatalf(`Unexpected error: %q`, err)
	}
	defer response.Body.Close()
	if response.StatusCode == 0 {
		t.Fatalf(`Invalid status: %d`, response.StatusCode)
	}
	if response.StatusCode != 200 {
		t.Fatalf(`Unexpected status: %d`, response.StatusCode)
	}
}
