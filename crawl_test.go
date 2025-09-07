package main

import "testing"

func TestCrawl(t *testing.T) {
	tests := []struct {
		seed string
		want []string
	}{
		{testPaths[0], []string{
			testPaths[0],
			testPaths[1],
			testPaths[2],
			testPaths[3],
		}},
	}

	ts := getTestServer()
	defer ts.Close()

	for _, test := range tests {
		got, _ := crawl(ts.URL + test.seed)
		if len(got) != len(test.want) {
			t.Errorf("For seed %q, got %d URLs but wanted %d\n", test.seed, len(got), len(test.want))
		}
		count := 0
		for _, v := range test.want {
			_, exists := got[ts.URL+v]
			if exists {
				count++
			}
		}
		if count != len(test.want) {
			t.Errorf("For seed %q, got %v but wanted %v\n", test.seed, got, test.want)
		}
	}
}
