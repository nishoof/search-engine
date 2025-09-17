package main

import (
	"fmt"
	"sort"

	"github.com/kljensen/snowball"
)

//-- InvertedIndex

type Result struct {
	url         string
	occurrences int
	score       float32
}

type Results []Result

type InvertedIndex struct {
	frequency map[string]map[string]int // maps words to their FrequencyMap
	wordCount map[string]int            // maps document names to their word counts
}

func NewInvertedIndex(mp map[string][]string) *InvertedIndex {
	ii := new(InvertedIndex)
	ii.frequency = make(map[string]map[string]int)
	ii.wordCount = make(map[string]int)
	for url, words := range mp {
		for _, w := range words {
			stemmed, err := snowball.Stem(w, "english", true)
			if err != nil {
				fmt.Println("Error stemming word:", err)
				return nil
			}
			ii.Increment(stemmed, url)
			ii.wordCount[url]++
		}
	}
	return ii
}

func (ii *InvertedIndex) GetFrequency(word, documentName string) int {
	fm, exists := ii.frequency[word]
	if !exists {
		return 0
	}
	return fm[documentName]
}

func (ii *InvertedIndex) GetNumDocs() int {
	return len(ii.wordCount)
}

func (ii *InvertedIndex) GetNumDocsWithWord(word string) int {
	fm, exists := ii.frequency[word]
	if !exists {
		return 0
	}
	return len(fm)
}

func (ii *InvertedIndex) GetWordCount(documentName string) int {
	return ii.wordCount[documentName]
}

func (ii *InvertedIndex) Increment(word, documentName string) {
	if _, exists := ii.frequency[word]; !exists {
		ii.frequency[word] = make(map[string]int)
	}
	ii.frequency[word][documentName]++
	ii.wordCount[documentName]++
}

func (ii *InvertedIndex) Search(word string) Results {
	// Stem the search word
	word, err := snowball.Stem(word, "english", true)
	if err != nil {
		fmt.Println("Error stemming word:", err)
		return nil
	}

	// Calculate TF-IDF for each document
	numDocs := ii.GetNumDocs()
	results := make(Results, 0)
	for url := range ii.wordCount {
		occurrences := ii.GetFrequency(word, url)
		if occurrences == 0 {
			continue
		}
		tfidf := tfidf(word, url, ii.GetWordCount(url), numDocs, ii)
		results = append(results, Result{url, occurrences, float32(tfidf)})
	}

	// Sort results by score
	sort.Slice(results, func(i, j int) bool {
		return results[i].score > results[j].score
	})

	return results
}
