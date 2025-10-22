package main

type IndexType bool

const IN_MEM IndexType = false
const SQLITE IndexType = true

type Index interface {
	GetDocs() []string
	GetFrequency(word, documentName string) int
	GetNumDocs() int
	GetNumDocsWithWord(word string) int
	GetWordCount(documentName string) int
	Increment(word, documentName string, count int)
}
