package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		idx  Index
		word string
		want Results
	}{
		{
			startServer(SQLITE),
			"Romeo",
			Results{
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneII_30.2.html", 41, 0.671309888},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneII_30.4.html", 41, 0.664264619},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneII_30.1.html", 9, 0.589442790},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIII_30.2.html", 26, 0.522244871},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneV_30.2.html", 5, 0.469845712},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIII_30.1.html", 30, 0.456610620},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneI_30.2.html", 14, 0.455693752},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIV_30.3.html", 8, 0.450268805},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIII_30.3.html", 27, 0.450268805},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneI_30.4.html", 16, 0.418651879},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneII_30.6.html", 5, 0.388721287},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneII_30.0.html", 2, 0.360215068},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneV_30.1.html", 10, 0.356648564},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneII_30.3.html", 12, 0.344886750},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneI_30.5.html", 14, 0.313880324},
				Result{"http://localhost:8080/top10", 1, 0.284380317},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneI_30.1.html", 21, 0.277088493},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneV_30.3.html", 26, 0.254730493},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIII_30.5.html", 16, 0.211373135},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneI_30.0.html", 1, 0.168850809},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIV_30.1.html", 5, 0.125656411},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneII_30.5.html", 3, 0.123267509},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIII_30.4.html", 1, 0.087857328},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIV_30.4.html", 1, 0.078307621},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/index.html", 10, 0.070492186},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIV_30.2.html", 1, 0.065493643},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneI_30.3.html", 1, 0.028740562},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20The%20Picture%20of%20Dorian%20Gray,%20by%20Oscar%20Wilde/chap04.html", 4, 0.025043920},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20The%20Picture%20of%20Dorian%20Gray,%20by%20Oscar%20Wilde/chap07.html", 3, 0.022174662},
				Result{"http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIV_30.5.html", 1, 0.020312879},
			},
		},
	}

	for testIdx, test := range tests {
		// Make sure we got the expected number of results
		got := Search(test.word, test.idx)
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
