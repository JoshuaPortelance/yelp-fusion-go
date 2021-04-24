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

type Client struct {
	key     string
	timeout int
}

func (c *Client) NewRequest(path string, endpoint string) *Request {
	return &Request{
		Client:   Client{c.key, c.timeout},
		path:     BaseYelpUrl + path,
		endpoint: endpoint,
		params:   make(map[string]string),
	}
}

type Request struct {
	Client
	path     string
	endpoint string
	params   map[string]string
}

func (r *Request) AddParameter(name string, value string) {
	r.params[name] = value
}

func (r *Request) GetResponse() (*http.Response, error) {
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
	client := &http.Client{Timeout: time.Duration(r.Client.timeout) * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return res, err
	}
	if res.StatusCode != 200 {
		return res, errors.New(fmt.Sprint(res.StatusCode))
	}
	return res, nil
}

// Provides a way for us to have a semi-generic return type
// without the need for type assertions or repeated code
// due to scoped variable assignment of varying types.
type GetReturnType struct {
	Business
	Businesses
	Reviews
	Autocomplete
	Event
	Events
	Categories
	CategoryWrapper
}

func (r *Request) Get() (*GetReturnType, error) {
	var err error
	response, err := r.GetResponse()
	if err != nil {
		return nil, err
	}

	responsebody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data := GetReturnType{}

	// Figuring out the type to use for unmarshalling.
	switch r.endpoint {
	case "BusinessSearch", "PhoneSearch", "TransactionSearch", "BusinessMatch":
		data.Businesses = Businesses{}
		err = json.Unmarshal(responsebody, &data.Businesses)
	case "BusinessDetails":
		data.Business = Business{}
		err = json.Unmarshal(responsebody, &data.Business)
	case "Reviews":
		data.Reviews = Reviews{}
		err = json.Unmarshal(responsebody, &data.Reviews)
	case "Autocomplete":
		data.Autocomplete = Autocomplete{}
		err = json.Unmarshal(responsebody, &data.Autocomplete)
	case "EventLookup", "FeaturedEvent":
		data.Event = Event{}
		err = json.Unmarshal(responsebody, &data.Event)
	case "EventSearch":
		data.Events = Events{}
		err = json.Unmarshal(responsebody, &data.Events)
	case "AllCategories":
		data.Categories = Categories{}
		err = json.Unmarshal(responsebody, &data.Categories)
	case "CategoryDetails":
		data.CategoryWrapper = CategoryWrapper{}
		err = json.Unmarshal(responsebody, &data.CategoryWrapper)
	default:
		return nil, errors.New("unknown endpoint key")
	}

	if err != nil {
		return nil, err
	}
	return &data, nil
}
