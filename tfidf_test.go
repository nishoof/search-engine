package main

import (
	"testing"
)

func TestTfIdf(t *testing.T) {
	tests := []struct {
		word, doc              string
		numWordsInDoc, numDocs int
		invertedIndex          InvertedIndex
		want                   float64
	}{
		{
			"simple", "index", 3, 1, InvertedIndex{
				map[string]map[string]int{"simple": {"index": 1}},
				map[string]int{"index": 3},
			}, 0.16666666666666667,
		},
	}

	for testIdx, test := range tests {
		tfidf := tfidf(test.word, test.doc, test.numWordsInDoc, test.numDocs, &test.invertedIndex)
		if tfidf != test.want {
			t.Errorf("test %d failed. got %f, want %f\n",
				testIdx, tfidf, test.want)
		}
	}
}
