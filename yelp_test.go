package yelp

import (
	"fmt"
	"os"
	"testing"
)

func TestClientCreation(t *testing.T) {
	apiKey := "test"
	timeout := 5000

	client := Client{
		Key:     apiKey,
		Timeout: timeout,
	}
	if client.Key != apiKey {
		t.Fatalf(`Client.key = %q, expected %q`, client.Key, apiKey)
	}
	if client.Timeout != timeout {
		t.Fatalf(`Client.timeout = %q, expected %q`, client.Timeout, timeout)
	}
}

func TestClientDefaultCreation(t *testing.T) {
	apiKey := ""
	timeout := 0

	client := Client{}
	if client.Key != apiKey {
		t.Fatalf(`Client.key = %q, expected %q`, client.Key, apiKey)
	}
	if client.Timeout != timeout {
		t.Fatalf(`Client.timeout = %q, expected %q`, client.Timeout, timeout)
	}
}

func TestClientSingleFieldCreation(t *testing.T) {
	apiKey := "test"

	client := Client{Key: apiKey}
	if client.Key != apiKey {
		t.Fatalf(`Client.key = %q, expected %q`, client.Key, apiKey)
	}
}

func TestClientNewRequest(t *testing.T) {
	client := Client{}
	path := "/businesses/search"
	req := client.NewRequest(path, "BusinessSearch")

	if req.path != BaseYelpUrl+path {
		t.Fatalf(`Request.path = %q, expected %q`, req.path, path)
	}
}

func TestRequestAddParameter(t *testing.T) {
	client := Client{}
	req := client.NewRequest("", "")
	expected := "test"

	req.AddParameter(expected, expected)

	actual, found := req.params[expected]
	if !found {
		t.Fatalf(`Request.param key %q, is not in map`, expected)
	}
	if actual != expected {
		t.Fatalf(`Request.param[%q] = %q, expected %q`, expected, actual, expected)
	}
}

func TestRequestGetResponse(t *testing.T) {
	client := Client{
		Key: os.Getenv("YELP_API_KEY"),
	}
	req := client.NewRequest("/businesses/search", "BusinessSearch")

	req.AddParameter("location", "Victoria, Canada")

	res, err := req.GetResponse()
	if err != nil {
		t.Fatalf(`Request.Get error: %q`, err)
	}
	defer res.Body.Close()

	if res.StatusCode == 0 {
		t.Fatalf(`Request.Get Invalid status: %d`, res.StatusCode)
	}
}

func TestRequestGetResponseWithTimeout(t *testing.T) {
	client := Client{
		Key:     os.Getenv("YELP_API_KEY"),
		Timeout: 5000,
	}
	req := client.NewRequest("/businesses/search", "BusinessSearch")

	req.AddParameter("location", "Victoria, Canada")

	res, err := req.GetResponse()
	if err != nil {
		t.Fatalf(`Request.Get error: %q`, err)
	}
	defer res.Body.Close()

	if res.StatusCode == 0 {
		t.Fatalf(`Request.Get Invalid status: %d`, res.StatusCode)
	}
}

func TestRequestGetInvalidEndpoint(t *testing.T) {
	client := Client{
		Key:     os.Getenv("YELP_API_KEY"),
		Timeout: 30,
	}
	req := client.NewRequest("/businesses/search", "")

	req.AddParameter("location", "Victoria, Canada")

	_, err := req.Get()
	if err == nil {
		t.Fatalf(`Unexpected success`)
	}
}

func TestRequestGetResponseInvalidKey(t *testing.T) {
	client := Client{
		Key: "Invalid",
	}
	req := client.NewRequest("/businesses/search", "BusinessSearch")

	req.AddParameter("location", "Victoria, Canada")

	res, err := req.GetResponse()
	if err == nil {
		t.Fatalf(`Unexpected success`)
	}
	defer res.Body.Close()

	if res.StatusCode == 0 {
		t.Fatalf(`Invalid status: %d`, res.StatusCode)
	}
	if res.StatusCode != 400 {
		t.Fatalf(`Unexpected status: %d`, res.StatusCode)
	}
	if fmt.Sprint(res.StatusCode) != err.Error() {
		t.Fatalf(`Response code %d != error message code %s`, res.StatusCode, err)
	}
}

func TestRequestInvalidParams(t *testing.T) {
	client := Client{
		Key: os.Getenv("YELP_API_KEY"),
	}
	req := client.NewRequest("/businesses/search", "BusinessSearch")

	// Not adding any parameters. Yelp Fusion requests at least the
	// 'location' parameter to be given for the business search.

	res, err := req.GetResponse()
	if err == nil {
		t.Fatalf(`Unexpected success`)
	}
	defer res.Body.Close()

	if res.StatusCode == 0 {
		t.Fatalf(`Invalid status: %d`, res.StatusCode)
	}
	if res.StatusCode != 400 {
		t.Fatalf(`Unexpected status: %d`, res.StatusCode)
	}
	if fmt.Sprint(res.StatusCode) != err.Error() {
		t.Fatalf(`Response code %d != error message code %s`, res.StatusCode, err)
	}
}
