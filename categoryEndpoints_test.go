package yelp

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestNewAllCategoriesGet(t *testing.T) {
	client := YelpClient{
		key: os.Getenv("YELP_API_KEY"),
	}
	allCatReq := client.NewAllCategories()
	_, err := allCatReq.Get()
	if err != nil {
		t.Fatalf(`Unexpected error: %q`, err)
	}
}

func TestNewAllCategoriesGetResponse(t *testing.T) {
	client := YelpClient{
		key: os.Getenv("YELP_API_KEY"),
	}
	allCatReq := client.NewAllCategories()
	response, err := allCatReq.GetResponse()
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

func TestNewCategoryDetailsGet(t *testing.T) {
	client := YelpClient{
		key: os.Getenv("YELP_API_KEY"),
	}
	catDetailsReq := client.NewCategoryDetails("arts")
	_, err := catDetailsReq.Get()
	if err != nil {
		t.Fatalf(`Unexpected error: %q`, err)
	}
}

func TestNewCategoryDetailsGetResponse(t *testing.T) {
	client := YelpClient{
		key: os.Getenv("YELP_API_KEY"),
	}
	catDetailsReq := client.NewAllCategories()
	response, err := catDetailsReq.GetResponse()
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
