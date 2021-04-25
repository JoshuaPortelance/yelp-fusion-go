package yelp

import (
	"fmt"
)

/*
Request type and override(s) for the All Categories endpoint.
https://www.yelp.com/developers/documentation/v3/all_categories
*/
type AllCategoriesRequest struct {
	Request
}

func (c *Client) NewAllCategories() *AllCategoriesRequest {
	return &AllCategoriesRequest{*c.NewRequest("/categories", "AllCategories")}
}

func (r *AllCategoriesRequest) Get() (*Categories, error) {
	data, err := r.Request.Get()
	if err != nil {
		return nil, err
	}
	return &data.Categories, nil
}

/*
Request type and override(s) for the Category Details endpoint.
https://www.yelp.com/developers/documentation/v3/category
*/
type CategoryDetailsRequest struct {
	Request
}

func (c *Client) NewCategoryDetails(categoryAlias string) *CategoryDetailsRequest {
	return &CategoryDetailsRequest{
		*c.NewRequest(fmt.Sprintf("/categories/%s", categoryAlias), "CategoryDetails"),
	}
}

func (r *CategoryDetailsRequest) Get() (*Category, error) {
	data, err := r.Request.Get()
	if err != nil {
		return nil, err
	}
	return &data.CategoryWrapper.Category, nil
}
