package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, url := range urls {
		go check(url, c)
	}

	// waits for the channel to receive a message
	for u := range c {
		go func(url string) {
			time.Sleep(5 * time.Second)
			check(url, c)
		}(u)
	}
}

func check(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "might be down")
		c <- url
		return
	}
	fmt.Println(url, "is up!")
	c <- url
}
