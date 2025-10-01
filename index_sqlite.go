package main

import (
	"database/sql"
	"fmt"

	"github.com/kljensen/snowball"
	_ "github.com/mattn/go-sqlite3"
)

type IndexSQLite struct {
	db *sql.DB
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func NewIndexSQLite(mp map[string][]string) IndexSQLite {
	db, err := sql.Open("sqlite3", "db.db")
	checkErr(err)

	initTables(db)
	idx := IndexSQLite{db}

	for url, words := range mp {
		for _, w := range words {
			stemmed, err := snowball.Stem(w, "english", true)
			if err != nil {
				panic(err)
			}
			idx.Increment(stemmed, url)
		}
	}

	return idx
}

func (idx IndexSQLite) GetDocs() []string {
	db := idx.db

	rows, err := db.Query(`
		SELECT name FROM documents;
	`)
	checkErr(err)
	defer rows.Close()

	documentNames := make([]string, 0)
	for rows.Next() {
		var documentName string
		rows.Scan(&documentName)
		documentNames = append(documentNames, documentName)
	}
	return documentNames
}

func (idx IndexSQLite) GetFrequency(word, documentName string) int {
	db := idx.db

	wordId := getWordId(db, word)
	if wordId == -1 {
		return 0 // word doesn't occur in any document
	}

	documentId := getDocumentId(db, documentName)
	if documentId == -1 {
		return 0 // document doesn't exist
	}

	rows, err := db.Query(`
		SELECT count FROM frequencies WHERE word_id = ? AND doc_id = ?
	`, wordId, documentId)
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		var count int
		err = rows.Scan(&count)
		checkErr(err)
		return count
	}

	return 0
}

func (idx IndexSQLite) GetNumDocs() int {
	db := idx.db

	row := db.QueryRow(`
		SELECT COUNT(*) FROM documents
	`)

	var count int
	err := row.Scan(&count)
	checkErr(err)
	return count
}

func (idx IndexSQLite) GetNumDocsWithWord(word string) int {
	// fm, exists := idx.frequency[word]
	// if !exists {
	// 	return 0
	// }
	// return len(fm)

	db := idx.db

	wordId := getWordId(db, word)
	if wordId == -1 {
		return 0 // the word doesn't occur in any documents
	}

	row := db.QueryRow(`
		SELECT COUNT(*) FROM frequencies WHERE word_id = ?
	`, wordId)

	var numDocs int
	err := row.Scan(&numDocs)
	checkErr(err)
	return numDocs
}

func (idx IndexSQLite) GetWordCount(documentName string) int {
	db := idx.db

	documentId := getDocumentId(db, documentName)
	if documentId == -1 {
		return 0
	}

	row := db.QueryRow(`
		SELECT word_count FROM documents WHERE id = ?;
	`, documentId)

	var wordCount int
	err := row.Scan(&wordCount)
	checkErr(err)
	return wordCount
}

func (idx IndexSQLite) Increment(word, documentName string) {
	db := idx.db

	// get the wordId
	wordId := getWordId(db, word)
	if wordId == -1 {
		// word doesn't occur in any document, so add it to the words table
		_, err := db.Exec(`
			INSERT INTO words(word) VALUES(?);
		`, word)
		checkErr(err)
		wordId = getWordId(db, word)
	}

	// get the documentId
	documentId := getDocumentId(db, documentName)
	if documentId == -1 {
		// document doesn't exist, so add it to the documents table
		fmt.Printf("New document %s\n", documentName)
		_, err := db.Exec(`
			INSERT INTO documents(name, word_count) VALUES(?, 0)
		`, documentName)
		checkErr(err)
		documentId = getDocumentId(db, documentName)
	}

	// update frequency
	frequencyId, frequencyCount := getFrequency(db, wordId, documentId)
	if frequencyId == -1 {
		// frequency doesn't exist (first occurrence of this word in this document), so add it to the frequencies table
		_, err := db.Exec(`
			INSERT INTO frequencies(word_id, doc_id, count) VALUES(?, ?, 1)
		`, wordId, documentId)
		checkErr(err)
	} else {
		// frequency already exists, we need to update it now
		_, err := db.Exec(`
			UPDATE frequencies SET count = ? WHERE id = ?
		`, frequencyCount+1, frequencyId)
		checkErr(err)
	}

	// update word count
	wordCount := idx.GetWordCount(documentName)
	_, err := db.Exec(`
		UPDATE documents SET word_count = ? WHERE id = ?
	`, wordCount+1, documentId)
	checkErr(err)
}

func (idx IndexSQLite) Close() {
	idx.db.Close()
}

func initTables(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS documents (
			id INTEGER PRIMARY KEY,
			name TEXT,
			word_count INTEGER
		);
	`)
	checkErr(err)

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS words (
			id INTEGER PRIMARY KEY,
			word TEXT
		);
	`)
	checkErr(err)

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS frequencies (
			id INTEGER PRIMARY KEY,
			word_id INTEGER,
			doc_id INTEGER,
			count INTEGER
		);
	`)
	checkErr(err)
}

func getWordId(db *sql.DB, word string) int {
	rows, err := db.Query(`
		SELECT id FROM words WHERE word = ?;
	`, word)
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		var id int
		rows.Scan(&id)
		return id
	}

	return -1
}

func getDocumentId(db *sql.DB, documentName string) int {
	rows, err := db.Query(`
		SELECT id FROM documents WHERE name = ?;
	`, documentName)
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		var id int
		rows.Scan(&id)
		return id
	}

	return -1
}

func getFrequency(db *sql.DB, wordId, documentId int) (int, int) {
	rows, err := db.Query(`
		SELECT id, count FROM frequencies WHERE word_id = ? AND doc_id = ?;
	`, wordId, documentId)
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		var id, count int
		rows.Scan(&id, &count)
		return id, count
	}

	return -1, -1
}
