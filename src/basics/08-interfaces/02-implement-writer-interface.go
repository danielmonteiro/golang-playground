package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println(resp)
	fmt.Println("1 ============================")

	// 1st way to print the response body
	bs := make([]byte, 999999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
	fmt.Println("2 ============================")

	// 2nd way to print the response body (will be empty as it was printed before)
	io.Copy(os.Stdout, resp.Body)
	fmt.Println("\n3 ============================")

	// Implementing our own Writer (will be empty as it was printed before)
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}
