package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		word string
		want Results
	}{
		{
			"Romeo",
			Results{
				{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneII_30.2.html", 41, 0.66329426},
				{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneII_30.4.html", 41, 0.65633315},
				{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneII_30.1.html", 9, 0.5824047},
				{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIII_30.2.html", 26, 0.5160091},
				{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneV_30.2.html", 5, 0.46423563},
				{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIII_30.1.html", 30, 0.45115855},
				{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneI_30.2.html", 14, 0.45025262},
				{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIII_30.3.html", 27, 0.44489247},
				{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIV_30.3.html", 8, 0.44489247},
				{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneI_30.4.html", 16, 0.41365305},
				{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneII_30.6.html", 5, 0.38407984},
				{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneII_30.0.html", 2, 0.35591397},
				{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneV_30.1.html", 10, 0.35239008},
				{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneII_30.3.html", 12, 0.3407687},
				{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneI_30.5.html", 14, 0.3101325},
				{"http://localhost:8080/top10", 1, 0.28098473},
				{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneI_30.1.html", 21, 0.27378},
				{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneV_30.3.html", 26, 0.25168893},
				{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIII_30.5.html", 16, 0.20884928},
				{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneI_30.0.html", 1, 0.16683468},
				{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIV_30.1.html", 5, 0.124156035},
				{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneII_30.5.html", 3, 0.12179566},
				{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIII_30.4.html", 1, 0.08680829},
				{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIV_30.4.html", 1, 0.0773726},
				{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/index.html", 10, 0.069650486},
				{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIV_30.2.html", 1, 0.06471163},
				{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneI_30.3.html", 1, 0.028397392},
				{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20The%20Picture%20of%20Dorian%20Gray,%20by%20Oscar%20Wilde/chap04.html", 4, 0.024744889},
				{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20The%20Picture%20of%20Dorian%20Gray,%20by%20Oscar%20Wilde/chap07.html", 3, 0.021909889},
				{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIV_30.5.html", 1, 0.020070337},
			},
		},
	}

	for testIdx, test := range tests {
		got := Search(test.word, idx)

		// Make sure we got the expected number of results
		if len(got) != len(test.want) {
			t.Errorf("Test %d: Got %d results but wanted %d\n", testIdx, len(got), len(test.want))
		}

		// Convert the wanted Result slice to a Result set
		wantSet := make(map[Result]struct{})
		for _, result := range test.want {
			wantSet[result] = struct{}{}
		}

		// Make sure each result is in the wanted set
		for _, result := range got {
			_, exists := wantSet[result]
			if !exists {
				t.Errorf("Test %d: got unexpected result %v\n", testIdx, result)
			}
		}
	}
}
