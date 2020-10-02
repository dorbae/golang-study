package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

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
		res := "OK"
		err := hitURL(url)
		if err != nil {
			res = "FAILED"

		}
		results[url] = res
	}

	for url, res := range results {
		fmt.Println(url, res)
	}

	// #3.6 URLChecker + Go Routines
	// https://nomadcoders.co/go-for-beginners/lectures/1525
	fmt.Println("Goroutine and channel")
	c := make(chan requestResult)
	for _, url := range urls {
		go hitURLWithConcurrency(url, c)
	}
	for i := 0; i < len(urls); i++ {
		r := <-c
		fmt.Println(r)
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

func hitURLWithConcurrency(url string, c chan<- requestResult) {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	fmt.Println("[DEBUG]", err, resp.StatusCode)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
}
