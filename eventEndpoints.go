package yelp

import "fmt"

// YelpRequest for the Event Lookup endpoint.
//
// https://www.yelp.com/developers/documentation/v3/event
func (c *YelpClient) NewEventLookup(eventId string) *YelpRequest {
	return c.NewRequest(fmt.Sprintf("/events/%s", eventId))
}

// YelpRequest for the Event Search endpoint.
//
// https://www.yelp.com/developers/documentation/v3/event_search
func (c *YelpClient) NewEventSearch() *YelpRequest {
	return c.NewRequest("/events")
}

// YelpRequest for the Featured Event endpoint.
//
// https://www.yelp.com/developers/documentation/v3/featured_event
func (c *YelpClient) NewFeaturedEvent() *YelpRequest {
	return c.NewRequest("/events/featured")
}
