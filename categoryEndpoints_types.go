package yelp

type Category struct {
	Alias            string   `json:"alias"`
	Title            string   `json:"title"`
	ParentAliases    []string `json:"parent_aliases"`
	CountryWhitelist []string `json:"country_whitelist"`
	CountryBlacklist []string `json:"country_blacklist"`
}

type CategoryWrapper struct {
	Category Category `json:"category"`
}

type AllCategories struct {
	Categories []Category `json:"categories"`
}
