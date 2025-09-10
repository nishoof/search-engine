package main

import (
	"reflect"
	"testing"
)

func TestCrawl(t *testing.T) {
	ts := getTestServer()
	defer ts.Close()

	tests := []struct {
		seed string
		want map[string][]string
	}{
		{testPaths[0], map[string][]string{
			ts.URL + testPaths[0]: {"simple", "html", "href", "html", "style", "html"},
			ts.URL + testPaths[1]: {"hello", "cs", "272", "there", "are", "no", "links", "here"},
			ts.URL + testPaths[2]: {"for", "a", "simple", "example", "see", "simple", "html"},
			ts.URL + testPaths[3]: {"style", "here", "is", "a", "blue", "link", "to", "href", "html", "and", "a", "red", "link", "to", "simple", "html"},
		}},
	}

	for _, test := range tests {
		got := crawl(ts.URL + test.seed)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("For seed %q, got %v but wanted %v\n", test.seed, got, test.want)
		}
	}
}
