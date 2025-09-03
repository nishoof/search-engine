package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"unicode"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

/* Downloads the contents of the website of the given url and returns a ReadCloser on it */
func download(url string) io.ReadCloser {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error downloading %q: %v\n", url, err)
		return nil
	}
	return resp.Body
}

/* Extracts the words and hrefs from the contents of a website represented as a Reader, returning them as 2 maps (to avoid duplicates) */
func extract(reader *bufio.Reader) (map[string]struct{}, map[string]struct{}) {
	tree, err := html.Parse(reader)
	if err != nil {
		fmt.Printf("Error parsing HTML: %v\n", err)
		return nil, nil
	}

	words := make(map[string]struct{})
	hrefs := make(map[string]struct{})
	extractDfsHelper(tree, words, hrefs)
	return words, hrefs
}

/* Does a recursive dfs on the HTML node tree, extracting words and hrefs into the given sets, and skipping nodes that we don't want (such as style) */
func extractDfsHelper(node *html.Node, words map[string]struct{}, hrefs map[string]struct{}) {
	if node.Type == html.TextNode {
		extractWords(node, words)
	} else if node.Type == html.ElementNode && node.DataAtom == atom.A {
		extractHrefs(node, hrefs)
	} else if node.Type == html.ElementNode && (node.DataAtom == atom.Style || node.DataAtom == atom.Title || node.DataAtom == atom.Script) {
		return // skip the <style>, <title>, and <script> subtrees
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		extractDfsHelper(child, words, hrefs)
	}
}

/* Extracts words from the given text node and adds them to the given words set */
func extractWords(node *html.Node, words map[string]struct{}) {
	// Verify that the node is a text node
	if node.Type != html.TextNode {
		panic("Invalid node")
	}

	delimiterFunc := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	newWords := strings.FieldsFunc(node.Data, delimiterFunc)
	for _, word := range newWords {
		word = strings.ToLower(word)
		words[word] = struct{}{}
	}
}

/* Extracts href attributes from the given anchor node and adds them to the given hrefs set */
func extractHrefs(node *html.Node, hrefs map[string]struct{}) {
	// Verify that the node is an anchor element node
	if node.Type != html.ElementNode || node.DataAtom != atom.A {
		panic("Invalid node")
	}

	for _, attribute := range node.Attr {
		if attribute.Key == "href" {
			hrefs[attribute.Val] = struct{}{}
			break
		}
	}
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

		for word := range words {
			fmt.Printf("%s ", word)
		}
		fmt.Printf("\n\n")

		for word := range words {
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

func main() {
	_, words := crawl("https://usf-cs272-f25.github.io/")
	fmt.Printf("Extracted Words: \n")
	for word := range words {
		fmt.Printf("%s ", word)
	}
	fmt.Println()
}
