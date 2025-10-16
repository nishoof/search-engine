package main

import (
	"bufio"
	"fmt"
	"net/url"
	"time"
)

/* Crawls the website starting from the given seed URL and returns a map that maps a page URL to the words found on the page */
func crawl(seed string, ignoreCrawlDelay bool) map[string][]string {
	q := make([]string, 0)
	q = append(q, seed)
	visitedSet := make(map[string]struct{})
	mp := make(map[string][]string) // maps URL to list of words
	host := extractHost(seed)
	if host == "" {
		return nil
	}
	stopper := NewStopper() // used by the extract function

	robotsTxtUrl, err := url.JoinPath(host, "robots.txt")
	if err != nil {
		panic(err)
	}
	rules := parseRobotsTxt(robotsTxtUrl)
	if ignoreCrawlDelay {
		rules.SetCrawlDelay(0)
	}

	fmt.Printf("seed: %s\n", seed)
	fmt.Printf("host: %s\n", host)
	fmt.Printf("robotsTxtUrl: %s\n", robotsTxtUrl)

	for len(q) > 0 {
		url := q[0]
		q = q[1:]
		visitedSet[url] = struct{}{}

		if rules.Disallowed(url) {
			continue // skip urls disallowed by robots.txt
		}

		fmt.Printf("crawling %s\n", url)
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

		mp[url] = append(mp[url], words...)

		cleanedHrefs := cleanHrefs(host, hrefs)
		for _, href := range cleanedHrefs {
			_, visited := visitedSet[href]
			if !visited && extractHost(href) == host {
				q = append(q, href)
			}
		}
	}

	return mp
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
