package yelp

import "fmt"

type Category struct {
	Alias            string   `json:"alias"`
	Title            string   `json:"title"`
	ParentAliases    []string `json:"parent_aliases"`
	CountryWhitelist []string `json:"country_whitelist"`
	CountryBlacklist []string `json:"country_blacklist"`
}

type AllCategories struct {
	Categories []Category `json:"categories"`
}

// YelpRequest for the All Categories endpoint.
func (c *YelpClient) NewAllCategories() *YelpRequest {
	return c.NewRequest("/categories")
}

// YelpRequest for the Category Details endpoint.
func (c *YelpClient) NewCategoryDetails(categoryAlias string) *YelpRequest {
	return c.NewRequest(fmt.Sprintf("/categories/%s", categoryAlias))
}
