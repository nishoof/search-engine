package main

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	ts := getTestServer()
	defer ts.Close()

	tests := []struct {
		seed, word string
		want       FrequencyMap
	}{
		{ts.URL + "/test-data/rnj/sceneI_30.0.html", "Verona", map[string]int{ts.URL + "/test-data/rnj/sceneI_30.0.html": 1}},
		{ts.URL + "/test-data/rnj/sceneI_30.1.html", "Benvolio", map[string]int{ts.URL + "/test-data/rnj/sceneI_30.1.html": 26}},
		{ts.URL + "/test-data/rnj/", "Romeo", map[string]int{
			ts.URL + "/test-data/rnj/sceneI_30.0.html":  2,
			ts.URL + "/test-data/rnj/sceneI_30.1.html":  22,
			ts.URL + "/test-data/rnj/sceneI_30.3.html":  2,
			ts.URL + "/test-data/rnj/sceneI_30.4.html":  17,
			ts.URL + "/test-data/rnj/sceneI_30.5.html":  15,
			ts.URL + "/test-data/rnj/sceneII_30.2.html": 42,
			ts.URL + "/test-data/rnj/":                  200,
			ts.URL + "/test-data/rnj/sceneI_30.2.html":  15,
			ts.URL + "/test-data/rnj/sceneII_30.0.html": 3,
			ts.URL + "/test-data/rnj/sceneII_30.1.html": 10,
			ts.URL + "/test-data/rnj/sceneII_30.3.html": 13,
		}},
	}

	for testIdx, test := range tests {
		got := search(test.seed, test.word)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("test %d: got %v, want %v", testIdx, got, test.want)
		}
	}
}
