package yelp

import (
	"testing"
)

func TestYelpClientCreation(t *testing.T) {
	apiKey := "test"
	timeout := 5000

	client := YelpClient{
		key:     apiKey,
		timeout: timeout,
	}
	if client.key != apiKey {
		t.Fatalf(`YelpClient.key = %q, expected %q`, client.key, apiKey)
	}
	if client.timeout != timeout {
		t.Fatalf(`YelpClient.timeout = %q, expected %q`, client.timeout, timeout)
	}
}

func TestYelpClientDefaultCreation(t *testing.T) {
	apiKey := ""
	timeout := 0

	client := YelpClient{}
	if client.key != apiKey {
		t.Fatalf(`YelpClient.key = %q, expected %q`, client.key, apiKey)
	}
	if client.timeout != timeout {
		t.Fatalf(`YelpClient.timeout = %q, expected %q`, client.timeout, timeout)
	}
}

func TestYelpClientSingleFieldCreation(t *testing.T) {
	apiKey := "test"

	client := YelpClient{key: apiKey}
	if client.key != apiKey {
		t.Fatalf(`YelpClient.key = %q, expected %q`, client.key, apiKey)
	}
}

func TestYelpClientNewRequest(t *testing.T) {
	client := YelpClient{}
	path := BaseYelpUrl + "/businesses/search"
	req := client.NewRequest(path)

	if req.path != path {
		t.Fatalf(`YelpRequest.path = %q, expected %q`, req.path, path)
	}
}

func TestYelpRequestAddParameter(t *testing.T) {
	client := YelpClient{}
	req := client.NewRequest("")
	expected := "test"

	req.AddParameter(expected, expected)

	actual, found := req.params[expected]
	if !found {
		t.Fatalf(`YelpRequest.param key %q, doesn't exist`, expected)
	}
	if actual != expected {
		t.Fatalf(`YelpRequest.param[%q] = %q, expected %q`, expected, actual, expected)
	}
}

func TestYelpRequestGet(t *testing.T) {
	client := YelpClient{
		key: "h4lGIGYbwWvSf-TJVWUU2sp-WRTtUQb8n8N-UmCOSn9vF4Aa8LjtycdFqkwtcSArTcQgLlJLEf-T7KfSJKakKiRE5kmNldcjQ7sTK4bwefaewZfRorg-0n3v02ZEX3Yx",
	}
	req := client.NewRequest(BaseYelpUrl + "/businesses/search")

	req.AddParameter("location", "Vancouver, Canada")

	res, err := req.Get()
	if err != nil {
		t.Fatalf(`YelpRequest.Get error: %q`, err)
	}
	if res.StatusCode == 0 {
		t.Fatalf(`YelpRequest.Get Invalid status: %d`, res.StatusCode)
	}
}

func TestYelpRequestGetInvalidKey(t *testing.T) {
	client := YelpClient{
		key: "Invalid",
	}
	req := client.NewRequest(BaseYelpUrl + "/businesses/search")

	req.AddParameter("location", "Vancouver, Canada")

	res, err := req.Get()
	if err != nil {
		t.Fatalf(`YelpRequest.Get error: %q`, err)
	}
	if res.StatusCode == 0 {
		t.Fatalf(`YelpRequest.Get Invalid status: %d`, res.StatusCode)
	}
	if res.StatusCode != 400 {
		t.Fatalf(`YelpRequest.Get Unexpected status: %d`, res.StatusCode)
	}

}

func TestYelpRequestInvalidParams(t *testing.T) {
	client := YelpClient{
		key: "h4lGIGYbwWvSf-TJVWUU2sp-WRTtUQb8n8N-UmCOSn9vF4Aa8LjtycdFqkwtcSArTcQgLlJLEf-T7KfSJKakKiRE5kmNldcjQ7sTK4bwefaewZfRorg-0n3v02ZEX3Yx",
	}
	req := client.NewRequest(BaseYelpUrl + "/businesses/search")

	// Not adding any parameters. Yelp Fusion requests at least the
	// 'location' parameter to be give nfor the business search.

	res, err := req.Get()

	if err != nil {
		t.Fatalf(`YelpRequest.Get error: %q`, err)
	}
	if res.StatusCode == 0 {
		t.Fatalf(`YelpRequest.Get Invalid status: %d`, res.StatusCode)
	}
	if res.StatusCode != 400 {
		t.Fatalf(`YelpRequest.Get Unexpected status: %d`, res.StatusCode)
	}
}
