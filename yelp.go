package yelp

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const BaseYelpUrl = "https://api.yelp.com/v3"

type YelpClient struct {
	key     string
	timeout int
}

func (c *YelpClient) NewRequest(path string, endpoint string) *YelpRequest {
	return &YelpRequest{
		YelpClient: YelpClient{c.key, c.timeout},
		path:       BaseYelpUrl + path,
		endpoint:   endpoint,
		params:     make(map[string]string),
	}
}

type YelpRequest struct {
	YelpClient
	path     string
	endpoint string
	params   map[string]string
}

func (r *YelpRequest) AddParameter(name string, value string) {
	r.params[name] = value
}

func (r *YelpRequest) GetResponse() (*http.Response, error) {
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
	client := &http.Client{Timeout: time.Duration(r.YelpClient.timeout) * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return res, err
	}
	if res.StatusCode != 200 {
		return res, errors.New(fmt.Sprint(res.StatusCode))
	}
	return res, nil
}

func (r *YelpRequest) Get() (interface{}, error) {
	response, err := r.GetResponse()
	if err != nil {
		return nil, err
	}

	responsebody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Figuring out the type to use for unmarshalling.
	switch r.endpoint {
	case "BusinessSearch", "PhoneSearch", "TransactionSearch", "BusinessMatch":
		data := Businesses{}
		err = json.Unmarshal(responsebody, &data)
		if err != nil {
			return nil, err
		}
		return data, nil
	case "BusinessDetails":
		data := Business{}
		err = json.Unmarshal(responsebody, &data)
		if err != nil {
			return nil, err
		}
		return data, nil
	case "Reviews":
		data := Reviews{}
		err = json.Unmarshal(responsebody, &data)
		if err != nil {
			return nil, err
		}
		return data, nil
	case "Autocomplete":
		data := Autocomplete{}
		err = json.Unmarshal(responsebody, &data)
		if err != nil {
			return nil, err
		}
		return data, nil
	case "EventLookup", "FeaturedEvent":
		data := Event{}
		err = json.Unmarshal(responsebody, &data)
		if err != nil {
			return nil, err
		}
		return data, nil
	case "EventSearch":
		data := Events{}
		err = json.Unmarshal(responsebody, &data)
		if err != nil {
			return nil, err
		}
		return data, nil
	case "AllCategories":
		data := AllCategories{}
		err = json.Unmarshal(responsebody, &data)
		if err != nil {
			return nil, err
		}
		return data, nil
	case "CategoryDetails":
		data := CategoryWrapper{}
		err = json.Unmarshal(responsebody, &data)
		if err != nil {
			return nil, err
		}
		return data, nil
	default:
		return nil, errors.New("unknown endpoint key")
	}
}
