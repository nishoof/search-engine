package main

func tf(word, doc string, numWordsInDoc int, idx Index) float64 {
	if numWordsInDoc == 0 {
		return 0.0
	}
	occurrences := idx.GetFrequency(word, doc)
	return float64(occurrences) / float64(numWordsInDoc)
}

func idf(word string, numDocs int, idx Index) float64 {
	numDocsWithWord := idx.GetNumDocsWithWord(word)
	return float64(numDocs) / (float64(numDocsWithWord) + 1.0) // not using log
}

func tfidf(word, doc string, numWordsInDoc, numDocs int, idx Index) float64 {
	tf := tf(word, doc, numWordsInDoc, idx)
	idf := idf(word, numDocs, idx)
	return tf * idf
}
