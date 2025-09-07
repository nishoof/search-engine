package main

import (
	"fmt"
)

func main() {
	_, words := crawl("https://usf-cs272-f25.github.io/")
	fmt.Printf("Extracted Words: \n")
	for word := range words {
		fmt.Printf("%s ", word)
	}
	fmt.Println()
}
