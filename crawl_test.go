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
			ts.URL + testPaths[0]: {"simple", "href", "style"},
			ts.URL + testPaths[1]: {"272", "links"},
			ts.URL + testPaths[2]: {"simple", "simple"},
			ts.URL + testPaths[3]: {"style", "blue", "link", "href", "red", "link", "simple"},
		}},
	}

	for _, test := range tests {
		got := crawl(ts.URL+test.seed, true, nil)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("For seed %q, got %v but wanted %v\n", test.seed, got, test.want)
		}
	}
}
