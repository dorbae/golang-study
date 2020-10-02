package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	// #3.0 hitURL
	// https://nomadcoders.co/go-for-beginners/lectures/1519
	urls := []string{
		"https://www.naver.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.soundcloud.com/",
		"https://www.instagram.com/",
	}

	// var result = map[string]string{}
	var results = make(map[string]string)

	for _, url := range urls {
		// hitURL(url)
		// #3.1 Slow URLChecker
		// https://nomadcoders.co/go-for-beginners/lectures/1520
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"

		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	// https://golang.org/pkg/
	resp, err := http.Get(url)
	fmt.Println("[DEBUG]", err, resp.StatusCode)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
