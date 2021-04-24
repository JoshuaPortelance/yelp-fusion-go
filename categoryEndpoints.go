package yelp

import (
	"fmt"
	"net/http"
)

/*
Request for the All Categories endpoint.
https://www.yelp.com/developers/documentation/v3/all_categories
*/
type AllCategoriesRequest struct {
	Request
}

func (yc *Client) NewAllCategories() *AllCategoriesRequest {
	return &AllCategoriesRequest{*yc.NewRequest("/categories", "AllCategories")}
}

func (acr *AllCategoriesRequest) Get() (*Categories, error) {
	data, err := acr.Request.Get()
	if err != nil {
		return nil, err
	}
	return &data.Categories, nil
}

func (acr *AllCategoriesRequest) GetResponse() (*http.Response, error) {
	return acr.Request.GetResponse()
}

/*
Request for the Category Details endpoint.
https://www.yelp.com/developers/documentation/v3/category
*/
type CategoryDetailsRequest struct {
	Request
}

func (yc *Client) NewCategoryDetails(categoryAlias string) *CategoryDetailsRequest {
	return &CategoryDetailsRequest{
		*yc.NewRequest(fmt.Sprintf("/categories/%s", categoryAlias), "CategoryDetails"),
	}
}

func (acr *CategoryDetailsRequest) Get() (*Category, error) {
	data, err := acr.Request.Get()
	if err != nil {
		return nil, err
	}
	return &data.CategoryWrapper.Category, nil
}

func (acr *CategoryDetailsRequest) GetResponse() (*http.Response, error) {
	return acr.Request.GetResponse()
}
