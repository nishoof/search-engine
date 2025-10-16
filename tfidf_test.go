package main

import (
	"testing"
)

func TestTfIdf(t *testing.T) {
	tests := []struct {
		word, doc              string
		numWordsInDoc, numDocs int
		want                   float64
	}{
		{
			"blood",
			"http://localhost:8080/top10/Dracula%20%7C%20Project%20Gutenberg/chap10.html",
			1663, 331, 0.016404221154672147,
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
