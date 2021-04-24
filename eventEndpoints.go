package yelp

import (
	"fmt"
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
		*yc.NewRequest(fmt.Sprintf("/events/%s", eventId), "EventLookup"),
	}
}

func (elr *EventLookupRequest) Get() (*Event, error) {
	data, err := elr.YelpRequest.Get()
	if err != nil {
		return nil, err
	}
	event := data.(Event)
	return &event, nil
}

func (elr *EventLookupRequest) GetResponse() (*http.Response, error) {
	return elr.YelpRequest.GetResponse()
}

// YelpRequest for the Event Search endpoint.
//
// https://www.yelp.com/developers/documentation/v3/event_search
type EventSearchRequest struct {
	YelpRequest
}

func (yc *YelpClient) NewEventSearch() *EventSearchRequest {
	return &EventSearchRequest{*yc.NewRequest("/events", "EventSearch")}
}

func (esr *EventSearchRequest) Get() (*Events, error) {
	data, err := esr.YelpRequest.Get()
	if err != nil {
		return nil, err
	}
	events := data.(Events)
	return &events, nil
}

func (esr *EventSearchRequest) GetResponse() (*http.Response, error) {
	return esr.YelpRequest.GetResponse()
}

// YelpRequest for the Featured Event endpoint.
//
// https://www.yelp.com/developers/documentation/v3/featured_event
type FeaturedEventRequest struct {
	YelpRequest
}

func (yc *YelpClient) NewFeaturedEvent() *FeaturedEventRequest {
	return &FeaturedEventRequest{*yc.NewRequest("/events/featured", "FeaturedEvent")}
}

func (fer *FeaturedEventRequest) Get() (*Event, error) {
	data, err := fer.YelpRequest.Get()
	if err != nil {
		return nil, err
	}
	event := data.(Event)
	return &event, nil
}

func (fer *FeaturedEventRequest) GetResponse() (*http.Response, error) {
	return fer.YelpRequest.GetResponse()
}
