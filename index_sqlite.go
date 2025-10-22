package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type IndexSQLite struct {
	db                 *sql.DB
	preparedStatements PreparedStatements
	pendingFrequencies []Frequency
}

type PreparedStatements struct {
	getDocs            *sql.Stmt
	getFreqCount       *sql.Stmt
	getFreqIdCount     *sql.Stmt
	getNumDocs         *sql.Stmt
	getNumDocsWithWord *sql.Stmt
	getWordCount       *sql.Stmt
	addWord            *sql.Stmt
	addDoc             *sql.Stmt
	addFreq            *sql.Stmt
	updateWordCount    *sql.Stmt
	getWordId          *sql.Stmt
	getDocumentId      *sql.Stmt
}

type Frequency struct {
	wordId, documentId, count int
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func NewIndexSQLite() IndexSQLite {
	db, err := sql.Open("sqlite3", "db.db")
	checkErr(err)

	initTables(db)

	var preparedStatements PreparedStatements
	preparedStatements.getDocs, err = db.Prepare(`SELECT name FROM documents;`)
	checkErr(err)
	preparedStatements.getFreqCount, err = db.Prepare(`SELECT count FROM frequencies WHERE word_id = ? AND doc_id = ?;`)
	checkErr(err)
	preparedStatements.getFreqIdCount, err = db.Prepare(`SELECT id, count FROM frequencies WHERE word_id = ? AND doc_id = ?;`)
	checkErr(err)
	preparedStatements.getNumDocs, err = db.Prepare(`SELECT COUNT(*) FROM documents;`)
	checkErr(err)
	preparedStatements.getNumDocsWithWord, err = db.Prepare(`SELECT COUNT(*) FROM frequencies WHERE word_id = ?;`)
	checkErr(err)
	preparedStatements.getWordCount, err = db.Prepare(`SELECT word_count FROM documents WHERE id = ?;`)
	checkErr(err)
	preparedStatements.addWord, err = db.Prepare(`INSERT INTO words(word) VALUES(?);`)
	checkErr(err)
	preparedStatements.addDoc, err = db.Prepare(`INSERT INTO documents(name, word_count) VALUES(?, 0);`)
	checkErr(err)
	preparedStatements.addFreq, err = db.Prepare(`INSERT INTO frequencies(word_id, doc_id, count) VALUES(?, ?, ?);`)
	checkErr(err)
	preparedStatements.updateWordCount, err = db.Prepare(`UPDATE documents SET word_count = ? WHERE id = ?;`)
	checkErr(err)
	preparedStatements.getWordId, err = db.Prepare(`SELECT id FROM words WHERE word = ?;`)
	checkErr(err)
	preparedStatements.getDocumentId, err = db.Prepare(`SELECT id FROM documents WHERE name = ?;`)
	checkErr(err)

	idx := IndexSQLite{db, preparedStatements, make([]Frequency, 1000)}

	return idx
}

func (idx IndexSQLite) GetDocs() []string {
	stmt := idx.preparedStatements.getDocs
	rows, err := stmt.Query()
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
	wordId := getWordId(idx.preparedStatements, word)
	if wordId == -1 {
		return 0 // word doesn't occur in any document
	}

	documentId := getDocumentId(idx.preparedStatements, documentName)
	if documentId == -1 {
		return 0 // document doesn't exist
	}

	stmt := idx.preparedStatements.getFreqCount
	rows, err := stmt.Query(wordId, documentId)
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
	stmt := idx.preparedStatements.getNumDocs
	row := stmt.QueryRow()

	var count int
	err := row.Scan(&count)
	checkErr(err)
	return count
}

func (idx IndexSQLite) GetNumDocsWithWord(word string) int {
	wordId := getWordId(idx.preparedStatements, word)
	if wordId == -1 {
		return 0 // the word doesn't occur in any documents
	}

	stmt := idx.preparedStatements.getNumDocsWithWord
	row := stmt.QueryRow(wordId)

	var numDocs int
	err := row.Scan(&numDocs)
	checkErr(err)
	return numDocs
}

func (idx IndexSQLite) GetWordCount(documentName string) int {
	documentId := getDocumentId(idx.preparedStatements, documentName)
	if documentId == -1 {
		return 0
	}

	stmt := idx.preparedStatements.getWordCount
	row := stmt.QueryRow(documentId)

	var wordCount int
	err := row.Scan(&wordCount)
	checkErr(err)
	return wordCount
}

func (idx IndexSQLite) Increment(word, documentName string, count int) {
	// get the wordId
	wordId := getWordId(idx.preparedStatements, word)
	if wordId == -1 {
		// word doesn't occur in any document, so add it to the words table
		stmt := idx.preparedStatements.addWord
		_, err := stmt.Exec(word)
		checkErr(err)
		wordId = getWordId(idx.preparedStatements, word)
	}

	// get the documentId
	documentId := getDocumentId(idx.preparedStatements, documentName)
	if documentId == -1 {
		// document doesn't exist, so add it to the documents table
		stmt := idx.preparedStatements.addDoc
		_, err := stmt.Exec(documentName)
		checkErr(err)
		documentId = getDocumentId(idx.preparedStatements, documentName)
	}

	// update frequency
	idx.pendingFrequencies = append(idx.pendingFrequencies, Frequency{wordId, documentId, count})
	if len(idx.pendingFrequencies) == 1000 {
		idx.Flush()
	}

	// update word count
	wordCount := idx.GetWordCount(documentName)
	stmt := idx.preparedStatements.updateWordCount
	_, err := stmt.Exec(wordCount+count, documentId)
	checkErr(err)
}

func (idx IndexSQLite) Flush() {
	tx, err := idx.db.Begin()
	checkErr(err)

	for _, pf := range idx.pendingFrequencies {
		stmt := idx.preparedStatements.addFreq
		tx.Stmt(stmt).Exec(pf.wordId, pf.documentId, pf.count)
	}

	tx.Commit()
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
			count INTEGER,
			FOREIGN KEY (word_id) REFERENCES words(id),
			FOREIGN KEY (doc_id) REFERENCES documents(id)
		);
	`)
	checkErr(err)

	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS idx_documents_name ON documents(name);`)
	checkErr(err)

	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS idx_words_word ON words(word);`)
	checkErr(err)

	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS idx_frequencies_word_doc ON frequencies(word_id, doc_id);`)
	checkErr(err)

	_, err = db.Exec(`PRAGMA foreign_keys = ON;`)
	checkErr(err)
}

func getWordId(preparedStatements PreparedStatements, word string) int {
	stmt := preparedStatements.getWordId
	rows, err := stmt.Query(word)
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		var id int
		rows.Scan(&id)
		return id
	}

	return -1
}

func getDocumentId(preparedStatements PreparedStatements, documentName string) int {
	stmt := preparedStatements.getDocumentId
	rows, err := stmt.Query(documentName)
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		var id int
		rows.Scan(&id)
		return id
	}

	return -1
}

func getFrequency(preparedStatements PreparedStatements, wordId, documentId int) (int, int) {
	stmt := preparedStatements.getFreqIdCount
	rows, err := stmt.Query(wordId, documentId)
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		var id, count int
		rows.Scan(&id, &count)
		return id, count
	}

	return -1, -1
}
