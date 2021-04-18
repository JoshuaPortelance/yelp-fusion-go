package yelp

import (
	"fmt"
	"net/http"
	"net/url"
)

const BaseYelpUrl = "https://api.yelp.com/v3"

type YelpClient struct {
	key     string
	timeout int
}

func (c *YelpClient) NewRequest(path string) *YelpRequest {
	return &YelpRequest{c.key, c.timeout, path, make(map[string]string)}
}

type YelpRequest struct {
	key     string
	timeout int
	path    string
	params  map[string]string
}

func (r *YelpRequest) AddParameter(name string, value string) {
	r.params[name] = value
}

func (r *YelpRequest) Get() (*http.Response, error) {
	// Creating a standalone request as we need to modify
	// it a bit before sending it.
	req, err := http.NewRequest("GET", r.path, nil)
	if err != nil {
		return &http.Response{}, err
	}

	// Adding API key to the headers.
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.key))

	// Adding all of the parameters to the URL.
	params := url.Values{}
	for key, value := range r.params {
		params.Add(key, value)
	}
	req.URL.RawQuery = params.Encode()

	// Sending the request.
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return &http.Response{}, err
	}
	defer res.Body.Close()
	return res, nil
}
