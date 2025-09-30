package main

import (
	"testing"
)

func TestSQL(t *testing.T) {
	idx := NewIndexSQLite(map[string][]string{"abc": []string{"a", "b", "a", "a"}, "def": []string{"d", "e", "f", "a"}})

	if idx.GetFrequency("a", "abc") != 3 {
		t.Error()
	}
	if idx.GetNumDocs() != 2 {
		t.Error()
	}
	if idx.GetNumDocsWithWord("a") != 2 {
		t.Error()
	}
	if idx.GetNumDocsWithWord("b") != 1 {
		t.Error()
	}
	if idx.GetNumDocsWithWord("x") != 0 {
		t.Error()
	}
	if idx.GetWordCount("abc") != 4 {
		t.Error()
	}

	idx.Close()
}
