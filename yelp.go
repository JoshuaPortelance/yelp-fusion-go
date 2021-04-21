package yelp

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

const BaseYelpUrl = "https://api.yelp.com/v3"

type YelpClient struct {
	key     string
	timeout int
}

func (c *YelpClient) NewRequest(endpointPath string) *YelpRequest {
	return &YelpRequest{
		YelpClient{c.key, c.timeout},
		BaseYelpUrl + endpointPath,
		make(map[string]string),
	}
}

type YelpRequest struct {
	YelpClient
	path   string
	params map[string]string
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
		return res, err
	}
	if res.StatusCode != 200 {
		return res, errors.New(fmt.Sprint(res.StatusCode))
	}
	return res, nil
}
