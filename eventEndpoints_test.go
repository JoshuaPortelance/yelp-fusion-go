package yelp

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestNewEventSearchAndLookup(t *testing.T) {
	client := Client{
		key: os.Getenv("YELP_API_KEY"),
	}
	eventsReq := client.NewEventSearch()
	events, eventsErr := eventsReq.Get()
	if eventsErr != nil {
		t.Fatalf(`Unexpected error: %q`, eventsErr)
	}
	if events.Total == 0 {
		t.Fatalf(`Didn't find any events. %v+`, events)
	}
	if events.Events[0].Name == "" {
		t.Fatalf(`Invalid event. %v+`, events)
	}

	// Doing a detailed lookup using the found Id.
	eventReq := client.NewEventLookup(events.Events[0].Id)
	event, eventErr := eventReq.Get()
	if eventErr != nil {
		t.Fatalf(`Unexpected error: %q`, eventErr)
	}
	if event.Name == "" {
		t.Fatalf(`Invalid event. %v+`, event)
	}
}

func TestNewEventSearch(t *testing.T) {
	client := Client{
		key: os.Getenv("YELP_API_KEY"),
	}
	eventsReq := client.NewEventSearch()
	data, err := eventsReq.Get()
	if err != nil {
		t.Fatalf(`Unexpected error: %q`, err)
	}
	if data.Total == 0 {
		t.Fatalf(`Didn't find any events. %v+`, data)
	}
	if data.Events[0].Name == "" {
		t.Fatalf(`Invalid event. %v+`, data)
	}
}

func TestNewEventSearchResponse(t *testing.T) {
	client := Client{
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
	client := Client{
		key: os.Getenv("YELP_API_KEY"),
	}
	eventReq := client.NewFeaturedEvent()
	eventReq.AddParameter("location", "Seattle, Washington")
	data, err := eventReq.Get()
	if err != nil {
		t.Fatalf(`Unexpected error: %q`, err)
	}
	if data.Name == "" {
		t.Fatalf(`Invalid event. %v+`, data)
	}
}

func TestNewFeaturedEventGetResponse(t *testing.T) {
	client := Client{
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
