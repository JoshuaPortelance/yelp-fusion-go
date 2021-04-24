package yelp

import (
	"fmt"
	"io"
	"os"
	"testing"
)

// Phone Search endpoint tests.

func TestNewPhoneSearchGet(t *testing.T) {
	client := Client{
		key: os.Getenv("YELP_API_KEY"),
	}

	phoneReq := client.NewPhoneSearch()
	phoneReq.AddParameter("phone", "+12502986877")
	data, err := phoneReq.Get()
	if err != nil {
		t.Fatalf(`Unexpected error: %q`, err)
	}
	if data.Total == 0 {
		t.Fatalf(`Didn't find any businesses. %v+`, data)
	}
	if data.Businesses[0].Name == "" {
		t.Fatalf(`Business has no name. %v+`, data)
	}
}

func TestNewPhoneSearchGetResponse(t *testing.T) {
	client := Client{
		key: os.Getenv("YELP_API_KEY"),
	}
	phoneReq := client.NewPhoneSearch()
	phoneReq.AddParameter("phone", "+12502986877")
	response, err := phoneReq.GetResponse()
	if err != nil {
		body, _ := io.ReadAll(response.Body)
		fmt.Println(string(body))
		t.Fatalf(`Unexpected error: %q`, err)
	}
	defer response.Body.Close()
	if response.StatusCode == 0 {
		t.Fatalf(`Invalid status: %d`, response.StatusCode)
	}
	if response.StatusCode != 200 {
		t.Fatalf(`Unexpected status: %d`, response.StatusCode)
	}
}

// Transaction Search endpoint tests.

func TestNewTransactionSearchGet(t *testing.T) {
	client := Client{
		key: os.Getenv("YELP_API_KEY"),
	}
	transactionReq := client.NewTransactionSearch("delivery")
	transactionReq.AddParameter("location", "Seattle, Washington")
	data, err := transactionReq.Get()
	if err != nil {
		t.Fatalf(`Unexpected error: %q`, err)
	}
	if data.Total == 0 {
		t.Fatalf(`Didn't find any businesses. %v+`, data)
	}
	if data.Businesses[0].Name == "" {
		t.Fatalf(`Business has no name. %v+`, data)
	}
}

func TestNewTransactionSearchGetResponse(t *testing.T) {
	client := Client{
		key: os.Getenv("YELP_API_KEY"),
	}
	transactionReq := client.NewTransactionSearch("delivery")
	transactionReq.AddParameter("location", "Seattle, Washington")
	response, err := transactionReq.GetResponse()
	if err != nil {
		body, _ := io.ReadAll(response.Body)
		fmt.Println(string(body))
		t.Fatalf(`Unexpected error: %q`, err)
	}
	defer response.Body.Close()
	if response.StatusCode == 0 {
		t.Fatalf(`Invalid status: %d`, response.StatusCode)
	}
	if response.StatusCode != 200 {
		t.Fatalf(`Unexpected status: %d`, response.StatusCode)
	}
}

// Business Match endpoint tests.

func TestNewBusinessMatchGet(t *testing.T) {
	client := Client{
		key: os.Getenv("YELP_API_KEY"),
	}
	businessMatchReq := client.NewBusinessMatch()
	businessMatchReq.AddParameter("name", "Red Fish Blue Fish")
	businessMatchReq.AddParameter("address1", "1006 Wharf Street")
	businessMatchReq.AddParameter("city", "Victoria")
	businessMatchReq.AddParameter("state", "BC")
	businessMatchReq.AddParameter("country", "CA")
	data, err := businessMatchReq.Get()
	if err != nil {
		t.Fatalf(`Unexpected error: %q`, err)
	}
	if len(data.Businesses) == 0 {
		t.Fatalf(`Didn't find any businesses. %v+`, data)
	}
	if data.Businesses[0].Name == "" {
		t.Fatalf(`Business has no name. %v+`, data)
	}
}

func TestNewBusinessMatchGetResponse(t *testing.T) {
	client := Client{
		key: os.Getenv("YELP_API_KEY"),
	}
	businessMatchReq := client.NewBusinessMatch()
	businessMatchReq.AddParameter("name", "Red Fish Blue Fish")
	businessMatchReq.AddParameter("address1", "1006 Wharf Street")
	businessMatchReq.AddParameter("city", "Victoria")
	businessMatchReq.AddParameter("state", "BC")
	businessMatchReq.AddParameter("country", "CA")
	response, err := businessMatchReq.GetResponse()
	if err != nil {
		body, _ := io.ReadAll(response.Body)
		fmt.Println(string(body))
		t.Fatalf(`Unexpected error: %q`, err)
	}
	defer response.Body.Close()
	if response.StatusCode == 0 {
		t.Fatalf(`Invalid status: %d`, response.StatusCode)
	}
	if response.StatusCode != 200 {
		t.Fatalf(`Unexpected status: %d`, response.StatusCode)
	}
}

// Business Autocomplete endpoint tests.

func TestNewAutocompleteGet(t *testing.T) {
	client := Client{
		key: os.Getenv("YELP_API_KEY"),
	}
	autocompleteReq := client.NewAutocomplete()
	autocompleteReq.AddParameter("text", "Red Fish")
	autocompleteReq.AddParameter("latitude", "48.4243002")
	autocompleteReq.AddParameter("longitude", "-123.3706659")
	data, err := autocompleteReq.Get()
	if err != nil {
		t.Fatalf(`Unexpected error: %q`, err)
	}
	if len(data.Terms) == 0 {
		t.Fatalf(`Didn't autocomplete to any terms. %v+`, data)
	}
	if data.Terms[0].Text == "" {
		t.Fatalf(`Invalid term. %v+`, data)
	}
}

func TestNewAutocompleteGetResponse(t *testing.T) {
	client := Client{
		key: os.Getenv("YELP_API_KEY"),
	}
	autocompleteReq := client.NewAutocomplete()
	autocompleteReq.AddParameter("text", "Red Fish")
	autocompleteReq.AddParameter("latitude", "48.4243002")
	autocompleteReq.AddParameter("longitude", "-123.3706659")
	response, err := autocompleteReq.GetResponse()
	if err != nil {
		body, _ := io.ReadAll(response.Body)
		fmt.Println(string(body))
		t.Fatalf(`Unexpected error: %q`, err)
	}
	defer response.Body.Close()
	if response.StatusCode == 0 {
		t.Fatalf(`Invalid status: %d`, response.StatusCode)
	}
	if response.StatusCode != 200 {
		t.Fatalf(`Unexpected status: %d`, response.StatusCode)
	}
}

// Business Search endpoint tests.
// Business Details endpoint tests.
// Business Reviews endpoint tests.

func TestNewBusinessSearchAndDetailsAndReviewsGet(t *testing.T) {
	client := Client{
		key: os.Getenv("YELP_API_KEY"),
	}
	businessesReq := client.NewBusinessSearch()
	businessesReq.AddParameter("location", "Victoria, Canada")
	businesses, businessesErr := businessesReq.Get()
	if businessesErr != nil {
		t.Fatalf(`Unexpected error: %q`, businessesErr)
	}
	if businesses.Total == 0 {
		t.Fatalf(`Didn't find any businesses. %v+`, businesses)
	}
	if businesses.Businesses[0].Name == "" {
		t.Fatalf(`Business has no name. %v+`, businesses)
	}

	businessReq := client.NewBusinessDetails(string(businesses.Businesses[0].Id))
	business, businessErr := businessReq.Get()
	if businessErr != nil {
		t.Fatalf(`Unexpected error: %q`, businessErr)
	}
	if business.Name == "" {
		t.Fatalf(`Business has no name. %v+`, business)
	}

	reviewsReq := client.NewReviews(string(businesses.Businesses[0].Id))
	reviews, reviewsErr := reviewsReq.Get()
	if reviewsErr != nil {
		t.Fatalf(`Unexpected error: %q`, reviewsErr)
	}
	if reviews.Total == 0 {
		t.Fatalf(`Found no reviews. %v+`, reviews)
	}
	if reviews.Reviews[0].Text == "" {
		t.Fatalf(`Invalid review. %v+`, reviews)
	}
}
