package yelp

import (
	"fmt"
	"io"
	"os"
	"testing"
)

// Phone Search endpoint tests.

func TestNewPhoneSearchGet(t *testing.T) {
	client := YelpClient{
		key: os.Getenv("YELP_API_KEY"),
	}

	phoneReq := client.NewPhoneSearch()
	phoneReq.AddParameter("phone", "+12502986877")
	_, err := phoneReq.Get()
	if err != nil {
		t.Fatalf(`Unexpected error: %q`, err)
	}
}

func TestNewPhoneSearchGetResponse(t *testing.T) {
	client := YelpClient{
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
	client := YelpClient{
		key: os.Getenv("YELP_API_KEY"),
	}
	transactionReq := client.NewTransactionSearch("delivery")
	transactionReq.AddParameter("location", "Seattle, Washington")
	_, err := transactionReq.Get()
	if err != nil {
		t.Fatalf(`Unexpected error: %q`, err)
	}
}

func TestNewTransactionSearchGetResponse(t *testing.T) {
	client := YelpClient{
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
	client := YelpClient{
		key: os.Getenv("YELP_API_KEY"),
	}
	businessMatchReq := client.NewBusinessMatch()
	businessMatchReq.AddParameter("name", "Red Fish Blue Fish")
	businessMatchReq.AddParameter("address1", "1006 Wharf Street")
	businessMatchReq.AddParameter("city", "Victoria")
	businessMatchReq.AddParameter("state", "BC")
	businessMatchReq.AddParameter("country", "CA")
	_, err := businessMatchReq.Get()
	if err != nil {
		t.Fatalf(`Unexpected error: %q`, err)
	}
}

func TestNewBusinessMatchGetResponse(t *testing.T) {
	client := YelpClient{
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
	client := YelpClient{
		key: os.Getenv("YELP_API_KEY"),
	}
	autocompleteReq := client.NewAutocomplete()
	autocompleteReq.AddParameter("text", "Red Fish")
	autocompleteReq.AddParameter("latitude", "48.4243002")
	autocompleteReq.AddParameter("longitude", "-123.3706659")
	_, err := autocompleteReq.Get()
	if err != nil {
		t.Fatalf(`Unexpected error: %q`, err)
	}
}

func TestNewAutocompleteGetResponse(t *testing.T) {
	client := YelpClient{
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
	client := YelpClient{
		key: os.Getenv("YELP_API_KEY"),
	}
	businessesReq := client.NewBusinessSearch()
	businessesReq.AddParameter("location", "Victoria, Canada")
	businesses, businessesErr := businessesReq.Get()
	if businessesErr != nil {
		t.Fatalf(`Unexpected error: %q`, businessesErr)
	}

	businessReq := client.NewBusinessDetails(string(businesses.Businesses[0].Id))
	_, businessErr := businessReq.Get()
	if businessErr != nil {
		t.Fatalf(`Unexpected error: %q`, businessErr)
	}

	reviewsReq := client.NewReviews(string(businesses.Businesses[0].Id))
	_, reviewsErr := reviewsReq.Get()
	if reviewsErr != nil {
		t.Fatalf(`Unexpected error: %q`, reviewsErr)
	}
}
