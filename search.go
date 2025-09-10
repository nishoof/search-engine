package main

import (
	"fmt"

	"github.com/kljensen/snowball"
)

func search(seed, word string) FrequencyMap {
	word, err := snowball.Stem(word, "english", true)
	if err != nil {
		fmt.Println("Error stemming word:", err)
		return nil
	}

	mp := crawl(seed)

	ii := make(InvertedIndex)
	for url, words := range mp {
		for _, w := range words {
			stemmed, err := snowball.Stem(w, "english", true)
			if err != nil {
				fmt.Println("Error stemming word:", err)
				continue
			}
			ii.Increment(stemmed, url)
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
