package searcher

import (
	"log/slog"
	"sort"

	"github.com/kljensen/snowball"
	"github.com/nishoof/search-engine/index"
)

type Result struct {
	URL         string
	Occurrences int
	Score       float32
	Title       string
}

type Results []Result

func Search(word string, idx index.Index) Results {
	// Stem the search word
	word, err := snowball.Stem(word, "english", true)
	if err != nil {
		slog.Error("Error stemming word", "word", word, "error", err)
		return nil
	}

	// Calculate TF-IDF for each document
	numDocs := idx.GetNumDocs()
	results := make(Results, 0)
	for _, url := range idx.GetDocs() {
		occurrences := idx.GetFrequency(word, url)
		if occurrences == 0 {
			continue
		}
		numWordsInDoc := idx.GetWordCount(url)
		tfidf := tfidf(word, url, numWordsInDoc, numDocs, idx)
		docTitle := idx.GetTitle(url)
		results = append(results, Result{url, occurrences, float32(tfidf), docTitle})
	}

	// Sort results by score
	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	return results
}
