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
			startServer(),
			"Romeo",
			Results{
				{url: "http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneII_30.2.html", occurrences: 41, score: 0.33565494},
				{url: "http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneII_30.4.html", occurrences: 41, score: 0.33213231},
				{url: "http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneII_30.1.html", occurrences: 9, score: 0.29472139},
				{url: "http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIII_30.2.html", occurrences: 26, score: 0.26112244},
				{url: "http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneV_30.2.html", occurrences: 5, score: 0.23492286},
				{url: "http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIII_30.1.html", occurrences: 30, score: 0.22830531},
				{url: "http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneI_30.2.html", occurrences: 14, score: 0.22784688},
				{url: "http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIV_30.3.html", occurrences: 8, score: 0.22513440},
				{url: "http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIII_30.3.html", occurrences: 27, score: 0.22513440},
				{url: "http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneI_30.4.html", occurrences: 16, score: 0.20932594},
				{url: "http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneII_30.6.html", occurrences: 5, score: 0.19436064},
				{url: "http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneII_30.0.html", occurrences: 2, score: 0.18010753},
				{url: "http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneV_30.1.html", occurrences: 10, score: 0.17832428},
				{url: "http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneII_30.3.html", occurrences: 12, score: 0.17244337},
				{url: "http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneI_30.5.html", occurrences: 14, score: 0.15694016},
				{url: "http://localhost:8080/top10", occurrences: 1, score: 0.14219016},
				{url: "http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneI_30.1.html", occurrences: 21, score: 0.13854425},
				{url: "http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneV_30.3.html", occurrences: 26, score: 0.12736525},
				{url: "http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIII_30.5.html", occurrences: 16, score: 0.10568657},
				{url: "http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneI_30.0.html", occurrences: 1, score: 0.084425405},
				{url: "http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIV_30.1.html", occurrences: 5, score: 0.062828206},
				{url: "http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneII_30.5.html", occurrences: 3, score: 0.061633755},
				{url: "http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIII_30.4.html", occurrences: 1, score: 0.043928664},
				{url: "http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIV_30.4.html", occurrences: 1, score: 0.03915381},
				{url: "http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/index.html", occurrences: 10, score: 0.035246093},
				{url: "http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIV_30.2.html", occurrences: 1, score: 0.03274682},
				{url: "http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneI_30.3.html", occurrences: 1, score: 0.014370281},
				{url: "http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20The%20Picture%20of%20Dorian%20Gray,%20by%20Oscar%20Wilde/chap04.html", occurrences: 4, score: 0.01252196},
				{url: "http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20The%20Picture%20of%20Dorian%20Gray,%20by%20Oscar%20Wilde/chap07.html", occurrences: 3, score: 0.011087331},
				{url: "http://localhost:8080/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare/sceneIV_30.5.html", occurrences: 1, score: 0.01015644},
			},
		},
	}

	for testIdx, test := range tests {
		// Make sure we got the expected number of results
		got := test.idx.Search(test.word)
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
