package main

import (
	"fmt"
	"sort"

	"github.com/kljensen/snowball"
)

type Result struct {
	url         string
	occurrences int
	score       float32
}

type Results []Result

func search(seed, word string) Results {
	// Stem the search word
	word, err := snowball.Stem(word, "english", true)
	if err != nil {
		fmt.Println("Error stemming word:", err)
		return nil
	}

	// Crawl the seed URL, giving us a map of URLs to words at that url
	mp := crawl(seed)

	// Build the inverted index from the crawled data
	// Build a map from document names to their word counts
	ii := make(InvertedIndex)
	wordCount := make(map[string]int)
	for url, words := range mp {
		for _, w := range words {
			stemmed, err := snowball.Stem(w, "english", true)
			if err != nil {
				fmt.Println("Error stemming word:", err)
				return nil
			}
			ii.Increment(stemmed, url)
			wordCount[url]++
		}
	}

	// Calculate TF-IDF for each document
	numDocs := len(mp)
	results := make(Results, 0)
	for url := range mp {
		occurrences := ii.GetFrequency(word, url)
		if occurrences == 0 {
			continue
		}
		tfidf := tfidf(word, url, wordCount[url], numDocs, ii)
		results = append(results, Result{url, occurrences, float32(tfidf)})
	}

	// Sort results by score
	sort.Slice(results, func(i, j int) bool {
		return results[i].score > results[j].score
	})

	return results
}
