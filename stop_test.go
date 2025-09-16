package main

import (
	"testing"
)

func TestStop(t *testing.T) {
	tests := []struct {
		word string
		want bool
	}{
		{"the", true},
		{"and", true},
		{"is", true},
		{"hello", true},
		{"class", false},
		{"nishil", false},
		{"phil", false},
	}

	for testIdx, test := range tests {
		got := isStopWord(test.word)
		if got != test.want {
			t.Errorf("test %d (%s): got %v, want %v", testIdx, test.word, got, test.want)
		}
	}
}
