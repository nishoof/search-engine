package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		ii   *InvertedIndex
		want Results
	}{
		{
			startServer(),
			Results{
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneII_30.3.html", 107, 0.25268582},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneII_30.1.html", 107, 0.25268582},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneII_30.5.html", 107, 0.25268582},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneII_30.6.html", 107, 0.25268582},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneII_30.2.html", 107, 0.25268582},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneII_30.0.html", 107, 0.25268582},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneII_30.4.html", 107, 0.25268582},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIII_30.3.html", 96, 0.18465991},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIII_30.1.html", 96, 0.18465991},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIII_30.4.html", 96, 0.18465991},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIII_30.2.html", 96, 0.18465991},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIII_30.5.html", 96, 0.18465991},
				Result{"http://localhost:8080/top10", 1, 0.14219016},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneI_30.4.html", 62, 0.13993317},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneI_30.2.html", 62, 0.13993317},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneI_30.3.html", 62, 0.13993317},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneI_30.1.html", 62, 0.13993317},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneI_30.5.html", 62, 0.13993317},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneV_30.3.html", 39, 0.13983132},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneV_30.1.html", 39, 0.13983132},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneV_30.2.html", 39, 0.13983132},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneI_30.0.html", 1, 0.084425405},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIV_30.4.html", 12, 0.045373484},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIV_30.1.html", 12, 0.045373484},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIV_30.2.html", 12, 0.045373484},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIV_30.5.html", 12, 0.045373484},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIV_30.3.html", 12, 0.045373484},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/index.html", 10, 0.035246093},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20The%20Picture%20of%20Dorian%20Gray,%20by%20Oscar%20Wilde/chap04.html", 4, 0.01252196},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20The%20Picture%20of%20Dorian%20Gray,%20by%20Oscar%20Wilde/chap07.html", 3, 0.011087331},
			},
		},
	}

	for testIdx, test := range tests {
		// Make sure we got the expected number of results
		got := test.ii.Search("Romeo")
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
