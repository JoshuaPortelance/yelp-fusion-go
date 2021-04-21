package yelp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
YelpRequest for the Event Lookup endpoint.
https://www.yelp.com/developers/documentation/v3/event
*/
type EventLookupRequest struct {
	YelpRequest
}

func (yc *YelpClient) NewEventLookup(eventId string) *EventLookupRequest {
	return &EventLookupRequest{
		*yc.NewRequest(fmt.Sprintf("/events/%s", eventId)),
	}
}

func (elr *EventLookupRequest) Get() (*Event, error) {
	response, err := elr.GetResponse()
	if err != nil {
		return &Event{}, err
	}

	responsebody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &Event{}, err
	}
	defer response.Body.Close()

	event := Event{}
	err = json.Unmarshal(responsebody, &event)
	if err != nil {
		return &Event{}, err
	}
	return &event, nil
}

func (elr *EventLookupRequest) GetResponse() (*http.Response, error) {
	return elr.YelpRequest.Get()
}

// YelpRequest for the Event Search endpoint.
//
// https://www.yelp.com/developers/documentation/v3/event_search
type EventSearchRequest struct {
	YelpRequest
}

func (yc *YelpClient) NewEventSearch() *EventSearchRequest {
	return &EventSearchRequest{*yc.NewRequest("/events")}
}

func (esr *EventSearchRequest) Get() (*Events, error) {
	response, err := esr.GetResponse()
	if err != nil {
		return &Events{}, err
	}

	responsebody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &Events{}, err
	}
	defer response.Body.Close()

	events := Events{}
	err = json.Unmarshal(responsebody, &events)
	if err != nil {
		return &Events{}, err
	}
	return &events, nil
}

func (esr *EventSearchRequest) GetResponse() (*http.Response, error) {
	return esr.YelpRequest.Get()
}

// YelpRequest for the Featured Event endpoint.
//
// https://www.yelp.com/developers/documentation/v3/featured_event
type FeaturedEventRequest struct {
	YelpRequest
}

func (yc *YelpClient) NewFeaturedEvent() *FeaturedEventRequest {
	return &FeaturedEventRequest{*yc.NewRequest("/events/featured")}
}

func (fer *FeaturedEventRequest) Get() (*Event, error) {
	response, err := fer.GetResponse()
	if err != nil {
		return &Event{}, err
	}

	responsebody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &Event{}, err
	}
	defer response.Body.Close()

	event := Event{}
	err = json.Unmarshal(responsebody, &event)
	if err != nil {
		return &Event{}, err
	}
	return &event, nil
}

func (fer *FeaturedEventRequest) GetResponse() (*http.Response, error) {
	return fer.YelpRequest.Get()
}
