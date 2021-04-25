package yelp

import (
	"fmt"
)

/*
Request type and override(s) for the Business Search endpoint.
https://www.yelp.com/developers/documentation/v3/business_search
*/
type BusinessSearchRequest struct {
	Request
}

func (c *Client) NewBusinessSearch() *BusinessSearchRequest {
	return &BusinessSearchRequest{*c.NewRequest("/businesses/search", "BusinessSearch")}
}

func (r *BusinessSearchRequest) Get() (*Businesses, error) {
	data, err := r.Request.Get()
	if err != nil {
		return nil, err
	}
	return &data.Businesses, nil
}

/*
Request type and override(s) for the Phone Search endpoint.
https://www.yelp.com/developers/documentation/v3/business_search_phone
*/
type PhoneSearchRequest struct {
	Request
}

func (c *Client) NewPhoneSearch() *PhoneSearchRequest {
	return &PhoneSearchRequest{*c.NewRequest("/businesses/search/phone", "PhoneSearch")}
}

func (r *PhoneSearchRequest) Get() (*Businesses, error) {
	data, err := r.Request.Get()
	if err != nil {
		return nil, err
	}
	return &data.Businesses, nil
}

/*
Request type and override(s) for the Transaction Search endpoint.
https://www.yelp.com/developers/documentation/v3/transaction_search
*/
type TransactionSearchRequest struct {
	Request
}

func (c *Client) NewTransactionSearch(transactionType string) *TransactionSearchRequest {
	return &TransactionSearchRequest{
		*c.NewRequest(fmt.Sprintf("/transactions/%s/search", transactionType), "TransactionSearch"),
	}
}

func (r *TransactionSearchRequest) Get() (*Businesses, error) {
	data, err := r.Request.Get()
	if err != nil {
		return nil, err
	}
	return &data.Businesses, nil
}

/*
Request type and override(s) for the Business Details endpoint.
https://www.yelp.com/developers/documentation/v3/business
*/
type BusinessDetailsRequest struct {
	Request
}

func (c *Client) NewBusinessDetails(businessId string) *BusinessDetailsRequest {
	return &BusinessDetailsRequest{
		*c.NewRequest(fmt.Sprintf("/businesses/%s", businessId), "BusinessDetails"),
	}
}

func (r *BusinessDetailsRequest) Get() (*Business, error) {
	data, err := r.Request.Get()
	if err != nil {
		return nil, err
	}
	return &data.Business, nil
}

/*
Request type and override(s) for the Business Match endpoint.
https://www.yelp.com/developers/documentation/v3/business_match
*/
type BusinessMatchRequest struct {
	Request
}

func (c *Client) NewBusinessMatch() *BusinessMatchRequest {
	return &BusinessMatchRequest{*c.NewRequest("/businesses/matches", "BusinessMatch")}
}

func (r *BusinessMatchRequest) Get() (*Businesses, error) {
	data, err := r.Request.Get()
	if err != nil {
		return nil, err
	}
	return &data.Businesses, nil
}

/*
Request type and override(s) for the Reviews endpoint.
https://www.yelp.com/developers/documentation/v3/business_reviews
*/
type ReviewsRequest struct {
	Request
}

func (c *Client) NewReviews(businessId string) *ReviewsRequest {
	return &ReviewsRequest{
		*c.NewRequest(fmt.Sprintf("/businesses/%s/reviews", businessId), "Reviews"),
	}
}

func (r *ReviewsRequest) Get() (*Reviews, error) {
	data, err := r.Request.Get()
	if err != nil {
		return nil, err
	}
	return &data.Reviews, nil
}

/*
Request type and override(s) for the Autocomplete endpoint.
https://www.yelp.com/developers/documentation/v3/autocomplete
*/
type AutocompleteRequest struct {
	Request
}

func (c *Client) NewAutocomplete() *AutocompleteRequest {
	return &AutocompleteRequest{*c.NewRequest("/autocomplete", "Autocomplete")}
}

func (r *AutocompleteRequest) Get() (*Autocomplete, error) {
	data, err := r.Request.Get()
	if err != nil {
		return nil, err
	}
	return &data.Autocomplete, nil
}
