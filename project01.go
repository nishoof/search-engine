package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
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

/* Extracts the words and hrefs from the contents of a website represented as a Reader */
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

/* Extracts words from the given text node and adds them to the words slice */
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

/* Extracts href attributes from the given anchor node and adds them to the hrefs slice */
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

func main() {
	body := download("https://usf-cs272-f25.github.io/test-data/project01/style.html")
	defer body.Close()

	reader := bufio.NewReader(body)
	words, hrefs := extract(reader)

	fmt.Println("Words:", strings.Join(words, ", "))
	fmt.Println("Hrefs:", strings.Join(hrefs, ", "))
}
