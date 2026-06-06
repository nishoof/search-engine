package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/nishoof/search-engine/index"
)

func startServer(seed string, indexType index.IndexType, fastMode bool) index.Index {
	fmt.Println("Crawling and building index...")
	var idx index.Index
	if indexType == index.IN_MEM {
		idx = index.NewIndexInMemory()
	} else {
		sqliteIdx := index.NewIndexSQLite()
		idx = &sqliteIdx
	}
	crawl(seed, fastMode, &idx)

	fmt.Println("Done\nhttp://localhost:8080/")

	// Handle requests to / (show a search bar)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./templates/search-bar.html")
	})

	// Handle requests to /search (show the search results)
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")
		results := Search(q, idx)

		t, err := template.ParseFiles("./templates/results.html")
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			fmt.Println("Template error:", err)
			return
		}

		templateData := struct {
			Query   string
			Results []Result
		}{
			Query:   q,
			Results: results,
		}

		err = t.Execute(w, templateData)
		if err != nil {
			panic(err)
		}
	})

	go http.ListenAndServe(":8080", nil)

	return idx
}

func main() {
	indexFlag := flag.String("index", "", "Specify which index to use: inmem or sqlite")
	fastFlag := flag.Bool("fast", false, "Enable fast mode (ignores crawl-delay)")
	seedFlag := flag.String("seed", "", "URL to crawl (e.g. https://nishilanand.com/)")
	flag.Parse()

	if *seedFlag == "" {
		panic("Please specify a seed URL with -seed=<url>")
	}

	var indexType index.IndexType
	switch *indexFlag {
	case "inmem":
		indexType = index.IN_MEM
	case "sqlite":
		indexType = index.SQLITE
	default:
		panic("Please specify a valid index with -index=inmem or -index=sqlite")
	}

	startServer(*seedFlag, indexType, *fastFlag)
	for {
		time.Sleep(100 * time.Millisecond)
	}
}
