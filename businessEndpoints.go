package yelp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	return &BusinessSearchRequest{*yc.NewRequest("/businesses/search")}
}

func (bsr *BusinessSearchRequest) Get() (*Businesses, error) {
	response, err := bsr.GetResponse()
	if err != nil {
		return &Businesses{}, err
	}

	responsebody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &Businesses{}, err
	}
	defer response.Body.Close()

	businesses := Businesses{}
	err = json.Unmarshal(responsebody, &businesses)
	if err != nil {
		return &Businesses{}, err
	}
	return &businesses, nil
}

func (bsr *BusinessSearchRequest) GetResponse() (*http.Response, error) {
	return bsr.YelpRequest.Get()
}

/*
YelpRequest for the Phone Search endpoint.
https://www.yelp.com/developers/documentation/v3/business_search_phone
*/
type PhoneSearchRequest struct {
	YelpRequest
}

func (yc *YelpClient) NewPhoneSearch() *PhoneSearchRequest {
	return &PhoneSearchRequest{*yc.NewRequest("/businesses/search/phone")}
}

func (psr *PhoneSearchRequest) Get() (*Businesses, error) {
	response, err := psr.GetResponse()
	if err != nil {
		return &Businesses{}, err
	}

	responsebody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &Businesses{}, err
	}
	defer response.Body.Close()

	businesses := Businesses{}
	err = json.Unmarshal(responsebody, &businesses)
	if err != nil {
		return &Businesses{}, err
	}
	return &businesses, nil
}

func (psr *PhoneSearchRequest) GetResponse() (*http.Response, error) {
	return psr.YelpRequest.Get()
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
		*yc.NewRequest(fmt.Sprintf("/transactions/%s/search", transactionType)),
	}
}

func (tsr *TransactionSearchRequest) Get() (*Businesses, error) {
	response, err := tsr.GetResponse()
	if err != nil {
		return &Businesses{}, err
	}

	responsebody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &Businesses{}, err
	}
	defer response.Body.Close()

	businesses := Businesses{}
	err = json.Unmarshal(responsebody, &businesses)
	if err != nil {
		return &Businesses{}, err
	}
	return &businesses, nil
}

func (tsr *TransactionSearchRequest) GetResponse() (*http.Response, error) {
	return tsr.YelpRequest.Get()
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
		*yc.NewRequest(fmt.Sprintf("/businesses/%s", businessId)),
	}
}

func (bdr *BusinessDetailsRequest) Get() (*Business, error) {
	response, err := bdr.GetResponse()
	if err != nil {
		return &Business{}, err
	}

	responsebody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &Business{}, err
	}
	defer response.Body.Close()

	business := Business{}
	err = json.Unmarshal(responsebody, &business)
	if err != nil {
		return &Business{}, err
	}
	return &business, nil
}

func (bdr *BusinessDetailsRequest) GetResponse() (*http.Response, error) {
	return bdr.YelpRequest.Get()
}

/*
YelpRequest for the Business Match endpoint.
https://www.yelp.com/developers/documentation/v3/business_match
*/
type BusinessMatchRequest struct {
	YelpRequest
}

func (yc *YelpClient) NewBusinessMatch() *BusinessMatchRequest {
	return &BusinessMatchRequest{*yc.NewRequest("/businesses/matches")}
}

func (bmr *BusinessMatchRequest) Get() (*Businesses, error) {
	response, err := bmr.GetResponse()
	if err != nil {
		return &Businesses{}, err
	}

	responsebody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &Businesses{}, err
	}
	defer response.Body.Close()

	businesses := Businesses{}
	err = json.Unmarshal(responsebody, &businesses)
	if err != nil {
		return &Businesses{}, err
	}
	return &businesses, nil
}

func (bmr *BusinessMatchRequest) GetResponse() (*http.Response, error) {
	return bmr.YelpRequest.Get()
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
		*yc.NewRequest(fmt.Sprintf("/businesses/%s/reviews", businessId)),
	}
}

func (rr *ReviewsRequest) Get() (*Reviews, error) {
	response, err := rr.GetResponse()
	if err != nil {
		return &Reviews{}, err
	}

	responsebody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &Reviews{}, err
	}
	defer response.Body.Close()

	reviews := Reviews{}
	err = json.Unmarshal(responsebody, &reviews)
	if err != nil {
		return &Reviews{}, err
	}
	return &reviews, nil
}

func (rr *ReviewsRequest) GetResponse() (*http.Response, error) {
	return rr.YelpRequest.Get()
}

/*
YelpRequest for the Autocomplete endpoint.
https://www.yelp.com/developers/documentation/v3/autocomplete
*/
type AutocompleteRequest struct {
	YelpRequest
}

func (yc *YelpClient) NewAutocomplete() *AutocompleteRequest {
	return &AutocompleteRequest{*yc.NewRequest("/autocomplete")}
}

func (acr *AutocompleteRequest) Get() (*Autocomplete, error) {
	response, err := acr.GetResponse()
	if err != nil {
		return &Autocomplete{}, err
	}

	responsebody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &Autocomplete{}, err
	}
	defer response.Body.Close()

	autocomplete := Autocomplete{}
	err = json.Unmarshal(responsebody, &autocomplete)
	if err != nil {
		return &Autocomplete{}, err
	}
	return &autocomplete, nil
}

func (acr *AutocompleteRequest) GetResponse() (*http.Response, error) {
	return acr.YelpRequest.Get()
}
