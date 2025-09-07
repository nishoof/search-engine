package main

import (
	"bufio"
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

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
	} else if node.Type == html.ElementNode && (node.DataAtom == atom.Style || node.DataAtom == atom.Script) {
		return // skip the <style> and <script> subtrees
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
