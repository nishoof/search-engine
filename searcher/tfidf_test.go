package searcher

import (
	"testing"

	"github.com/nishoof/search-engine/crawler"
	"github.com/nishoof/search-engine/index"
	"github.com/nishoof/search-engine/testutils"
)

func TestTfIdf(t *testing.T) {
	top10TestServer := testutils.NewTop10TestServer()
	defer top10TestServer.Close()

	var idx index.Index = index.NewIndexInMemory()
	crawler.Crawl(top10TestServer.URL+"/top10", true, &idx)

	tests := []struct {
		word, doc              string
		numWordsInDoc, numDocs int
		want                   float64
	}{
		{
			"blood", top10TestServer.URL + "/top10/Dracula%20%7C%20Project%20Gutenberg/chap10.html",
			1663, 331,
			0.016404221154672147,
		},
	}

	for testIdx, test := range tests {
		tfidf := tfidf(test.word, test.doc, test.numWordsInDoc, test.numDocs, idx)
		if tfidf != test.want {
			t.Errorf("test %d failed. got %f, want %f\n",
				testIdx, tfidf, test.want)
		}
	}
}
