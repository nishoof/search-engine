package main

import (
	"net/http"
	"net/http/httptest"
)

var testData = [][]string{
	{"<html>", "<body>", "<ul>", "<li>", "<a href=\"/testdata/simple/simple.html\">simple.html</a>", "</li>", "<li>", "<a href=\"/testdata/simple/href.html\">href.html</a>", "</li>", "<li>", "<a href=\"/testdata/simple/style.html\">style.html</a>", "</li>", "</ul>", "</body>", "</html>"},
	{"<html>", "<body>", "Hello CS 272, there are no links here.", "</body>", "</html>"},
	{"<html>", "<body>", "For a simple example, see <a href=\"/testdata/simple/simple.html\">simple.html</a>", "</body>", "</html>"},
	{"<html>", "<head>", "<title>Style</title>", "<style>", "a.blue {", "color: blue;", "}", "a.red {", "color: red;", "}", "</style>", "</head>", "<body>", "<p>", "Here is a blue link to <a class=\"blue\" href=\"/testdata/simple/href.html\">href.html</a>", "</p>", "<p>", "And a red link to <a class=\"red\" href=\"/testdata/simple/simple.html\">simple.html</a>", "</p>", "</body>", "</html>"},
}

var testPaths = []string{
	"/testdata/simple",
	"/testdata/simple/simple.html",
	"/testdata/simple/href.html",
	"/testdata/simple/style.html",
}

func getTestServer() *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case testPaths[0]:
			http.ServeFile(w, r, "testdata/simple/index.html")
		case testPaths[1]:
			http.ServeFile(w, r, "testdata/simple/simple.html")
		case testPaths[2]:
			http.ServeFile(w, r, "testdata/simple/href.html")
		case testPaths[3]:
			http.ServeFile(w, r, "testdata/simple/style.html")
		default:
			http.NotFound(w, r)
		}
	}))
	return ts
}
