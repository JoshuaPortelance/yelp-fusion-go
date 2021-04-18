package yelp

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

// Testing YelpRequest for the All Categories endpoint and unmarshalling using the provided AllCategories type.
func TestNewAllCategories(t *testing.T) {
	var err error

	client := YelpClient{
		key: "h4lGIGYbwWvSf-TJVWUU2sp-WRTtUQb8n8N-UmCOSn9vF4Aa8LjtycdFqkwtcSArTcQgLlJLEf-T7KfSJKakKiRE5kmNldcjQ7sTK4bwefaewZfRorg-0n3v02ZEX3Yx",
	}
	allCatReq := client.NewAllCategories()
	res, err := allCatReq.Get()

	if err != nil {
		t.Fatalf(`Unexpected error: %q`, err)
	}
	if res.StatusCode == 0 {
		t.Fatalf(`Invalid response: %d`, res.StatusCode)
	}
	if res.StatusCode != 200 {
		t.Fatalf(`Unexpected response: %d`, res.StatusCode)
	}

	resbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf(`Unexpected error reading body: %q`, err)
	}
	defer res.Body.Close()

	allCat := AllCategories{}
	err = json.Unmarshal(resbody, &allCat)
	if err != nil {
		t.Fatalf(`Unexpected error unmarshalling body: %q`, err)
	}
}

// Testing YelpRequest for the Category Details endpoint and unmarshalling using the provided Category type.
func TestNewCategoryDetails(t *testing.T) {
	client := YelpClient{
		key: "h4lGIGYbwWvSf-TJVWUU2sp-WRTtUQb8n8N-UmCOSn9vF4Aa8LjtycdFqkwtcSArTcQgLlJLEf-T7KfSJKakKiRE5kmNldcjQ7sTK4bwefaewZfRorg-0n3v02ZEX3Yx",
	}
	catDetailsReq := client.NewCategoryDetails("arts")
	res, err := catDetailsReq.Get()

	if err != nil {
		t.Fatalf(`Unexpected error: %q`, err)
	}
	if res.StatusCode == 0 {
		t.Fatalf(`Invalid response: %d`, res.StatusCode)
	}
	if res.StatusCode != 200 {
		t.Fatalf(`Unexpected response: %d`, res.StatusCode)
	}

	resbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf(`Unexpected error reading body: %q`, err)
	}
	defer res.Body.Close()

	cat := Category{}
	err = json.Unmarshal(resbody, &cat)
	if err != nil {
		t.Fatalf(`Unexpected error unmarshalling body: %q`, err)
	}
}
