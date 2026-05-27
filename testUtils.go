package main

import (
	"net/http"
	"net/http/httptest"
)

var simpleTestdataPaths = []string{
	"testdata/simple/index.html",
	"testdata/simple/simple.html",
	"testdata/simple/href.html",
	"testdata/simple/style.html",
} // full path (including "testdata") is needed since TestDownload uses the path to find the file to compare against

func getSimpleTestServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.Handle("/testdata/simple/", http.StripPrefix("/testdata/simple/", http.FileServer(http.Dir("testdata/simple"))))
	return httptest.NewServer(mux)
}

func getTop10TestServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.Handle("/top10/", http.StripPrefix("/top10/", http.FileServer(http.Dir("testdata/top10"))))
	mux.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/robots.txt")
	})
	return httptest.NewServer(mux)
}
