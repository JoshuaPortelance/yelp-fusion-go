package yelp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
YelpRequest for the All Categories endpoint.
https://www.yelp.com/developers/documentation/v3/all_categories
*/
type AllCategoriesRequest struct {
	YelpRequest
}

func (yc *YelpClient) NewAllCategories() *AllCategoriesRequest {
	return &AllCategoriesRequest{*yc.NewRequest("/categories")}
}

func (acr *AllCategoriesRequest) Get() (*AllCategories, error) {
	response, err := acr.GetResponse()
	if err != nil {
		return &AllCategories{}, err
	}

	responsebody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &AllCategories{}, err
	}
	defer response.Body.Close()

	allCategories := AllCategories{}
	err = json.Unmarshal(responsebody, &allCategories)
	if err != nil {
		return &AllCategories{}, err
	}
	return &allCategories, nil
}

func (acr *AllCategoriesRequest) GetResponse() (*http.Response, error) {
	return acr.YelpRequest.Get()
}

/*
YelpRequest for the Category Details endpoint.
https://www.yelp.com/developers/documentation/v3/category
*/
type CategoryDetailsRequest struct {
	YelpRequest
}

func (yc *YelpClient) NewCategoryDetails(categoryAlias string) *CategoryDetailsRequest {
	return &CategoryDetailsRequest{
		*yc.NewRequest(fmt.Sprintf("/categories/%s", categoryAlias)),
	}
}

func (acr *CategoryDetailsRequest) Get() (*Category, error) {
	response, err := acr.GetResponse()
	if err != nil {
		return &Category{}, err
	}

	responsebody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &Category{}, err
	}
	defer response.Body.Close()

	category := Category{}
	err = json.Unmarshal(responsebody, &category)
	if err != nil {
		return &Category{}, err
	}
	return &category, nil
}

func (acr *CategoryDetailsRequest) GetResponse() (*http.Response, error) {
	return acr.YelpRequest.Get()
}
