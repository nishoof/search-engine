package main

import (
	"net/http"
	"net/http/httptest"
)

var testData = [][]string{
	{"<html>", "<body>", "  <ul>", "    <li>", "      <a href=\"/test-data/lab03/simple.html\">simple.html</a>", "    </li>", "    <li>", "      <a href=\"/test-data/lab03/href.html\">href.html</a>", "    </li>", "    <li>", "      <a href=\"/test-data/lab03/style.html\">style.html</a>", "  </ul>", "</body>", "</html>"},
	{"<html>", "<body>", "Hello CS 272, there are no links here.", "</body>", "</html>"},
	{"<html>", "<body>", "For a simple example, see <a href=\"/test-data/lab03/simple.html\">simple.html</a>", "</body>", "</html>"},
	{"<html>", "<head>", "  <title>Style</title>", "  <style>", "    a.blue {", "      color: blue;", "    }", "    a.red {", "      color: red;", "    }", "  </style>", "<body>", "  <p>", "    Here is a blue link to <a class=\"blue\" href=\"/test-data/lab03/href.html\">href.html</a>", "  </p>", "  <p>", "    And a red link to <a class=\"red\" href=\"/test-data/lab03/simple.html\">simple.html</a>", "  </p>", "</body>", "</html>"},
}

var testPaths = []string{
	"/test-data/lab03",
	"/test-data/lab03/simple.html",
	"/test-data/lab03/href.html",
	"/test-data/lab03/style.html",
}

func getTestServer() *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case testPaths[0]:
			http.ServeFile(w, r, "test-data/index.html")
		case testPaths[1]:
			http.ServeFile(w, r, "test-data/simple.html")
		case testPaths[2]:
			http.ServeFile(w, r, "test-data/href.html")
		case testPaths[3]:
			http.ServeFile(w, r, "test-data/style.html")
		default:
			http.NotFound(w, r)
		}
	}))
	return ts
}
