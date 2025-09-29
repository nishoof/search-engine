package main

type Result struct {
	url         string
	occurrences int
	score       float32
}

type Results []Result

type IndexType bool

const IN_MEM IndexType = false
const SQLITE IndexType = true

type Index interface {
	GetFrequency(word, documentName string) int
	GetNumDocs() int
	GetNumDocsWithWord(word string) int
	GetWordCount(documentName string) int
	Increment(word, documentName string)
	Search(word string) Results
}
