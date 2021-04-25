package yelp

import (
	"fmt"
)

/*
Request type and override(s) for the Event Lookup endpoint.
https://www.yelp.com/developers/documentation/v3/event
*/
type EventLookupRequest struct {
	Request
}

func (c *Client) NewEventLookup(eventId string) *EventLookupRequest {
	return &EventLookupRequest{
		*c.NewRequest(fmt.Sprintf("/events/%s", eventId), "EventLookup"),
	}
}

func (r *EventLookupRequest) Get() (*Event, error) {
	data, err := r.Request.Get()
	if err != nil {
		return nil, err
	}
	return &data.Event, nil
}

/*
Request type and override(s) for the Event Search endpoint.
https://www.yelp.com/developers/documentation/v3/event_search
*/
type EventSearchRequest struct {
	Request
}

func (c *Client) NewEventSearch() *EventSearchRequest {
	return &EventSearchRequest{*c.NewRequest("/events", "EventSearch")}
}

func (r *EventSearchRequest) Get() (*Events, error) {
	data, err := r.Request.Get()
	if err != nil {
		return nil, err
	}
	return &data.Events, nil
}

/*
Request type and override(s) for the Featured Event endpoint.
https://www.yelp.com/developers/documentation/v3/featured_event
*/
type FeaturedEventRequest struct {
	Request
}

func (c *Client) NewFeaturedEvent() *FeaturedEventRequest {
	return &FeaturedEventRequest{*c.NewRequest("/events/featured", "FeaturedEvent")}
}

func (r *FeaturedEventRequest) Get() (*Event, error) {
	data, err := r.Request.Get()
	if err != nil {
		return nil, err
	}
	return &data.Event, nil
}
