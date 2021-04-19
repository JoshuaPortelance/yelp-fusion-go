package yelp

import "fmt"

// YelpRequest for the Business Search endpoint.
//
// https://www.yelp.com/developers/documentation/v3/business_search
func (c *YelpClient) NewBusinessSearch() *YelpRequest {
	return c.NewRequest("/businesses/search")
}

// YelpRequest for the Phone Search endpoint.
//
// https://www.yelp.com/developers/documentation/v3/business_search_phone
func (c *YelpClient) NewPhoneSearch() *YelpRequest {
	return c.NewRequest("/businesses/search/phone")
}

// YelpRequest for the Transaction Search endpoint.
//
// https://www.yelp.com/developers/documentation/v3/transaction_search
func (c *YelpClient) NewTransactionSearch(transactionType string) *YelpRequest {
	return c.NewRequest(fmt.Sprintf("/transactions/%s/search", transactionType))
}

// YelpRequest for the Business Details endpoint.
//
// https://www.yelp.com/developers/documentation/v3/business
func (c *YelpClient) NewBusinessDetails(businessId string) *YelpRequest {
	return c.NewRequest(fmt.Sprintf("/transactions/%s/search", businessId))
}

// YelpRequest for the Business Match endpoint.
//
// https://www.yelp.com/developers/documentation/v3/business_match
func (c *YelpClient) NewBusinessMatch() *YelpRequest {
	return c.NewRequest("/businesses/matches")
}

// YelpRequest for the Reviews endpoint.
//
// https://www.yelp.com/developers/documentation/v3/business_reviews
func (c *YelpClient) NewReviews(businessId string) *YelpRequest {
	return c.NewRequest(fmt.Sprintf("/businesses/%s/reviews", businessId))
}

// YelpRequest for the Autocomplete endpoint.
//
// https://www.yelp.com/developers/documentation/v3/autocomplete
func (c *YelpClient) NewAutocomplete() *YelpRequest {
	return c.NewRequest("/autocomplete")
}
