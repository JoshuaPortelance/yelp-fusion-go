package yelp

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	want := "Hello, World!"
	msg := HelloWorld()
	if want != msg {
		t.Fatalf(`HelloWorld() = %q, want match for %#q, nil`, msg, want)
	}
}

func TestYelpClientCreation(t *testing.T) {
	apiKey := "dawefaf324125hQ#$%afsdfjl2jr2k342"
	timeout := 5000

	client := YelpClient{
		key:     apiKey,
		timeout: timeout,
	}
	if client.key != apiKey {
		t.Fatalf(`YelpClient.key = %q, expected %q`, client.key, apiKey)
	}
	if client.timeout != timeout {
		t.Fatalf(`YelpClient.timeout = %q, expected %q`, client.timeout, timeout)
	}
}

func TestYelpClientDefaultCreation(t *testing.T) {
	apiKey := ""
	timeout := 0

	client := YelpClient{}
	if client.key != apiKey {
		t.Fatalf(`YelpClient.key = %q, expected %q`, client.key, apiKey)
	}
	if client.timeout != timeout {
		t.Fatalf(`YelpClient.timeout = %q, expected %q`, client.timeout, timeout)
	}
}

func TestSearch(t *testing.T) {
	apiKey := "h4lGIGYbwWvSf-TJVWUU2sp-WRTtUQb8n8N-UmCOSn9vF4Aa8LjtycdFqkwtcSArTcQgLlJLEf-T7KfSJKakKiRE5kmNldcjQ7sTK4bwefaewZfRorg-0n3v02ZEX3Yx"
	timeout := 5000

	client := YelpClient{
		key:     apiKey,
		timeout: timeout,
	}

	client.Search()
}
