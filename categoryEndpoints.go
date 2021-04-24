package yelp

import (
	"fmt"
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
	return &AllCategoriesRequest{*yc.NewRequest("/categories", "AllCategories")}
}

func (acr *AllCategoriesRequest) Get() (*AllCategories, error) {
	data, err := acr.YelpRequest.Get()
	if err != nil {
		return nil, err
	}
	return &data.AllCategories, nil
}

func (acr *AllCategoriesRequest) GetResponse() (*http.Response, error) {
	return acr.YelpRequest.GetResponse()
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
		*yc.NewRequest(fmt.Sprintf("/categories/%s", categoryAlias), "CategoryDetails"),
	}
}

func (acr *CategoryDetailsRequest) Get() (*Category, error) {
	data, err := acr.YelpRequest.Get()
	if err != nil {
		return nil, err
	}
	return &data.CategoryWrapper.Category, nil
}

func (acr *CategoryDetailsRequest) GetResponse() (*http.Response, error) {
	return acr.YelpRequest.GetResponse()
}
