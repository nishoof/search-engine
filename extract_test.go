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
		{testData[1], []string{"Hello", "CS", "272", "there", "are", "no", "links", "here"}, []string{}},
		{testData[2], []string{"For", "a", "simple", "example", "see", "html"}, []string{"/test-data/project01/simple.html"}},
		{testData[3], []string{"Here", "is", "a", "blue", "link", "to", "href", "html", "And", "red", "simple"}, []string{"/test-data/project01/href.html", "/test-data/project01/simple.html"}},
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
			if gotWords[word] != struct{}{} {
				t.Errorf("For test %d, word %q is missing from gotWords\n", testIdx, word)
			}
		}

		// hrefs
		if len(gotHrefs) != len(test.wantHrefs) {
			t.Errorf("For test %d, got %d hrefs but wanted %d\n", testIdx, len(gotHrefs), len(test.wantHrefs))
		}
		for _, href := range test.wantHrefs {
			if gotHrefs[href] != struct{}{} {
				t.Errorf("For test %d, href %q is missing from gotHrefs\n", testIdx, href)
			}
		}
	}
}
