package main

import (
	"fmt"
	"net/url"
)

func main() {

	dummyURL := "https://www.example.com:8080/path/to/page?user=john&age=30#section1"

	parsedURL, err := url.Parse(dummyURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	fmt.Println("Full URL:", parsedURL.String())
	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Hostname:", parsedURL.Hostname())
	fmt.Println("Port:", parsedURL.Port())
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("Raw Query:", parsedURL.RawQuery)
	fmt.Println("Fragment:", parsedURL.Fragment)

	queryParams := parsedURL.Query()
	fmt.Println("Query Parameters:")
	for key, values := range queryParams {
		for _, value := range values {
			fmt.Printf("  %s: %s\n", key, value)
		}
	}
}
