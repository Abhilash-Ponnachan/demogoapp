package main

import "database/sql"

// repository interface for quotes
type quoteRepo interface {
	listQuotes(limit uint, qtsBuff *[]quote)
}

// dummy implementation for quote repo
type quoteRepoDummy struct{}

func (qr quoteRepoDummy) listQuotes(limit uint, qtsBuff *[]quote) {
	if qtsBuff != nil {
		var q quote
		q = quote{
			Author: "John Lennon",
			Quote:  "Life is what happens when you’re busy making other plans.",
		}
		*qtsBuff = append(*qtsBuff, q)
		q = quote{
			Author: "Albert Einstein",
			Quote:  "If you want to live a happy life, tie it to a goal, not to people or things.",
		}
		*qtsBuff = append(*qtsBuff, q)
		q = quote{
			Author: "Babe Ruth",
			Quote:  "Never let the fear of striking out keep you from playing the game.",
		}
		*qtsBuff = append(*qtsBuff, q)
		q = quote{
			Author: "Steve Jobs",
			Quote:  "Your time is limited, so don’t waste it living someone else’s life. Don’t be trapped by dogma – which is living with the results of other people’s thinking.",
		}
		*qtsBuff = append(*qtsBuff, q)
		q = quote{
			Author: "Soren Kierkegaard",
			Quote:  "Life is not a problem to be solved, but a reality to be experienced.",
		}
		*qtsBuff = append(*qtsBuff, q)
		q = quote{
			Author: "Charles Swindoll",
			Quote:  "Life is ten percent what happens to you and ninety percent how you respond to it.",
		}
		*qtsBuff = append(*qtsBuff, q)
	}
}

// sqllite implementation for quote repo
type quoteRepoSQL struct{}

func (qr quoteRepoSQL) listQuotes(limit uint, qtsBuff *[]quote) {
	sqlDb, err := sql.Open("sqlite3", config().DbConnection)
	if err != nil {
		panic(err.Error())
	}
	defer sqlDb.Close()

}

/*
 Install Sqlite driver for Go
	$ go get github.com/mattn/go-sqlite3

	>> download and install package in GOPATH(~/go)/pkg/mod/github.com/mattn/go-sqlite3@v1.14.8
	   ... approximately 9MB !!, has 'C' binding
	>> update mod.go file with 'require'
*/
