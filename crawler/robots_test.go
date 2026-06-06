package crawler

import (
	"strings"
	"testing"
	"time"

	"github.com/nishoof/search-engine/index"
	"github.com/nishoof/search-engine/searcher"
	"github.com/nishoof/search-engine/testutils"
)

func TestCrawlDelay(t *testing.T) {
	ts := testutils.NewTop10TestServer()
	defer ts.Close()
	seed := ts.URL + "/top10/Dracula%20%7C%20Project%20Gutenberg/index.html"

	t1 := time.Now()
	Crawl(seed, false, nil)
	t2 := time.Now()

	if t2.Sub(t1) < (10 * time.Second) {
		t.Errorf("Crawling was too fast\n")
	}
}

func TestDisallow(t *testing.T) {
	ts := testutils.NewTop10TestServer()
	defer ts.Close()
	var idx index.Index = index.NewIndexInMemory()
	Crawl(ts.URL+"/top10", true, &idx)

	got := searcher.Search("blood", idx)
	disallow := "chap21.html"

	for _, result := range got {
		if strings.Contains(result.URL, disallow) {
			t.Errorf("%s is in results when searching \"blood\"but it contains %s\n", result.URL, disallow)
		}
	}
}
