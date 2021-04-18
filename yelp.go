package yelp

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func HelloWorld() string {
	return "Hello, World!"
}

type YelpClient struct {
	key     string
	timeout int
}

func (c *YelpClient) Search() {
	var url = "https://api.yelp.com/v3/businesses/search?location=Vancouver"
	//var jsonStr = []byte(`{"location":"Vancouver, Canada", "categories":"restaurants"}`)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.key))
	//req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	resbody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(resbody))
	//reqbody, _ := ioutil.ReadAll(req.Body)
	//fmt.Println("request Body:", string(reqbody))
}

/*
func test() {
	key := "h4lGIGYbwWvSf-TJVWUU2sp-WRTtUQb8n8N-UmCOSn9vF4Aa8LjtycdFqkwtcSArTcQgLlJLEf-T7KfSJKakKiRE5kmNldcjQ7sTK4bwefaewZfRorg-0n3v02ZEX3Yx"
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}
	req, err := http.NewRequest("GET", , nil)
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)
}*/
