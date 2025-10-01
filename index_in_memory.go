package main

import (
	"github.com/kljensen/snowball"
)

type IndexInMemory struct {
	frequency map[string]map[string]int // maps words to their FrequencyMap
	wordCount map[string]int            // maps document names to their word counts
}

func NewIndexInMemory(mp map[string][]string) IndexInMemory {
	idx := new(IndexInMemory)
	idx.frequency = make(map[string]map[string]int)
	idx.wordCount = make(map[string]int)
	for url, words := range mp {
		for _, w := range words {
			stemmed, err := snowball.Stem(w, "english", true)
			if err != nil {
				panic(err)
			}
			idx.Increment(stemmed, url)
		}
	}
	return *idx
}

func (idx IndexInMemory) GetDocs() []string {
	// the document names are the keys of the wordCount map
	documentNames := make([]string, len(idx.wordCount))
	i := 0
	for documentName := range idx.wordCount {
		documentNames[i] = documentName
		i++
	}
	return documentNames
}

func (idx IndexInMemory) GetFrequency(word, documentName string) int {
	fm, exists := idx.frequency[word]
	if !exists {
		return 0
	}
	return fm[documentName]
}

func (idx IndexInMemory) GetNumDocs() int {
	return len(idx.wordCount)
}

func (idx IndexInMemory) GetNumDocsWithWord(word string) int {
	fm, exists := idx.frequency[word]
	if !exists {
		return 0
	}
	return len(fm)
}

func (idx IndexInMemory) GetWordCount(documentName string) int {
	return idx.wordCount[documentName]
}

func (idx IndexInMemory) Increment(word, documentName string) {
	if _, exists := idx.frequency[word]; !exists {
		idx.frequency[word] = make(map[string]int)
	}
	idx.frequency[word][documentName]++
	idx.wordCount[documentName]++
}
