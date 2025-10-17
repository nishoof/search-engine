package main

import (
	"testing"
)

func TestCrawl(t *testing.T) {
	ts := getTestServer()
	defer ts.Close()

	seed := ts.URL + testPaths[0]
	var idx Index = NewIndexInMemory()
	crawl(seed, true, &idx)
	tests := []struct {
		word string
		url  string
		want int
	}{
		{"href", testPaths[0], 1},
		{"272", testPaths[1], 1},
		{"simpl", testPaths[2], 2}, // "simple" stems to "simpl"
		{"style", testPaths[3], 1},
		{"blue", testPaths[3], 1},
		{"link", testPaths[3], 2},
		{"red", testPaths[3], 1},
		{"67", testPaths[0], 0},
	}

	for _, test := range tests {
		got := idx.GetFrequency(test.word, ts.URL+test.url)
		if got != test.want {
			t.Errorf("For word %q and url %q, got %d but wanted %d\n", test.word, test.url, got, test.want)
		}
	}
}
