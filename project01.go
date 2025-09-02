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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

/* Downloads the contents of the website of the given url and returns a ReadCloser on it */
func download(url string) io.ReadCloser {
	resp, err := http.Get(url)
	checkErr(err)
	return resp.Body
}

/* Extracts the words and hrefs from the contents of a website represented as a Reader, returning them as 2 slices */
func extract(reader *bufio.Reader) ([]string, []string) {
	tree, err := html.Parse(reader)
	checkErr(err)

	words := make([]string, 0)
	hrefs := make([]string, 0)
	isVisibleText := true
	for node := range tree.Descendants() {
		if node.Type == html.ElementNode {
			isVisibleText = node.DataAtom != atom.Style && node.DataAtom != atom.Title
		}
		if !isVisibleText {
			continue
		}
		if node.Type == html.TextNode {
			extractWords(node, &words)
		} else if node.Type == html.ElementNode && node.DataAtom == atom.A {
			extractHrefs(node, &hrefs)
		}
	}
	return words, hrefs
}

/* Extracts words from the given text node and appends them to the given words slice */
func extractWords(node *html.Node, words *[]string) {
	if node.Type != html.TextNode {
		panic("Invalid node")
	}

	delimiterFunc := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	newWords := strings.FieldsFunc(node.Data, delimiterFunc)
	*words = append(*words, newWords...)
}

/* Extracts href attributes from the given anchor node and appends them to the given hrefs slice */
func extractHrefs(node *html.Node, hrefs *[]string) {
	// Check if node is an anchor element node
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

/* Extracts the host from the given href */
func extractHost(href string) string {
	u, err := url.Parse(href)
	checkErr(err)
	u.Path = "/"
	return u.String()
}

/* Cleans the given href by resolving it against the given base url */
func cleanHref(base string, href string) string {
	b, err := url.Parse(base)
	checkErr(err)
	u, err := url.Parse(href)
	checkErr(err)
	cleaned := b.ResolveReference(u)
	return cleaned.String()
}

func main() {
	input := "https://usf-cs272-f25.github.io/test-data/project01/style.html"

	body := download(input)
	defer body.Close()

	reader := bufio.NewReader(body)
	words, hrefs := extract(reader)

	fmt.Println("Words:", strings.Join(words, ", "))
	fmt.Println("Hrefs:", strings.Join(hrefs, ", "))
	fmt.Println()

	host := extractHost(input)
	for _, href := range hrefs {
		fmt.Println(cleanHref(host, href))
	}
}
