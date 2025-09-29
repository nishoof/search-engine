package main

import (
	"fmt"
	"net/http"
	"time"
)

func startServer() Index {
	// Use http.Dir to serve the contents of ./static for GET requests
	http.Handle("/", http.FileServer(http.Dir("./static")))
	go http.ListenAndServe(":8080", nil)

	// Crawl the top 10 pages and build the index
	fmt.Println("Crawling...")
	mp := crawl("http://localhost:8080/top10")
	fmt.Println("Building index...")
	idx := NewIndexInMemory(mp)
	fmt.Println("Done\nhttp://localhost:8080/")

	// Handle /search requests
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")
		results := idx.Search(q)
		fmt.Fprintf(w, "search term: %s\n\n", q)
		for i, result := range results {
			fmt.Fprintf(w, "rank %3d\tscore: %f, occurrences: %d,\turl: %s\n", i+1, result.score, result.occurrences, result.url)
		}
	})

	return idx
}

func main() {
	startServer()
	for {
		time.Sleep(100 * time.Millisecond)
	}
}
