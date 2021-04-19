package yelp

import "fmt"

// YelpRequest for the Event Lookup endpoint.
func (c *YelpClient) NewEventLookup(eventId string) *YelpRequest {
	return c.NewRequest(fmt.Sprintf("/events/%s", eventId))
}

// YelpRequest for the Event Search endpoint.
func (c *YelpClient) NewEventSearch() *YelpRequest {
	return c.NewRequest("/events")
}

// YelpRequest for the Featured Event endpoint.
func (c *YelpClient) NewFeaturedEvent() *YelpRequest {
	return c.NewRequest("/events/featured")
}
