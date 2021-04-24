package yelp

import (
	"fmt"
)

/*
Request for the Business Search endpoint.
https://www.yelp.com/developers/documentation/v3/business_search
*/
type BusinessSearchRequest struct {
	Request
}

func (yc *Client) NewBusinessSearch() *BusinessSearchRequest {
	return &BusinessSearchRequest{*yc.NewRequest("/businesses/search", "BusinessSearch")}
}

func (bsr *BusinessSearchRequest) Get() (*Businesses, error) {
	data, err := bsr.Request.Get()
	if err != nil {
		return nil, err
	}
	return &data.Businesses, nil
}

/*
Request for the Phone Search endpoint.
https://www.yelp.com/developers/documentation/v3/business_search_phone
*/
type PhoneSearchRequest struct {
	Request
}

func (yc *Client) NewPhoneSearch() *PhoneSearchRequest {
	return &PhoneSearchRequest{*yc.NewRequest("/businesses/search/phone", "PhoneSearch")}
}

func (psr *PhoneSearchRequest) Get() (*Businesses, error) {
	data, err := psr.Request.Get()
	if err != nil {
		return nil, err
	}
	return &data.Businesses, nil
}

/*
Request for the Transaction Search endpoint.
https://www.yelp.com/developers/documentation/v3/transaction_search
*/
type TransactionSearchRequest struct {
	Request
}

func (yc *Client) NewTransactionSearch(transactionType string) *TransactionSearchRequest {
	return &TransactionSearchRequest{
		*yc.NewRequest(fmt.Sprintf("/transactions/%s/search", transactionType), "TransactionSearch"),
	}
}

func (tsr *TransactionSearchRequest) Get() (*Businesses, error) {
	data, err := tsr.Request.Get()
	if err != nil {
		return nil, err
	}
	return &data.Businesses, nil
}

/*
Request for the Business Details endpoint.
https://www.yelp.com/developers/documentation/v3/business
*/
type BusinessDetailsRequest struct {
	Request
}

func (yc *Client) NewBusinessDetails(businessId string) *BusinessDetailsRequest {
	return &BusinessDetailsRequest{
		*yc.NewRequest(fmt.Sprintf("/businesses/%s", businessId), "BusinessDetails"),
	}
}

func (bdr *BusinessDetailsRequest) Get() (*Business, error) {
	data, err := bdr.Request.Get()
	if err != nil {
		return nil, err
	}
	return &data.Business, nil
}

/*
Request for the Business Match endpoint.
https://www.yelp.com/developers/documentation/v3/business_match
*/
type BusinessMatchRequest struct {
	Request
}

func (yc *Client) NewBusinessMatch() *BusinessMatchRequest {
	return &BusinessMatchRequest{*yc.NewRequest("/businesses/matches", "BusinessMatch")}
}

func (bmr *BusinessMatchRequest) Get() (*Businesses, error) {
	data, err := bmr.Request.Get()
	if err != nil {
		return nil, err
	}
	return &data.Businesses, nil
}

/*
Request for the Reviews endpoint.
https://www.yelp.com/developers/documentation/v3/business_reviews
*/
type ReviewsRequest struct {
	Request
}

func (yc *Client) NewReviews(businessId string) *ReviewsRequest {
	return &ReviewsRequest{
		*yc.NewRequest(fmt.Sprintf("/businesses/%s/reviews", businessId), "Reviews"),
	}
}

func (rr *ReviewsRequest) Get() (*Reviews, error) {
	data, err := rr.Request.Get()
	if err != nil {
		return nil, err
	}
	return &data.Reviews, nil
}

/*
Request for the Autocomplete endpoint.
https://www.yelp.com/developers/documentation/v3/autocomplete
*/
type AutocompleteRequest struct {
	Request
}

func (yc *Client) NewAutocomplete() *AutocompleteRequest {
	return &AutocompleteRequest{*yc.NewRequest("/autocomplete", "Autocomplete")}
}

func (acr *AutocompleteRequest) Get() (*Autocomplete, error) {
	data, err := acr.Request.Get()
	if err != nil {
		return nil, err
	}
	return &data.Autocomplete, nil
}
