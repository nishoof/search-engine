package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
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
		{[]string{"For", "a", "simple", "example", "see", "simple", "html"}, []string{"/test-data/project01/simple.html"}},
		{[]string{"Here", "is", "a", "blue", "link", "to", "href", "html", "And", "a", "red", "link", "to", "simple", "html"}, []string{"/test-data/project01/href.html", "/test-data/project01/simple.html"}},
	}

	for testIdx, test := range tests {
		testFileStr := strings.Join(testData[testIdx], "\n")
		stringsReader := strings.NewReader(testFileStr)
		bufioReader := bufio.NewReader(stringsReader)
		gotWords, gotHrefs := extract(bufioReader)
		if !reflect.DeepEqual(gotWords, test.wantWords) {
			t.Errorf("For test %d, got words [%s] but wanted [%s]\n", testIdx, strings.Join(gotWords, ", "), strings.Join(test.wantWords, ", "))
		}
		if !reflect.DeepEqual(gotHrefs, test.wantHrefs) {
			t.Errorf("For test %d, got hrefs [%s] but wanted [%s]\n", testIdx, strings.Join(gotHrefs, ", "), strings.Join(test.wantHrefs, ", "))
		}
	}
}
