package main

import (
	"fmt"
	"net/url"
)

/* Cleans the given href by resolving it against the given base url */
func cleanHref(base string, href string) string {
	b, err := url.Parse(base)
	if err != nil {
		fmt.Printf("Error parsing base URL %q: %v\n", base, err)
		return ""
	}
	u, err := url.Parse(href)
	if err != nil {
		fmt.Printf("Error parsing href URL %q: %v\n", href, err)
		return ""
	}
	cleaned := b.ResolveReference(u)
	return cleaned.String()
}

/* Calls cleanHref() on each of the given hrefs and returns the cleaned hrefs in a slice */
func cleanHrefs(base string, hrefs map[string]struct{}) []string {
	cleanedSlice := make([]string, 0, len(hrefs))
	for href := range hrefs {
		clean := cleanHref(base, href)
		if clean != "" {
			cleanedSlice = append(cleanedSlice, clean)
		}
	}
	return cleanedSlice
}
