# yelp-fusion-go
[![Build Status](https://travis-ci.com/JoshuaPortelance/go-yelp.svg?token=qaKmxckZFKPT1MXTqhmw&branch=main)](https://travis-ci.com/JoshuaPortelance/go-yelp)
[![codecov](https://codecov.io/gh/JoshuaPortelance/yelp-fusion-go/branch/main/graph/badge.svg?token=F8G736FYQ3)](https://codecov.io/gh/JoshuaPortelance/yelp-fusion-go)

A Yelp Fusion API Golang package with no external dependencies.

Please refer to the official [Yelp Fusion documentation](https://www.yelp.com/developers/documentation/v3) for the request/response models.

# Table of Contents
[Usage Information:](#Usage-Information)

  * [Intallation](#Intallation)

  * Endpoint Request Interface:
    * [Get](#get)
    * [GetResponse](#get)

  * [Request Timeout](#Request-Timeout)

[Examples:](#Examples)

  * Business Endpoints:
    * [Business Search](#business-search)
    * [Phone](#phone-search)
    * [Delivery](#transaction-search)
    * [Business Details](#business-details)
    * [Match](#business-match)
    * [Reviews](#reviews)
    * [Autocomplete](#autocomplete)

  * Event Endpoints:
    * [Event Lookup](#event-lookup)
    * [Event Search](#event-search)
    * [Featured Event](#featured-event)

  * Category Endpoints:
    * [All Categories](#all-categories)
    * [Category Details](#category-details)

# Usage Information

## Intallation

```bash
go get github.com/JoshuaPortelance/yelp-fusion-go
```

## Endpoint Request Interface

### Get
Returns the unmarshalled JSON data following response model from the [Yelp Fusion documentation](https://www.yelp.com/developers/documentation/v3).

This is implemented per endpoint. The BusinessSearchRequest's definition of this is shown below. There is a similar definition for each endpoint request type.
```golang
func (r *BusinessSearchRequest) Get() (*Businesses, error)
```

### GetResponse
Returns an open and unmodified [http.Response](https://golang.org/pkg/net/http/#Response) pointer.
```golang
func (r *Request) GetResponse() (*http.Response, error)
```
The open [http.Response](https://golang.org/pkg/net/http/#Response) will need to be closed manually.

## Request Timeout
Timeout will set the [http.Client](https://golang.org/pkg/net/http/#Client) Timeout duration to abort the request if the server doesn't complete the response within that time in milliseconds.
```golang
client := yelp.Client{
    Key:     "YOUR_API_KEY",
    Timeout: 5000,
}
```

# Examples

## Business Endpoints 

### Business Search
```golang
client := yelp.Client{Key:"YOUR_API_KEY"}
search := client.NewBusinessSearch()
search.AddParameter("location", "Victoria, British Columbia")
search.AddParameter("term", "Red Fish Blue Fish")
data, _ := search.Get()
fmt.Println(data.Businesses[0].Name)
```

### Phone Search
```golang
client := yelp.Client{Key:"YOUR_API_KEY"}
search := client.NewPhoneSearch()
search.AddParameter("phone", "+12502986877")
data, _ := search.Get()
fmt.Println(data.Businesses[0].Name)
```

### Transaction Search
```golang
client := yelp.Client{Key:"YOUR_API_KEY"}
search := client.NewTransactionSearch("delivery")
search.AddParameter("location", "Seattle, Washington")
data, _ := search.Get()
fmt.Println(data.Businesses[0].Name)
```

### Business Details
```golang
client := yelp.Client{Key:"YOUR_API_KEY"}
search := client.NewBusinessDetails("red-fish-blue-fish-victoria")
search.AddParameter("location", "Victoria, British Columbia")
data, _ := search.Get()
fmt.Println(data.Name)
```

### Business Match
```golang
client := yelp.Client{Key:"YOUR_API_KEY"}
search := client.NewBusinessMatch()
search.AddParameter("name", "Red Fish Blue Fish")
search.AddParameter("address1", "1006 Wharf Street")
search.AddParameter("city", "Victoria")
search.AddParameter("state", "BC")
search.AddParameter("country", "CA")
data, _ := search.Get()
fmt.Println(data.Businesses[0].Name)
```

### Reviews
```golang
client := yelp.Client{Key:"YOUR_API_KEY"}
search := client.NewReviews("red-fish-blue-fish-victoria")
search.AddParameter("location", "Victoria, British Columbia")
data, _ := search.Get()
fmt.Println(data.Reviews[0].Text)
```

### Autocomplete
```golang
client := yelp.Client{Key:"YOUR_API_KEY"}
search := client.NewAutocomplete()
search.AddParameter("text", "Red Fish")
search.AddParameter("latitude", "48.4243002")
search.AddParameter("longitude", "-123.3706659")
data, _ := search.Get()
fmt.Println(data.Terms[0].Text)
```

## Event Endpoints

### Event Lookup
```golang
client := yelp.Client{Key:"YOUR_API_KEY"}
search := client.NewEventLookup("red-fish-blue-fish-victoria")
data, _ := search.Get()
fmt.Println(data.Description)
```

### Event Search
```golang
client := yelp.Client{Key:"YOUR_API_KEY"}
search := client.NewEventSearch()
search.AddParameter("location", "Victoria, British Columbia")
data, _ := search.Get()
fmt.Println(data.Events[0].Name)
```

### Featured Event
```golang
client := yelp.Client{Key:"YOUR_API_KEY"}
search := client.NewFeaturedEvent()
search.AddParameter("location", "Victoria, British Columbia")
data, _ := search.Get()
fmt.Println(data.Description)
```

## Category Endpoints 

### All Categories
```golang
client := yelp.Client{Key:"YOUR_API_KEY"}
search := client.NewAllCategories()
data, _ := search.Get()
fmt.Println(data.Categories[0].Alias)
```

### Category Details
```golang
client := yelp.Client{Key:"YOUR_API_KEY"}
search := client.NewCategoryDetails("arts")
data, _ := search.Get()
fmt.Println(data.Title)
```
