package data

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var dbUrl = "file:/mnt/c/Users/msh/DataGripProjects/Qim/equal-cannonball-sephix.turso.io"

func GetDb() *sql.DB {

	db, err := sql.Open("sqlite3", dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", dbUrl, err)
		os.Exit(1)
	}
	return db
}
