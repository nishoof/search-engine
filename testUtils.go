package main

import (
	"net/http"
	"net/http/httptest"
)

var testData = [][]string{
	{"<html>", "<body>", "  <ul>", "    <li>", "      <a href=\"/test-data/search-engine/simple.html\">simple.html</a>", "    </li>", "    <li>", "      <a href=\"/test-data/search-engine/href.html\">href.html</a>", "    </li>", "    <li>", "      <a href=\"/test-data/search-engine/style.html\">style.html</a>", "  </ul>", "</body>", "</html>"},
	{"<html>", "<body>", "Hello CS 272, there are no links here.", "</body>", "</html>"},
	{"<html>", "<body>", "For a simple example, see <a href=\"/test-data/search-engine/simple.html\">simple.html</a>", "</body>", "</html>"},
	{"<html>", "<head>", "  <title>Style</title>", "  <style>", "    a.blue {", "      color: blue;", "    }", "    a.red {", "      color: red;", "    }", "  </style>", "<body>", "  <p>", "    Here is a blue link to <a class=\"blue\" href=\"/test-data/search-engine/href.html\">href.html</a>", "  </p>", "  <p>", "    And a red link to <a class=\"red\" href=\"/test-data/search-engine/simple.html\">simple.html</a>", "  </p>", "</body>", "</html>"},
}

var testPaths = []string{
	"/test-data/search-engine",
	"/test-data/search-engine/simple.html",
	"/test-data/search-engine/href.html",
	"/test-data/search-engine/style.html",
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
		case "/test-data/rnj/sceneI_30.0.html":
			http.ServeFile(w, r, "test-data/rnj/sceneI_30.0.html")
		case "/test-data/rnj/sceneI_30.1.html":
			http.ServeFile(w, r, "test-data/rnj/sceneI_30.1.html")
		case "/test-data/rnj/sceneI_30.2.html":
			http.ServeFile(w, r, "test-data/rnj/sceneI_30.2.html")
		case "/test-data/rnj/sceneI_30.3.html":
			http.ServeFile(w, r, "test-data/rnj/sceneI_30.3.html")
		case "/test-data/rnj/sceneI_30.4.html":
			http.ServeFile(w, r, "test-data/rnj/sceneI_30.4.html")
		case "/test-data/rnj/sceneI_30.5.html":
			http.ServeFile(w, r, "test-data/rnj/sceneI_30.5.html")
		case "/test-data/rnj/sceneII_30.0.html":
			http.ServeFile(w, r, "test-data/rnj/sceneII_30.0.html")
		case "/test-data/rnj/sceneII_30.1.html":
			http.ServeFile(w, r, "test-data/rnj/sceneII_30.1.html")
		case "/test-data/rnj/sceneII_30.2.html":
			http.ServeFile(w, r, "test-data/rnj/sceneII_30.2.html")
		case "/test-data/rnj/sceneII_30.3.html":
			http.ServeFile(w, r, "test-data/rnj/sceneII_30.3.html")
		case "/test-data/rnj/":
			http.ServeFile(w, r, "test-data/rnj/index.html")
		default:
			http.NotFound(w, r)
		}
	}))
	return ts
}
