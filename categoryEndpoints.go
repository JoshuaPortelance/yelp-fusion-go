package yelp

import "fmt"

// YelpRequest for the All Categories endpoint.
//
// https://www.yelp.com/developers/documentation/v3/all_categories
func (c *YelpClient) NewAllCategories() *YelpRequest {
	return c.NewRequest("/categories")
}

// YelpRequest for the Category Details endpoint.
//
// https://www.yelp.com/developers/documentation/v3/category
func (c *YelpClient) NewCategoryDetails(categoryAlias string) *YelpRequest {
	return c.NewRequest(fmt.Sprintf("/categories/%s", categoryAlias))
}
