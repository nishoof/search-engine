package main

import "fmt"

//-- FrequencyMap

type FrequencyMap map[string]int // maps document names to their frequency counts

func (fm FrequencyMap) GetFrequency(documentName string) int {
	return fm[documentName]
}

func (fm FrequencyMap) Print() {
	for doc, freq := range fm {
		fmt.Println(doc, freq)
	}
}

//-- InvertedIndex

type InvertedIndex map[string]FrequencyMap // maps words to their FrequencyMap

func (ii InvertedIndex) GetFrequency(word, documentName string) int {
	fm, exists := ii[word]
	if !exists {
		return 0
	}
	return fm.GetFrequency(documentName)
}

func (ii InvertedIndex) GetNumDocsWithWord(word string) int { // TODO: for git, this is used in tfidf
	fm, exists := ii[word]
	if !exists {
		return 0
	}
	return len(fm)
}

func (ii InvertedIndex) Increment(word, documentName string) {
	if _, exists := ii[word]; !exists {
		ii[word] = make(FrequencyMap)
	}
	ii[word][documentName]++
}

func (ii InvertedIndex) Print() {
	for word, fm := range ii {
		fmt.Printf("----- Word: %s\n", word)
		fm.Print()
		fmt.Println()
	}
}
