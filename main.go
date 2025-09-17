package main

import (
	"fmt"
	"net/http"
	"time"
)

func startServer() {
	// Use http.Dir to serve the contents of ./static for GET requests
	http.Handle("/", http.FileServer(http.Dir("./static")))
	go http.ListenAndServe(":8080", nil)

	// Crawl the top 10 pages and build the inverted index
	fmt.Println("Crawling...")
	mp := crawl("http://localhost:8080/top10")
	fmt.Println("Building inverted index...")
	ii := NewInvertedIndex(mp)
	fmt.Println("Done")

	// Handle /search requests
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")
		results := ii.Search(q)
		fmt.Fprintf(w, "search term: %s\n\n", q)
		for i, result := range results {
			fmt.Fprintf(w, "rank %3d\tscore: %f, url: %s\n", i+1, result.score, result.url)
		}
	})
}

func main() {
	startServer()
	for {
		time.Sleep(100 * time.Millisecond)
	}
}
