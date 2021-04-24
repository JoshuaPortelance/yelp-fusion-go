package yelp

import (
	"fmt"
	"net/http"
)

/*
Request for the Event Lookup endpoint.
https://www.yelp.com/developers/documentation/v3/event
*/
type EventLookupRequest struct {
	Request
}

func (yc *Client) NewEventLookup(eventId string) *EventLookupRequest {
	return &EventLookupRequest{
		*yc.NewRequest(fmt.Sprintf("/events/%s", eventId), "EventLookup"),
	}
}

func (elr *EventLookupRequest) Get() (*Event, error) {
	data, err := elr.Request.Get()
	if err != nil {
		return nil, err
	}
	return &data.Event, nil
}

func (elr *EventLookupRequest) GetResponse() (*http.Response, error) {
	return elr.Request.GetResponse()
}

// Request for the Event Search endpoint.
//
// https://www.yelp.com/developers/documentation/v3/event_search
type EventSearchRequest struct {
	Request
}

func (yc *Client) NewEventSearch() *EventSearchRequest {
	return &EventSearchRequest{*yc.NewRequest("/events", "EventSearch")}
}

func (esr *EventSearchRequest) Get() (*Events, error) {
	data, err := esr.Request.Get()
	if err != nil {
		return nil, err
	}
	return &data.Events, nil
}

func (esr *EventSearchRequest) GetResponse() (*http.Response, error) {
	return esr.Request.GetResponse()
}

// Request for the Featured Event endpoint.
//
// https://www.yelp.com/developers/documentation/v3/featured_event
type FeaturedEventRequest struct {
	Request
}

func (yc *Client) NewFeaturedEvent() *FeaturedEventRequest {
	return &FeaturedEventRequest{*yc.NewRequest("/events/featured", "FeaturedEvent")}
}

func (fer *FeaturedEventRequest) Get() (*Event, error) {
	data, err := fer.Request.Get()
	if err != nil {
		return nil, err
	}
	return &data.Event, nil
}

func (fer *FeaturedEventRequest) GetResponse() (*http.Response, error) {
	return fer.Request.GetResponse()
}
