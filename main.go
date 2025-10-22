package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

func startServer(indexType IndexType, fastMode bool) Index {
	// Use http.Dir to serve the contents of ./static for GET requests
	http.Handle("/", http.FileServer(http.Dir("./static")))
	go http.ListenAndServe(":8080", nil)

	// Crawl the top 10 pages and build the index
	fmt.Println("Crawling and building index...")
	var idx Index
	if indexType == IN_MEM {
		idx = NewIndexInMemory()
	} else {
		idx = NewIndexSQLite()
	}
	crawl("http://localhost:8080/top10", fastMode, &idx)
	// crawl("https://www.usfca.edu/", fastMode, &idx)

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
	indexFlag := flag.String("index", "", "Specify which index to use: inmem or sqlite")
	fastFlag := flag.Bool("fast", false, "Enable fast mode (ignores crawl-delay)")
	flag.Parse()

	var indexType IndexType
	switch *indexFlag {
	case "inmem":
		indexType = IN_MEM
	case "sqlite":
		indexType = SQLITE
	default:
		panic("Please specify a valid index with -index=inmem or -index=sqlite")
	}

	startServer(indexType, *fastFlag)
	for {
		time.Sleep(100 * time.Millisecond)
	}
}
