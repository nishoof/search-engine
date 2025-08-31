package main

import (
	"fmt"
	"io"
	"net/http"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func download(url string) io.ReadCloser {
	resp, err := http.Get(url)
	checkErr(err)
	return resp.Body
}

func main() {
	fmt.Println("Hello, World!")

	// body := download("https://usf-cs272-f25.github.io/test-data/project01/simple.html")
}
