package main

type IndexInMemory struct {
	frequency      map[string]map[string]int // maps words to their FrequencyMap
	wordCount      map[string]int            // maps document names to their word counts
	documentTitles map[string]string         // maps document names to their titles
}

func NewIndexInMemory() IndexInMemory {
	idx := new(IndexInMemory)
	idx.frequency = make(map[string]map[string]int)
	idx.wordCount = make(map[string]int)
	idx.documentTitles = make(map[string]string)
	return *idx
}

func (idx IndexInMemory) AddDoc(documentName, title string) {
	if title == "" {
		title = documentName
	}
	_, exists := idx.wordCount[documentName]
	if !exists {
		idx.wordCount[documentName] = 0
	}
	idx.documentTitles[documentName] = title
}

func (idx IndexInMemory) GetDocs() []string {
	// the document names are the keys of the wordCount map
	documentNames := make([]string, len(idx.wordCount))
	i := 0
	for documentName := range idx.wordCount {
		documentNames[i] = documentName
		i++
	}
	return documentNames
}

func (idx IndexInMemory) GetFrequency(word, documentName string) int {
	fm, exists := idx.frequency[word]
	if !exists {
		return 0
	}
	return fm[documentName]
}

func (idx IndexInMemory) GetNumDocs() int {
	return len(idx.wordCount)
}

func (idx IndexInMemory) GetNumDocsWithWord(word string) int {
	fm, exists := idx.frequency[word]
	if !exists {
		return 0
	}
	return len(fm)
}

func (idx IndexInMemory) GetTitle(documentName string) string {
	return idx.documentTitles[documentName]
}

func (idx IndexInMemory) GetWordCount(documentName string) int {
	return idx.wordCount[documentName]
}

func (idx IndexInMemory) Increment(word, documentName string, count int) {
	_, exists := idx.frequency[word]
	if !exists {
		idx.frequency[word] = make(map[string]int)
	}
	idx.frequency[word][documentName] += count
	idx.wordCount[documentName] += count
}

func (idx IndexInMemory) Flush() {}
