package main

import (
	"bufio"
	"fmt"
	"net/url"
	"time"
)

/* Crawls the website starting from the given seed URL, then crawling all links found on that page, and so on for links found on those pages. fastMode ignores crawl-delay and prints less. If idx is not nil, crawl will also build the index using the index's Increment method */
func crawl(seed string, fastMode bool, idx *Index) {
	startTime := time.Now()

	q := make([]string, 0)
	q = append(q, seed)
	visitedSet := make(map[string]struct{})
	host := extractHost(seed)
	if host == "" {
		return
	}
	stopper := NewStopper() // used by the extract function

	robotsTxtUrl, err := url.JoinPath(host, "robots.txt")
	if err != nil {
		panic(err)
	}
	rules := parseRobotsTxt(robotsTxtUrl)
	if fastMode {
		rules.SetCrawlDelay(0)
	}

	for len(q) > 0 {
		url := q[0]
		q = q[1:]
		visitedSet[url] = struct{}{}

		if rules.Disallowed(url) {
			continue // skip urls disallowed by robots.txt
		}

		if !fastMode {
			fmt.Printf("crawling %s\n", url)
		}

		body := download(url)
		time.Sleep(rules.crawlDelay)
		if body == nil {
			continue
		}
		defer body.Close()

		reader := bufio.NewReader(body)
		words, hrefs := extract(reader, stopper)
		if words == nil || hrefs == nil {
			continue
		}

		if idx != nil {
			for _, word := range words {
				(*idx).Increment(word, url)
			}
		}

		cleanedHrefs := cleanHrefs(host, hrefs)
		for _, href := range cleanedHrefs {
			_, visited := visitedSet[href]
			if !visited && extractHost(href) == host {
				q = append(q, href)
			}
		}
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime).Seconds()
	numUrls := len(visitedSet)
	fmt.Printf("crawled %d urls in %.2f seconds (%.2f per second)\n", numUrls, duration, (float64)(numUrls)/duration)
}

/* Extracts the host from the given href */
func extractHost(href string) string {
	u, err := url.Parse(href)
	if err != nil {
		fmt.Printf("Error parsing href %q: %v\n", href, err)
		return ""
	}
	u.Path = "/"
	return u.String()
}
