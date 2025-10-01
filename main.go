package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func startServer(indexType IndexType) Index {
	// Use http.Dir to serve the contents of ./static for GET requests
	http.Handle("/", http.FileServer(http.Dir("./static")))
	go http.ListenAndServe(":8080", nil)

	// Crawl the top 10 pages and build the index
	fmt.Println("Crawling...")
	// mp := crawl("http://localhost:8080/top10")
	mp := crawl("https://usf-cs272-f25.github.io/")

	fmt.Println("Building index...")
	var idx Index
	if indexType == IN_MEM {
		idx = NewIndexInMemory(mp)
	} else {
		idx = NewIndexSQLite(mp)
	}

	fmt.Println("Done\nhttp://localhost:8080/")

	// Handle /search requests
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")
		results := Search(q, idx)
		fmt.Fprintf(w, "search term: %s\n\n", q)
		for i, result := range results {
			fmt.Fprintf(w, "rank %3d\tscore: %f, occurrences: %d,\turl: %s\n", i+1, result.score, result.occurrences, result.url)
		}
	})

	return idx
}

func main() {
	if len(os.Args) < 2 {
		panic("Please identify which index should be used. For example, `-index=inmem` or `-index=sqlite`.")
	}

	var indexType IndexType
	switch os.Args[1] {
	case "-index=inmem":
		indexType = IN_MEM
	case "-index=sqlite":
		indexType = SQLITE
	default:
		panic("Unknown argument: " + os.Args[1])
	}

	startServer(indexType)
	for {
		time.Sleep(100 * time.Millisecond)
	}
}
