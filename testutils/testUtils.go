package testutils

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
)

var SimpleTestdataPaths = []string{
	"testdata/simple/index.html",
	"testdata/simple/simple.html",
	"testdata/simple/href.html",
	"testdata/simple/style.html",
} // relative to root of the project

func NewSimpleTestServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.Handle("/testdata/simple/", http.StripPrefix("/testdata/simple/", http.FileServer(http.Dir(Root()+"/testdata/simple"))))
	return httptest.NewServer(mux)
}

func NewTop10TestServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.Handle("/top10/", http.StripPrefix("/top10/", http.FileServer(http.Dir(Root()+"/testdata/top10"))))
	mux.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, Root()+"/testdata/robots.txt")
	})
	return httptest.NewServer(mux)
}

func Root() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filepath.Dir(filename))
}
