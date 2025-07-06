package main

import (
	"fmt"
	"net/url"
)

func main() {
	//URL : Uniform resource locator
	//[Scheme://][userinfo@]host[:port][/path][?query][#fragment]
	rawUrl := "https://example.com:8080/path?query=param#fragment"

	parsedURL, err := url.Parse(rawUrl)
	if err != nil {
		fmt.Println("Error parsing URL : ", err)
		return
	}

	fmt.Println("Scheme : ", parsedURL.Scheme)
	fmt.Println("Host : ", parsedURL.Host)
	fmt.Println("Post : ", parsedURL.Port())
	fmt.Println("Path : ", parsedURL.Path)
	fmt.Println("RawQuery : ", parsedURL.RawQuery)
	fmt.Println("Fragment : ", parsedURL.Fragment)

	RawURL := "https://example.com/path?name=John&age=30"
	ParsedURL, err := url.Parse(RawURL)
	if err != nil {
		fmt.Println(err)
	}

	queryparams := ParsedURL.Query()
	fmt.Println("queryParams : ", queryparams)
	fmt.Println("Name : ", queryparams.Get("name"))
	fmt.Println("age : ", queryparams.Get("age"))

	// Building a URL :
	BaseURL := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "/path",
	}

	query := BaseURL.Query()
	query.Set("name", "John")
	query.Set("age", "30")
	BaseURL.RawQuery = query.Encode()

	fmt.Println("Built Url : ", BaseURL.String())

	// adding key value pair values:
	values := url.Values{}
	values.Add("name", "Jane")
	values.Add("age", "30")
	values.Add("city", "london")
	values.Add("country", "UK")

	//Encode:
	encodedqry := values.Encode()
	fmt.Println("Encoded Url", encodedqry)

	base := "https://example.com/search"
	fullURL := base + "?" + encodedqry
	fmt.Println(fullURL)

}
