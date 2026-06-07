package main

import (
	"flag"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"time"

	"github.com/nishoof/search-engine/crawler"
	"github.com/nishoof/search-engine/index"
	"github.com/nishoof/search-engine/logger"
	"github.com/nishoof/search-engine/searcher"
)

func startServer(seed string, indexType index.IndexType, fastMode bool) index.Index {
	slog.Info("Crawling and building index...")
	var idx index.Index
	if indexType == index.IN_MEM {
		idx = index.NewIndexInMemory()
	} else {
		sqliteIdx := index.NewIndexSQLite()
		idx = &sqliteIdx
	}
	crawlStartTime := time.Now()
	crawler.Crawl(seed, fastMode, &idx)

	numUrls := idx.GetNumDocs()
	durationSeconds := time.Since(crawlStartTime).Seconds()
	slog.Info(fmt.Sprintf("Done! Crawled %d urls in %.2f seconds (%.2f per second)", numUrls, durationSeconds, (float64)(numUrls)/durationSeconds))

	// Handle requests to / (show a search bar)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./templates/search-bar.html")
	})

	// Handle requests to /search (show the search results)
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")
		results := searcher.Search(q, idx)

		t, err := template.ParseFiles("./templates/results.html")
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			slog.Error("Template error", "error", err)
			return
		}

		templateData := struct {
			Query   string
			Results []searcher.Result
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
	slog.Info("Search at http://localhost:8080")

	return idx
}

func main() {
	seedFlag := flag.String("seed", "", "URL to crawl (e.g. https://nishilanand.com/)")
	indexFlag := flag.String("index", "", "Specify which index to use: inmem or sqlite")
	verboseFlag := flag.Bool("v", false, "Enable verbose logging (debug level)")
	fastFlag := flag.Bool("fast", false, "Enable fast mode (ignores crawl-delay)")
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

	if *verboseFlag {
		logger.Level.Set(slog.LevelDebug)
	} else {
		logger.Level.Set(slog.LevelInfo)
	}

	startServer(*seedFlag, indexType, *fastFlag)
	for {
		time.Sleep(100 * time.Second)
	}
}
