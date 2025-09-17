package main

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		ii   InvertedIndex
		want Results
	}{
		{
			InvertedIndex{
				map[string]map[string]int{
					"word1": {
						"doc1": 3,
						"doc2": 5,
						"doc3": 1,
						"doc4": 5,
					},
					"word2": {
						"doc1": 2,
					},
				},
				map[string]int{
					"doc1": 100,
					"doc2": 150,
					"doc3": 200,
					"doc4": 10,
				},
			},
			Results{
				Result{url: "doc4", occurrences: 5, score: 0.4},
				Result{url: "doc2", occurrences: 5, score: 0.026666667},
				Result{url: "doc1", occurrences: 3, score: 0.024},
				Result{url: "doc3", occurrences: 1, score: 0.004},
			},
		},
	}

	for testIdx, test := range tests {
		got := test.ii.Search("word1")
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Test %d: got %v but wanted %v\n", testIdx, got, test.want)
		}
	}
}
