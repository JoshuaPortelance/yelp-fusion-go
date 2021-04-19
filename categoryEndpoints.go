package yelp

import "fmt"

// YelpRequest for the All Categories endpoint.
func (c *YelpClient) NewAllCategories() *YelpRequest {
	return c.NewRequest("/categories")
}

// YelpRequest for the Category Details endpoint.
func (c *YelpClient) NewCategoryDetails(categoryAlias string) *YelpRequest {
	return c.NewRequest(fmt.Sprintf("/categories/%s", categoryAlias))
}
