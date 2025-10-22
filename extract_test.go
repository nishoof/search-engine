package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestExtract(t *testing.T) {
	tests := []struct {
		testData  []string
		wantWords map[string]int
		wantHrefs []string
	}{
		{testData[1], map[string]int{"272": 1, "link": 1}, []string{}},
		{testData[2], map[string]int{"simpl": 2}, []string{"/test-data/search-engine/simple.html"}},
		{testData[3], map[string]int{"style": 1, "blue": 1, "link": 2, "href": 1, "red": 1, "simpl": 1}, []string{"/test-data/search-engine/href.html", "/test-data/search-engine/simple.html"}},
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
