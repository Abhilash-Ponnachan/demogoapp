package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// repository interface for quotes
type quoteRepo interface {
	init()
	close()
	listQuotes(limit uint, qtsBuff *[]quote)
	addQuote(qtObj *quote) uint
	deleteQuote(rowId uint) uint
}

// sqllite implementation for quote repo
type quoteRepoSQL struct {
	db *sql.DB
}

func (qr *quoteRepoSQL) init() {
	db, err := sql.Open("sqlite3", config().DbConnection)
	checkerr(err)
	qr.db = db
	qr.prepareDB()
}

func (qr *quoteRepoSQL) close() {
	if qr.db != nil {
		qr.db.Close()
	}
}

func (qr *quoteRepoSQL) prepareDB() {
	cmd := `CREATE TABLE IF NOT EXISTS "quotes" (
		"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		"text"	TEXT,
		"author"	TEXT
	);`
	stmt, err := qr.db.Prepare(cmd)
	checkerr(err)
	_, err = stmt.Exec()
	checkerr(err)
}

func (qr *quoteRepoSQL) listQuotes(limit uint, qtsBuff *[]quote) {
	if qtsBuff != nil {
		cmd := `SELECT id, text, author FROM quotes ORDER BY id;`
		rows, err := qr.db.Query(cmd)
		checkerr(err)
		defer rows.Close()
		// iterate through rows cursor and append to buffer
		for rows.Next() {
			var id uint
			var author string
			var text string
			rows.Scan(&id, &text, &author)
			q := quote{id, author, text}
			*qtsBuff = append(*qtsBuff, q)
		}
	}
}

func (qr *quoteRepoSQL) addQuote(qtObj *quote) uint {
	var lastId int64 = 0
	if qtObj != nil {
		//log.Printf("Quote Obj = %v\n", qtObj)
		cmd := `INSERT INTO quotes (author, text) VALUES(?, ?);`
		stmt, err := qr.db.Prepare(cmd)
		checkerr(err)
		rslt, err := stmt.Exec(qtObj.Author, qtObj.Quote)
		checkerr(err)
		lastId, err = rslt.LastInsertId()
		checkerr(err)
	}
	return uint(lastId)
}

func (qr *quoteRepoSQL) deleteQuote(rowId uint) uint {
	var rowsDeleted int64 = 0
	cmd := `DELETE FROM quotes WHERE id = ?;`
	stmt, err := qr.db.Prepare(cmd)
	checkerr(err)
	rslt, err := stmt.Exec(rowId)
	checkerr(err)
	rowsDeleted, err = rslt.RowsAffected()
	checkerr(err)
	return uint(rowsDeleted)
}

/*
 Install Sqlite driver for Go
	$ go get github.com/mattn/go-sqlite3

	>> download and install package in GOPATH(~/go)/pkg/mod/github.com/mattn/go-sqlite3@v1.14.8
	   ... approximately 9MB !!, has 'C' binding
	>> update mod.go file with 'require'
*/
