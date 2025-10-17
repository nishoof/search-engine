package main

import (
	"strings"
	"testing"
	"time"
)

func TestCrawlDelay(t *testing.T) {
	t1 := time.Now()
	crawl("http://localhost:8080/top10/Dracula%20%7C%20Project%20Gutenberg/index.html", false, nil)
	t2 := time.Now()
	if t2.Sub(t1) < (10 * time.Second) {
		t.Errorf("TestDisallow was too fast\n")
	}
}

func TestDisallow(t *testing.T) {
	got := Search("blood", idx)
	disallow := "chap21.html"

	for _, result := range got {
		if strings.Contains(result.url, disallow) {
			t.Errorf("%s is in results when searching \"blood\"but it contains %s\n", result.url, disallow)
		}
	}
}
