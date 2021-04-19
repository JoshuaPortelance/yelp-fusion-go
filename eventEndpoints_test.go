package yelp

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func TestNewEventSearchAndLookup(t *testing.T) {
	client := YelpClient{
		key: os.Getenv("YELP_API_KEY"),
	}

	// Getting a list of events.
	eventsReq := client.NewEventSearch()
	eventsRes, eventsErr := eventsReq.Get()
	if eventsErr != nil {
		t.Fatalf(`Unexpected error getting events: %q`, eventsErr)
	}
	if eventsRes.StatusCode == 0 {
		t.Fatalf(`Invalid response getting events: %d`, eventsRes.StatusCode)
	}
	if eventsRes.StatusCode != 200 {
		t.Fatalf(`Unexpected response getting events: %d`, eventsRes.StatusCode)
	}

	// Reading and unmarshalling body so we can extract a valid Id.
	eventsResbody, err := ioutil.ReadAll(eventsRes.Body)
	if err != nil {
		t.Fatalf(`Unexpected error reading body: %q`, err)
	}
	defer eventsRes.Body.Close()

	events := Events{}
	err = json.Unmarshal(eventsResbody, &events)
	if err != nil {
		t.Fatalf(`Unexpected error unmarshalling body: %q`, err)
	}
	if events.Total == 0 {
		t.Fatalf(`Unable to find any events or unmarshalling failed: %+v`, events)
	}
	if events.Events[0].Id == "" {
		t.Fatalf(`events.Event[0].Id is invalid %+v`, events.Events[0])
	}

	// Doing a detailed lookup using the found Id.
	eventReq := client.NewEventLookup(events.Events[0].Id)
	eventRes, eventErr := eventReq.Get()
	if eventErr != nil {
		t.Fatalf(`Unexpected error getting event details: %q`, eventErr)
	}
	if eventRes.StatusCode == 0 {
		t.Fatalf(`Invalid response getting event details: %d`, eventRes.StatusCode)
	}
	if eventRes.StatusCode != 200 {
		t.Fatalf(`Unexpected response getting event details: %d`, eventRes.StatusCode)
	}

	// Reading and unmarshalling body so we can verify the Id.
	eventResbody, err := ioutil.ReadAll(eventRes.Body)
	if err != nil {
		t.Fatalf(`Unexpected error reading body: %q`, err)
	}

	defer eventRes.Body.Close()

	event := Event{}
	err = json.Unmarshal(eventResbody, &event)
	if err != nil {
		t.Fatalf(`Unexpected error unmarshalling body: %q`, err)
	}

	if event.Id == "" {
		t.Fatalf(`event.Id is invalid %+v`, event)
	}
	if event.Id != events.Events[0].Id {
		t.Fatalf(`event.Id does not match events.Events[0].Id: %q`, err)
	}
}

func TestNewEventSearch(t *testing.T) {
	client := YelpClient{
		key: os.Getenv("YELP_API_KEY"),
	}
	eventsReq := client.NewEventSearch()
	res, err := eventsReq.Get()
	if err != nil {
		t.Fatalf(`Unexpected error: %q`, err)
	}
	if res.StatusCode == 0 {
		t.Fatalf(`Invalid response: %d`, res.StatusCode)
	}
	if res.StatusCode != 200 {
		t.Fatalf(`Unexpected response: %d`, res.StatusCode)
	}

	resbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf(`Unexpected error reading body: %q`, err)
	}
	defer res.Body.Close()

	events := Events{}
	err = json.Unmarshal(resbody, &events)
	if err != nil {
		t.Fatalf(`Unexpected error unmarshalling body: %q`, err)
	}
}

func TestNewFeaturedEvent(t *testing.T) {
	client := YelpClient{
		key: os.Getenv("YELP_API_KEY"),
	}
	eventsReq := client.NewFeaturedEvent()
	eventsReq.AddParameter("location", "Vancouver, Canada")
	res, err := eventsReq.Get()
	if err != nil {
		t.Fatalf(`Unexpected error: %q`, err)
	}
	if res.StatusCode == 0 {
		t.Fatalf(`Invalid response: %d`, res.StatusCode)
	}
	if res.StatusCode != 200 {
		t.Fatalf(`Unexpected response: %d`, res.StatusCode)
	}

	resbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf(`Unexpected error reading body: %q`, err)
	}
	defer res.Body.Close()

	event := Event{}
	err = json.Unmarshal(resbody, &event)
	if err != nil {
		t.Fatalf(`Unexpected error unmarshalling body: %q`, err)
	}
}
