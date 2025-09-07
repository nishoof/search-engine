package main

import (
	"fmt"
	"io"
	"net/http"
)

/* Downloads the contents of the website of the given url and returns a ReadCloser on it */
func download(url string) io.ReadCloser {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error downloading %q: %v\n", url, err)
		return nil
	}
	return resp.Body
}
