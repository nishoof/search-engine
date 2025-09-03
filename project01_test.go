package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var testData = [][]string{
	{"<html>", "<body>", "  <ul>", "    <li>", "      <a href=\"/test-data/project01/simple.html\">simple.html</a>", "    </li>", "    <li>", "      <a href=\"/test-data/project01/href.html\">href.html</a>", "    </li>", "    <li>", "      <a href=\"/test-data/project01/style.html\">style.html</a>", "  </ul>", "</body>", "</html>"},
	{"<html>", "<body>", "Hello CS 272, there are no links here.", "</body>", "</html>"},
	{"<html>", "<body>", "For a simple example, see <a href=\"/test-data/project01/simple.html\">simple.html</a>", "</body>", "</html>"},
	{"<html>", "<head>", "  <title>Style</title>", "  <style>", "    a.blue {", "      color: blue;", "    }", "    a.red {", "      color: red;", "    }", "  </style>", "<body>", "  <p>", "    Here is a blue link to <a class=\"blue\" href=\"/test-data/project01/href.html\">href.html</a>", "  </p>", "  <p>", "    And a red link to <a class=\"red\" href=\"/test-data/project01/simple.html\">simple.html</a>", "  </p>", "</body>", "</html>"},
}

var testPaths = []string{
	"/test-data/project01",
	"/test-data/project01/simple.html",
	"/test-data/project01/href.html",
	"/test-data/project01/style.html",
}

func getTestServer() *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("r.URL.Path:", r.URL.Path)
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

func TestDownload(t *testing.T) {
	tests := []struct {
		path string
		want []string
	}{
		{testPaths[0], testData[0]},
		{testPaths[1], testData[1]},
		{testPaths[2], testData[2]},
	}

	ts := getTestServer()
	defer ts.Close()

	for testIdx, test := range tests {
		got := download(ts.URL + test.path)

		// Read through the downloaded content line by line with a scanner and compare it to expected
		scanner := bufio.NewScanner(got)
		for i := 0; scanner.Scan(); i++ {
			textGot := scanner.Text()
			textWant := test.want[i]
			if textGot != textWant {
				t.Errorf("For test %d at line %d, got \"%s\" but wanted \"%s\"\n", testIdx, i, textGot, textWant)
			}
		}
	}
}

func TestExtract(t *testing.T) {
	tests := []struct {
		testData, wantWords, wantHrefs []string
	}{
		{testData[1], []string{"Hello", "CS", "272", "there", "are", "no", "links", "here"}, []string{}},
		{testData[2], []string{"For", "a", "simple", "example", "see", "html"}, []string{"/test-data/project01/simple.html"}},
		{testData[3], []string{"Here", "is", "a", "blue", "link", "to", "href", "html", "And", "red", "simple"}, []string{"/test-data/project01/href.html", "/test-data/project01/simple.html"}},
	}

	for testIdx, test := range tests {
		testFileStr := strings.Join(test.testData, "\n")
		stringsReader := strings.NewReader(testFileStr)
		bufioReader := bufio.NewReader(stringsReader)
		gotWords, gotHrefs := extract(bufioReader)

		// words
		if len(gotWords) != len(test.wantWords) {
			t.Errorf("For test %d, got %d words but wanted %d\n", testIdx, len(gotWords), len(test.wantWords))
		}
		for _, word := range test.wantWords {
			if gotWords[word] != struct{}{} {
				t.Errorf("For test %d, word %q is missing from gotWords\n", testIdx, word)
			}
		}

		// hrefs
		if len(gotHrefs) != len(test.wantHrefs) {
			t.Errorf("For test %d, got %d hrefs but wanted %d\n", testIdx, len(gotHrefs), len(test.wantHrefs))
		}
		for _, href := range test.wantHrefs {
			if gotHrefs[href] != struct{}{} {
				t.Errorf("For test %d, href %q is missing from gotHrefs\n", testIdx, href)
			}
		}
	}
}

func TestCleanHref(t *testing.T) {
	tests := []struct {
		host, href, want string
	}{
		{"https://usf-cs272-f25.github.io/", "/", "https://usf-cs272-f25.github.io/"},
		{"https://usf-cs272-f25.github.io/", "/help/", "https://usf-cs272-f25.github.io/help/"},
		{"https://usf-cs272-f25.github.io/", "/syllabus/", "https://usf-cs272-f25.github.io/syllabus/"},
		{"https://usf-cs272-f25.github.io/", "https://gobyexample.com/", "https://gobyexample.com/"},
	}

	for _, test := range tests {
		got := cleanHref(test.host, test.href)
		if got != test.want {
			t.Errorf("For base %q and href %q, got %q but wanted %q\n", test.host, test.href, got, test.want)
		}
	}
}

func TestCrawl(t *testing.T) {
	tests := []struct {
		seed string
		want []string
	}{
		{testPaths[0], []string{
			testPaths[0],
			testPaths[1],
			testPaths[2],
			testPaths[3],
		}},
	}

	ts := getTestServer()
	defer ts.Close()

	for _, test := range tests {
		got, _ := crawl(ts.URL + test.seed)
		if len(got) != len(test.want) {
			t.Errorf("For seed %q, got %d URLs but wanted %d\n", test.seed, len(got), len(test.want))
		}
		count := 0
		for _, v := range test.want {
			_, exists := got[ts.URL+v]
			if exists {
				count++
			}
		}
		if count != len(test.want) {
			t.Errorf("For seed %q, got %v but wanted %v\n", test.seed, got, test.want)
		}
	}
}
