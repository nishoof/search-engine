package main

import (
	"net/http"
	"net/http/httptest"
)

const testdataDir = "testdata/simple"

var testdataPaths = []string{
	testdataDir + "/index.html",
	testdataDir + "/simple.html",
	testdataDir + "/href.html",
	testdataDir + "/style.html",
}

func getTestServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.Handle("/testdata/simple/", http.StripPrefix("/testdata/simple/", http.FileServer(http.Dir(testdataDir))))
	return httptest.NewServer(mux)
}
