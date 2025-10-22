package main

import (
	"bufio"
	"fmt"
	"strings"
	"unicode"

	"github.com/kljensen/snowball"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

/* Extracts the relevant words and hrefs from the contents of a website represented as a Reader, returning words as a map from the word to its count and hrefs as a slice */
func extract(reader *bufio.Reader, stopper *Stopper) (map[string]int, []string) {
	tree, err := html.Parse(reader)
	if err != nil {
		fmt.Printf("Error parsing HTML: %v\n", err)
		return nil, nil
	}

	words := make(map[string]int, 0)
	hrefs := make([]string, 0)
	extractDfsHelper(tree, &words, &hrefs, stopper)
	return words, hrefs
}

/* Does a recursive dfs on the HTML node tree, extracting relevant words and hrefs into the given structures, and skipping nodes that we don't want (such as style) */
func extractDfsHelper(node *html.Node, words *map[string]int, hrefs *[]string, stopper *Stopper) {
	if node.Type == html.TextNode {
		extractWords(node, words, stopper)
	} else if node.Type == html.ElementNode && node.DataAtom == atom.A {
		extractHrefs(node, hrefs)
	} else if node.Type == html.ElementNode && (node.DataAtom == atom.Style || node.DataAtom == atom.Script) {
		return // skip the <style> and <script> subtrees
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		extractDfsHelper(child, words, hrefs, stopper)
	}
}

/* Extracts relevant words from the given text node and adds them to the given words map */
func extractWords(node *html.Node, words *map[string]int, stopper *Stopper) {
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
		if stopper.isStopWord(word) {
			continue
		}
		stemmed, err := snowball.Stem(word, "english", true)
		if err != nil {
			panic(err)
		}
		(*words)[stemmed]++
	}
}

/* Extracts href attributes from the given anchor node and adds them to the given hrefs slice */
func extractHrefs(node *html.Node, hrefs *[]string) {
	// Verify that the node is an anchor element node
	if node.Type != html.ElementNode || node.DataAtom != atom.A {
		panic("Invalid node")
	}

	for _, attribute := range node.Attr {
		if attribute.Key == "href" {
			*hrefs = append(*hrefs, attribute.Val)
			break
		}
	}
}
