# yelp-go
[![Build Status](https://travis-ci.com/JoshuaPortelance/go-yelp.svg?token=qaKmxckZFKPT1MXTqhmw&branch=main)](https://travis-ci.com/JoshuaPortelance/go-yelp)
[![codecov](https://codecov.io/gh/JoshuaPortelance/go-yelp/branch/main/graph/badge.svg?token=F8G736FYQ3)](https://codecov.io/gh/JoshuaPortelance/go-yelp)

Yelp Fusion API golang package with no external dependencies.

## Business Endpoints 

### Business Search
```golang
package main

import (
	"github.com/JoshuaPortelance/go-yelp"
)

client := yelp.client{key:"YOUR_API_KEY"};
search := client.NewBusinessSearch()
search.AddParameter("location", "Victoria, British Columbia")
search.AddParameter("term", "Red Fish Blue Fish")
res, _ := search.Get() // *http.Response
resbody, _ := ioutil.ReadAll(res.Body)
defer res.Body.Close()
businesses := Businesses{}
json.Unmarshal(resbody, &businesses)

// Goal Option 1
client := yelp.client{key:"YOUR_API_KEY"};
search := client.NewBusinessSearch()
search.AddParameter("location", "Victoria, British Columbia")
search.AddParameter("term", "Red Fish Blue Fish")
res, unmarshalBody, err := search.Get() // res is already closed.

// Goal Option 2
client := yelp.client{key:"YOUR_API_KEY"};
search := client.NewBusinessSearch()
search.AddParameter("location", "Victoria, British Columbia")
search.AddParameter("term", "Red Fish Blue Fish")
res, err := search.GetOpenResponse()
//... do something with response ...
defer res.Body.Close() // Don't forget to close it.

client.search({
  term: 'Four Barrel Coffee',
  location: 'san francisco, ca',
}).then(response => {
  console.log(response.jsonBody.businesses[0].name);
}).catch(e => {
  console.log(e);
});
```