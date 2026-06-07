package searcher

import (
	"testing"

	"github.com/nishoof/search-engine/crawler"
	"github.com/nishoof/search-engine/index"
	"github.com/nishoof/search-engine/testutils"
)

func TestSearch(t *testing.T) {
	top10TestServer := testutils.NewTop10TestServer()
	defer top10TestServer.Close()

	var idx index.Index = index.NewIndexInMemory()
	crawler.Crawl(top10TestServer.URL+"/top10", true, &idx)

	const romeoTitle = "The Project Gutenberg eBook of Romeo and Juliet, by William Shakespeare"
	const dorianTitle = "The Project Gutenberg eBook of The Picture of Dorian Gray, by Oscar Wilde"
	base := top10TestServer.URL
	romeo := base + "/top10/The%20Project%20Gutenberg%20eBook%20of%20Romeo%20and%20Juliet,%20by%20William%20Shakespeare"
	dorian := base + "/top10/The%20Project%20Gutenberg%20eBook%20of%20The%20Picture%20of%20Dorian%20Gray,%20by%20Oscar%20Wilde"

	tests := []struct {
		word string
		want Results
	}{
		{
			"Romeo",
			Results{
				{base + "/top10", 1, 0.26570457, base + "/top10"},
				{dorian + "/chap04.html", 4, 0.023399245, dorianTitle},
				{dorian + "/chap07.html", 3, 0.020718414, dorianTitle},
				{romeo + "/sceneII_30.2.html", 41, 0.62722385, romeoTitle},
				{romeo + "/sceneII_30.4.html", 41, 0.6206413, romeoTitle},
				{romeo + "/sceneII_30.1.html", 9, 0.55073315, romeoTitle},
				{romeo + "/sceneIII_30.2.html", 26, 0.4879482, romeoTitle},
				{romeo + "/sceneV_30.2.html", 5, 0.43899018, romeoTitle},
				{romeo + "/sceneIII_30.1.html", 30, 0.42662427, romeoTitle},
				{romeo + "/sceneI_30.2.html", 14, 0.4257676, romeoTitle},
				{romeo + "/sceneIII_30.3.html", 27, 0.4206989, romeoTitle},
				{romeo + "/sceneIV_30.3.html", 8, 0.4206989, romeoTitle},
				{romeo + "/sceneI_30.4.html", 16, 0.3911583, romeoTitle},
				{romeo + "/sceneII_30.6.html", 5, 0.3631933, romeoTitle},
				{romeo + "/sceneII_30.0.html", 2, 0.33655915, romeoTitle},
				{romeo + "/sceneV_30.1.html", 10, 0.33322686, romeoTitle},
				{romeo + "/sceneII_30.3.html", 12, 0.32223746, romeoTitle},
				{romeo + "/sceneI_30.5.html", 14, 0.2932673, romeoTitle},
				{romeo + "/sceneI_30.1.html", 21, 0.25889164, romeoTitle},
				{romeo + "/sceneV_30.3.html", 26, 0.23800193, romeoTitle},
				{romeo + "/sceneIII_30.5.html", 16, 0.19749191, romeoTitle},
				{romeo + "/sceneI_30.0.html", 1, 0.1577621, romeoTitle},
				{romeo + "/sceneIV_30.1.html", 5, 0.11740435, romeoTitle},
				{romeo + "/sceneII_30.5.html", 3, 0.11517233, romeoTitle},
				{romeo + "/sceneIII_30.4.html", 1, 0.0820876, romeoTitle},
				{romeo + "/sceneIV_30.4.html", 1, 0.07316503, romeoTitle},
				{romeo + "/index.html", 10, 0.06586284, romeoTitle},
				{romeo + "/sceneIV_30.2.html", 1, 0.061192572, romeoTitle},
				{romeo + "/sceneI_30.3.html", 1, 0.026853124, romeoTitle},
				{romeo + "/sceneIV_30.5.html", 1, 0.0189789, romeoTitle},
			},
		},
	}

	for testIdx, test := range tests {
		got := Search(test.word, idx)

		// Make sure we got the expected number of results
		if len(got) != len(test.want) {
			t.Errorf("Test %d: Got %d results but wanted %d\n", testIdx, len(got), len(test.want))
		}

		// Convert the slice of wanted Results to a set of wanted Results
		wantSet := make(map[Result]struct{})
		for _, result := range test.want {
			wantSet[result] = struct{}{}
		}

		// Make sure each result we got is in the set of wanted Results
		for _, result := range got {
			_, exists := wantSet[result]
			if !exists {
				t.Errorf("Test %d: got unexpected result %v\n", testIdx, result)
			}
		}
	}
}
