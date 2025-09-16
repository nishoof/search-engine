package main

import (
	"math"
)

func tf(word, doc string, numWordsInDoc int, ii InvertedIndex) float64 {
	if numWordsInDoc == 0 {
		return 0.0
	}
	occurrences := ii.GetFrequency(word, doc)
	return float64(occurrences) / float64(numWordsInDoc)
}

func idf(word string, numDocs int, ii InvertedIndex) float64 {
	numDocsWithWord := ii.GetNumDocsWithWord(word)
	return math.Log(float64(numDocs) / (float64(numDocsWithWord) + 1.0))
}

func tfidf(word, doc string, numWordsInDoc, numDocs int, ii InvertedIndex) float64 {
	tf := tf(word, doc, numWordsInDoc, ii)
	idf := idf(word, numDocs, ii)
	return tf * idf
}
