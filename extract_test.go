package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestExtract(t *testing.T) {
	tests := []struct {
		testData, wantWords, wantHrefs []string
	}{
		{testData[1], []string{"hello", "cs", "272", "there", "are", "no", "links", "here"}, []string{}},
		{testData[2], []string{"for", "a", "simple", "example", "see", "simple", "html"}, []string{"/test-data/lab03/simple.html"}},
		{testData[3], []string{"style", "here", "is", "a", "blue", "link", "to", "href", "html", "and", "a", "red", "link", "to", "simple", "html"}, []string{"/test-data/lab03/href.html", "/test-data/lab03/simple.html"}},
	}

	for testIdx, test := range tests {
		testFileStr := strings.Join(test.testData, "\n")
		stringsReader := strings.NewReader(testFileStr)
		bufioReader := bufio.NewReader(stringsReader)
		gotWords, gotHrefs := extract(bufioReader)

		if !reflect.DeepEqual(gotWords, test.wantWords) {
			t.Errorf("For test %d, got words %v but wanted %v\n", testIdx, gotWords, test.wantWords)
		}

		if !reflect.DeepEqual(gotHrefs, test.wantHrefs) {
			t.Errorf("For test %d, got hrefs %v but wanted %v\n", testIdx, gotHrefs, test.wantHrefs)
		}
	}
}
