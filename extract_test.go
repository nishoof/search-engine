package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestExtract(t *testing.T) {
	tests := []struct {
		testData, wantWords, wantHrefs []string
	}{
		{testData[1], []string{"hello", "cs", "272", "there", "are", "no", "links", "here"}, []string{}},
		{testData[2], []string{"for", "a", "simple", "example", "see", "html"}, []string{"/test-data/lab03/simple.html"}},
		{testData[3], []string{"simple", "here", "is", "a", "blue", "link", "to", "href", "html", "and", "red", "simple"}, []string{"/test-data/lab03/href.html", "/test-data/lab03/simple.html"}},
	}

	for testIdx, test := range tests {
		testFileStr := strings.Join(test.testData, "\n")
		stringsReader := strings.NewReader(testFileStr)
		bufioReader := bufio.NewReader(stringsReader)
		gotWords, gotHrefs := extract(bufioReader)

		// words
		if len(gotWords) != len(test.wantWords) {
			t.Errorf("For test %d, got %d words but wanted %d\n", testIdx, len(gotWords), len(test.wantWords))
		}
		for _, word := range test.wantWords {
			_, exists := gotWords[word]
			if !exists {
				t.Errorf("For test %d, word %q is missing from gotWords\n", testIdx, word)
			}
		}

		// hrefs
		if len(gotHrefs) != len(test.wantHrefs) {
			t.Errorf("For test %d, got %d hrefs but wanted %d\n", testIdx, len(gotHrefs), len(test.wantHrefs))
		}
		for _, href := range test.wantHrefs {
			_, exists := gotHrefs[href]
			if !exists {
				t.Errorf("For test %d, href %q is missing from gotHrefs\n", testIdx, href)
			}
		}
	}
}
