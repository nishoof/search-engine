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
		{testData[1], []string{"272", "links"}, []string{}},
		{testData[2], []string{"simple", "simple"}, []string{"/test-data/search-engine/simple.html"}},
		{testData[3], []string{"style", "blue", "link", "href", "red", "link", "simple"}, []string{"/test-data/search-engine/href.html", "/test-data/search-engine/simple.html"}},
	}

	for testIdx, test := range tests {
		stopper := NewStopper()
		testFileStr := strings.Join(test.testData, "\n")
		stringsReader := strings.NewReader(testFileStr)
		bufioReader := bufio.NewReader(stringsReader)
		gotWords, gotHrefs := extract(bufioReader, stopper)

		if !reflect.DeepEqual(gotWords, test.wantWords) {
			t.Errorf("For test %d, got words %v but wanted %v\n", testIdx, gotWords, test.wantWords)
		}

		if !reflect.DeepEqual(gotHrefs, test.wantHrefs) {
			t.Errorf("For test %d, got hrefs %v but wanted %v\n", testIdx, gotHrefs, test.wantHrefs)
		}
	}
}
