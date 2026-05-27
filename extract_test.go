package main

import (
	"bufio"
	"bytes"
	"os"
	"reflect"
	"testing"
)

func TestExtract(t *testing.T) {
	tests := []struct {
		filepath  string
		wantWords map[string]int
		wantHrefs []string
		wantTitle string
	}{
		{simpleTestdataPaths[1], map[string]int{"272": 1, "link": 1}, []string{}, ""},
		{simpleTestdataPaths[2], map[string]int{"simpl": 2}, []string{"/testdata/simple/simple.html"}, ""},
		{simpleTestdataPaths[3], map[string]int{"style": 1, "blue": 1, "link": 2, "href": 1, "red": 1, "simpl": 1}, []string{"/testdata/simple/href.html", "/testdata/simple/simple.html"}, "Style"},
	}

	for testIdx, test := range tests {
		content, err := os.ReadFile(test.filepath)
		if err != nil {
			t.Fatalf("Unable to load test data for test %d: %v", testIdx, err)
		}

		stopper := NewStopper()
		bufioReader := bufio.NewReader(bytes.NewReader(content))
		gotWords, gotHrefs, gotTitle := extract(bufioReader, stopper)

		if !reflect.DeepEqual(gotWords, test.wantWords) {
			t.Errorf("For test %d, got words %v but wanted %v\n", testIdx, gotWords, test.wantWords)
		}
		if !reflect.DeepEqual(gotHrefs, test.wantHrefs) {
			t.Errorf("For test %d, got hrefs %v but wanted %v\n", testIdx, gotHrefs, test.wantHrefs)
		}
		if gotTitle != test.wantTitle {
			t.Errorf("For test %d, got title %q but wanted %q\n", testIdx, gotTitle, test.wantTitle)
		}
	}
}
