package main

import (
	"testing"

	"github.com/nishoof/search-engine/index"
)

func TestCrawl(t *testing.T) {
	tsURL := simpleTestServer.URL

	seed := tsURL + "/" + simpleTestdataPaths[0]
	var idx index.Index = index.NewIndexInMemory()
	crawl(seed, true, &idx)

	tests := []struct {
		word string
		url  string
		want int
	}{
		{"href", simpleTestdataPaths[0], 1},
		{"272", simpleTestdataPaths[1], 1},
		{"simpl", simpleTestdataPaths[2], 2}, // "simple" stems to "simpl"
		{"style", simpleTestdataPaths[3], 1},
		{"blue", simpleTestdataPaths[3], 1},
		{"link", simpleTestdataPaths[3], 2},
		{"red", simpleTestdataPaths[3], 1},
		{"67", simpleTestdataPaths[0], 0},
	}

	for _, test := range tests {
		got := idx.GetFrequency(test.word, tsURL+"/"+test.url)
		if got != test.want {
			t.Errorf("For word %q and url %q, got %d but wanted %d\n", test.word, test.url, got, test.want)
		}
	}
}
