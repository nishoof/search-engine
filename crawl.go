package main

import (
	"bufio"
	"fmt"
	"net/url"
)

/* Crawls the website starting from the given seed URL and returns a slice of all crawled URLs */
func crawl(seed string) (map[string]struct{}, map[string]struct{}) {
	q := make([]string, 0)
	q = append(q, seed)
	visitedSet := make(map[string]struct{})
	wordsSet := make(map[string]struct{})
	host := extractHost(seed)
	if host == "" {
		return nil, nil
	}

	for len(q) > 0 {
		url := q[0]
		q = q[1:]
		visitedSet[url] = struct{}{}
		fmt.Printf("Crawling: %s\n", url)

		body := download(url)
		if body == nil {
			continue
		}
		defer body.Close()

		reader := bufio.NewReader(body)
		words, hrefs := extract(reader)
		if words == nil || hrefs == nil {
			continue
		}
		cleanedHrefs := cleanHrefs(host, hrefs)

		for _, word := range words {
			fmt.Printf("%s ", word)
		}
		fmt.Printf("\n\n")

		for _, word := range words {
			wordsSet[word] = struct{}{}
		}

		for _, href := range cleanedHrefs {
			_, visited := visitedSet[href]
			if !visited && extractHost(href) == host {
				q = append(q, href)
			}
		}
	}

	return visitedSet, wordsSet
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
