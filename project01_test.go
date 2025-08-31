package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDownload(t *testing.T) {
	tests := []struct {
		want []string
	}{
		{[]string{"<html>", "<body>", "Hello CS 272, there are no links here.", "</body>", "</html>"}},
		{[]string{"<html>", "<body>", "For a simple example, see <a href=\"/test-data/project01/simple.html\">simple.html</a>", "</body>", "</html>"}},
		{[]string{"<html>", "<head>", "  <title>Style</title>", "  <style>", "    a.blue {", "      color: blue;", "    }", "    a.red {", "      color: red;", "    }", "  </style>", "<body>", "  <p>", "    Here is a blue link to <a class=\"blue\" href=\"/test-data/project01/href.html\">href.html</a>", "  </p>", "  <p>", "    And a red link to <a class=\"red\" href=\"/test-data/project01/simple.html\">simple.html</a>", "  </p>", "</body>", "</html>"}},
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
