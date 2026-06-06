package crawler

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"github.com/nishoof/search-engine/testutils"
)

func TestDownload(t *testing.T) {
	tests := []struct {
		path string
	}{
		{testutils.SimpleTestdataPaths[1]},
		{testutils.SimpleTestdataPaths[0]},
		{testutils.SimpleTestdataPaths[2]},
	}

	ts := testutils.NewSimpleTestServer()
	defer ts.Close()

	for testIdx, test := range tests {
		want, err := loadFileAsLines(test.path)
		if err != nil {
			t.Fatalf("Unable to load test data for test %d: %v", testIdx, err)
		}

		got := download(ts.URL + "/" + test.path)

		// Read through the downloaded content line by line with a scanner and compare it to expected
		scanner := bufio.NewScanner(got)
		for i := 0; scanner.Scan(); i++ {
			textGot := strings.TrimSpace(scanner.Text()) // ignore indentation differences
			textWant := want[i]
			if textGot != textWant {
				t.Errorf("For test %d at line %d, got \"%s\" but wanted \"%s\"\n", testIdx, i, textGot, textWant)
			}
		}
	}
}

func loadFileAsLines(path string) ([]string, error) {
	f, err := os.Open(testutils.Root() + "/" + path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}
	return lines, scanner.Err()
}
