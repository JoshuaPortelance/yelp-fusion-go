package yelp

import (
	"errors"
	"fmt"
	"net/http"
)

/*
YelpRequest for the Business Search endpoint.
https://www.yelp.com/developers/documentation/v3/business_search
*/
type BusinessSearchRequest struct {
	YelpRequest
}

func (yc *YelpClient) NewBusinessSearch() *BusinessSearchRequest {
	return &BusinessSearchRequest{*yc.NewRequest("/businesses/search", "BusinessSearch")}
}

func (bsr *BusinessSearchRequest) Get() (*Businesses, error) {
	data, err := bsr.YelpRequest.Get()
	if err != nil {
		return nil, err
	}
	businesses, ok := data.(Businesses)
	if !ok {
		return nil, errors.New("failed to convert interface to Businesses")
	}
	return &businesses, nil
}

func (bsr *BusinessSearchRequest) GetResponse() (*http.Response, error) {
	return bsr.YelpRequest.GetResponse()
}

/*
YelpRequest for the Phone Search endpoint.
https://www.yelp.com/developers/documentation/v3/business_search_phone
*/
type PhoneSearchRequest struct {
	YelpRequest
}

func (yc *YelpClient) NewPhoneSearch() *PhoneSearchRequest {
	return &PhoneSearchRequest{*yc.NewRequest("/businesses/search/phone", "PhoneSearch")}
}

func (psr *PhoneSearchRequest) Get() (*Businesses, error) {
	data, err := psr.YelpRequest.Get()
	if err != nil {
		return nil, err
	}
	businesses, ok := data.(Businesses)
	if !ok {
		return nil, errors.New("failed to convert interface to Businesses")
	}
	return &businesses, nil
}

func (psr *PhoneSearchRequest) GetResponse() (*http.Response, error) {
	return psr.YelpRequest.GetResponse()
}

/*
YelpRequest for the Transaction Search endpoint.
https://www.yelp.com/developers/documentation/v3/transaction_search
*/
type TransactionSearchRequest struct {
	YelpRequest
}

func (yc *YelpClient) NewTransactionSearch(transactionType string) *TransactionSearchRequest {
	return &TransactionSearchRequest{
		*yc.NewRequest(fmt.Sprintf("/transactions/%s/search", transactionType), "TransactionSearch"),
	}
}

func (tsr *TransactionSearchRequest) Get() (*Businesses, error) {
	data, err := tsr.YelpRequest.Get()
	if err != nil {
		return nil, err
	}
	businesses, ok := data.(Businesses)
	if !ok {
		return nil, errors.New("failed to convert interface to Businesses")
	}
	return &businesses, nil
}

func (tsr *TransactionSearchRequest) GetResponse() (*http.Response, error) {
	return tsr.YelpRequest.GetResponse()
}

/*
YelpRequest for the Business Details endpoint.
https://www.yelp.com/developers/documentation/v3/business
*/
type BusinessDetailsRequest struct {
	YelpRequest
}

func (yc *YelpClient) NewBusinessDetails(businessId string) *BusinessDetailsRequest {
	return &BusinessDetailsRequest{
		*yc.NewRequest(fmt.Sprintf("/businesses/%s", businessId), "BusinessDetails"),
	}
}

func (bdr *BusinessDetailsRequest) Get() (*Business, error) {
	data, err := bdr.YelpRequest.Get()
	if err != nil {
		return nil, err
	}
	business, ok := data.(Business)
	if !ok {
		return nil, errors.New("failed to convert interface to Business")
	}
	return &business, nil
}

func (bdr *BusinessDetailsRequest) GetResponse() (*http.Response, error) {
	return bdr.YelpRequest.GetResponse()
}

/*
YelpRequest for the Business Match endpoint.
https://www.yelp.com/developers/documentation/v3/business_match
*/
type BusinessMatchRequest struct {
	YelpRequest
}

func (yc *YelpClient) NewBusinessMatch() *BusinessMatchRequest {
	return &BusinessMatchRequest{*yc.NewRequest("/businesses/matches", "BusinessMatch")}
}

func (bmr *BusinessMatchRequest) Get() (*Businesses, error) {
	data, err := bmr.YelpRequest.Get()
	if err != nil {
		return nil, err
	}
	businesses, ok := data.(Businesses)
	if !ok {
		return nil, errors.New("failed to convert interface to Businesses")
	}
	return &businesses, nil
}

func (bmr *BusinessMatchRequest) GetResponse() (*http.Response, error) {
	return bmr.YelpRequest.GetResponse()
}

/*
YelpRequest for the Reviews endpoint.
https://www.yelp.com/developers/documentation/v3/business_reviews
*/
type ReviewsRequest struct {
	YelpRequest
}

func (yc *YelpClient) NewReviews(businessId string) *ReviewsRequest {
	return &ReviewsRequest{
		*yc.NewRequest(fmt.Sprintf("/businesses/%s/reviews", businessId), "Reviews"),
	}
}

func (rr *ReviewsRequest) Get() (*Reviews, error) {
	data, err := rr.YelpRequest.Get()
	if err != nil {
		return nil, err
	}
	reviews, ok := data.(Reviews)
	if !ok {
		return nil, errors.New("failed to convert interface to Reviews")
	}
	return &reviews, nil
}

func (rr *ReviewsRequest) GetResponse() (*http.Response, error) {
	return rr.YelpRequest.GetResponse()
}

/*
YelpRequest for the Autocomplete endpoint.
https://www.yelp.com/developers/documentation/v3/autocomplete
*/
type AutocompleteRequest struct {
	YelpRequest
}

func (yc *YelpClient) NewAutocomplete() *AutocompleteRequest {
	return &AutocompleteRequest{*yc.NewRequest("/autocomplete", "Autocomplete")}
}

func (acr *AutocompleteRequest) Get() (*Autocomplete, error) {
	data, err := acr.YelpRequest.Get()
	if err != nil {
		return nil, err
	}
	autocomplete, ok := data.(Autocomplete)
	if !ok {
		return nil, errors.New("failed to convert interface to Autocomplete")
	}
	return &autocomplete, nil
}

func (acr *AutocompleteRequest) GetResponse() (*http.Response, error) {
	return acr.YelpRequest.GetResponse()
}
