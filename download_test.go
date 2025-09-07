package main

import (
	"bufio"
	"testing"
)

func TestDownload(t *testing.T) {
	tests := []struct {
		path string
		want []string
	}{
		{testPaths[0], testData[0]},
		{testPaths[1], testData[1]},
		{testPaths[2], testData[2]},
	}

	ts := getTestServer()
	defer ts.Close()

	for testIdx, test := range tests {
		got := download(ts.URL + test.path)

		// Read through the downloaded content line by line with a scanner and compare it to expected
		scanner := bufio.NewScanner(got)
		for i := 0; scanner.Scan(); i++ {
			textGot := scanner.Text()
			textWant := test.want[i]
			if textGot != textWant {
				t.Errorf("For test %d at line %d, got \"%s\" but wanted \"%s\"\n", testIdx, i, textGot, textWant)
			}
		}
	}
}
