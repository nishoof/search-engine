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
				"simple": {"index": 1},
			}, -0.23104906018664842,
		},
	}

	for testIdx, test := range tests {
		tfidf := tfidf(test.word, test.doc, test.numWordsInDoc, test.numDocs, test.invertedIndex)
		if tfidf != test.want {
			t.Errorf("test %d failed. got %f, want %f\n",
				testIdx, tfidf, test.want)
		}
	}
}
