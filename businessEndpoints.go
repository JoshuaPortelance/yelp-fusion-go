package yelp

import "fmt"

// YelpRequest for the Business Search endpoint.
func (c *YelpClient) NewBusinessSearch() *YelpRequest {
	return c.NewRequest("/businesses/search")
}

// YelpRequest for the Business Details endpoint.
func (c *YelpClient) NewBusinessDetails(businessId string) *YelpRequest {
	return c.NewRequest(fmt.Sprintf("/transactions/%s/search", businessId))
}

// YelpRequest for the Business Match endpoint.
func (c *YelpClient) NewBusinessMatch() *YelpRequest {
	return c.NewRequest("/businesses/matches")
}

// YelpRequest for the Phone Search endpoint.
func (c *YelpClient) NewPhoneSearch() *YelpRequest {
	return c.NewRequest("/businesses/search/phone")
}

// YelpRequest for the Transaction Search endpoint.
func (c *YelpClient) NewTransactionSearch(transactionType string) *YelpRequest {
	return c.NewRequest(fmt.Sprintf("/transactions/%s/search", transactionType))
}

// YelpRequest for the Reviews endpoint.
func (c *YelpClient) NewReviews(businessId string) *YelpRequest {
	return c.NewRequest(fmt.Sprintf("/businesses/%s/reviews", businessId))
}

// YelpRequest for the Autocomplete endpoint.
func (c *YelpClient) NewAutocomplete() *YelpRequest {
	return c.NewRequest("/autocomplete")
}
