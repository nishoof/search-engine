package crawler

import (
	"fmt"
	"testing"

	"github.com/nishoof/search-engine/index"
	"github.com/nishoof/search-engine/testutils"
)

func TestCrawl(t *testing.T) {
	ts := testutils.NewSimpleTestServer()
	defer ts.Close()

	seed := ts.URL + "/" + testutils.SimpleTestdataPaths[0]
	var idx index.Index = index.NewIndexInMemory()
	fmt.Printf("Crawling %s\n", seed)
	Crawl(seed, true, &idx)
	fmt.Printf("Crawled %s\n", seed)

	tests := []struct {
		word string
		url  string
		want int
	}{
		{"href", testutils.SimpleTestdataPaths[0], 1},
		{"272", testutils.SimpleTestdataPaths[1], 1},
		{"simpl", testutils.SimpleTestdataPaths[2], 2}, // "simple" stems to "simpl"
		{"style", testutils.SimpleTestdataPaths[3], 1},
		{"blue", testutils.SimpleTestdataPaths[3], 1},
		{"link", testutils.SimpleTestdataPaths[3], 2},
		{"red", testutils.SimpleTestdataPaths[3], 1},
		{"67", testutils.SimpleTestdataPaths[0], 0},
	}

	for _, test := range tests {
		got := idx.GetFrequency(test.word, ts.URL+"/"+test.url)
		if got != test.want {
			t.Errorf("For word %q and url %q, got %d but wanted %d\n", test.word, test.url, got, test.want)
		}
	}
}
