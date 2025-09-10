package main

import (
	"fmt"
	"strings"
)

func search(seed, word string) FrequencyMap {
	word = strings.ToLower(word)
	mp := crawl(seed)

	ii := make(InvertedIndex)
	for url, words := range mp {
		for _, w := range words {
			ii.Increment(w, url)
		}
	}

	urlToFreq := ii[word]
	if urlToFreq == nil {
		println("Word not found")
		return nil
	}

	for url, freq := range urlToFreq {
		fmt.Printf("%s: %d\n", url, freq)
	}

	return urlToFreq
}
