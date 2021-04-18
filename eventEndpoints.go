package yelp

import "fmt"

type Location struct {
	Address1       string   `json:"address1"`
	Address2       string   `json:"address2"`
	Address3       string   `json:"address3"`
	City           string   `json:"city"`
	ZipCode        string   `json:"zip_code"`
	Country        string   `json:"country"`
	State          string   `json:"state"`
	DisplayAddress []string `json:"display_address"`
	Cross_streets  string   `json:"cross_streets"`
}

type Event struct {
	Attending_count int      `json:"attending_count"`
	Category        string   `json:"category"`
	Cost            float32  `json:"cost"`
	CostMax         float32  `json:"cost_max"`
	Description     string   `json:"description"`
	EventSiteUrl    string   `json:"event_site_url"`
	Id              string   `json:"id"`
	ImageUrl        string   `json:"image_url"`
	InterestedCount int      `json:"interested_count"`
	IsCanceled      bool     `json:"is_canceled"`
	IsFree          bool     `json:"is_free"`
	IsOfficial      bool     `json:"is_official"`
	Latitude        float32  `json:"latitude"`
	Longitude       float32  `json:"longitude"`
	Name            string   `json:"name"`
	TicketsUrl      string   `json:"tickets_url"`
	TimeEnd         string   `json:"time_end"`
	TimeStart       string   `json:"time_start"`
	Location        Location `json:"location"`
	BusinessId      string   `json:"business_id"`
}

type Events struct {
	Total  int     `json:"total"`
	Events []Event `json:"events"`
}

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
