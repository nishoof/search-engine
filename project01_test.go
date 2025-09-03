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
	{"<html>", "<body>", "Hello CS 272, there are no links here.", "</body>", "</html>"},
	{"<html>", "<body>", "For a simple example, see <a href=\"/test-data/project01/simple.html\">simple.html</a>", "</body>", "</html>"},
	{"<html>", "<head>", "  <title>Style</title>", "  <style>", "    a.blue {", "      color: blue;", "    }", "    a.red {", "      color: red;", "    }", "  </style>", "<body>", "  <p>", "    Here is a blue link to <a class=\"blue\" href=\"/test-data/project01/href.html\">href.html</a>", "  </p>", "  <p>", "    And a red link to <a class=\"red\" href=\"/test-data/project01/simple.html\">simple.html</a>", "  </p>", "</body>", "</html>"},
}

func TestDownload(t *testing.T) {
	tests := []struct {
		want []string
	}{
		{testData[0]},
		{testData[1]},
		{testData[2]},
	}

	for testIdx, test := range tests {
		// Create a test HTTP server that serves the expected HTML
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for _, v := range test.want {
				fmt.Fprintln(w, v)
			}
		}))

		got := download(testServer.URL)

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
		wantWords, wantHrefs []string
	}{
		{[]string{"Hello", "CS", "272", "there", "are", "no", "links", "here"}, []string{}},
		{[]string{"For", "a", "simple", "example", "see", "html"}, []string{"/test-data/project01/simple.html"}},
		{[]string{"Here", "is", "a", "blue", "link", "to", "href", "html", "And", "red", "simple"}, []string{"/test-data/project01/href.html", "/test-data/project01/simple.html"}},
	}

	for testIdx, test := range tests {
		testFileStr := strings.Join(testData[testIdx], "\n")
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
		{"https://usf-cs272-f25.github.io/test-data/project01/", []string{
			"https://usf-cs272-f25.github.io/test-data/project01/",
			"https://usf-cs272-f25.github.io/test-data/project01/simple.html",
			"https://usf-cs272-f25.github.io/test-data/project01/href.html",
			"https://usf-cs272-f25.github.io/test-data/project01/style.html",
		}},
	}

	for _, test := range tests {
		got, _ := crawl(test.seed)
		if len(got) != len(test.want) {
			t.Errorf("For seed %q, got %d URLs but wanted %d\n", test.seed, len(got), len(test.want))
		}
		for _, v := range test.want {
			if got[v] != struct{}{} {
				t.Errorf("For seed %q, got %v but wanted %v\n", test.seed, got, test.want)
				break
			}
		}
	}
}
