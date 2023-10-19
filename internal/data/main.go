package data

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "modernc.org/sqlite"
)

var dbUrl = "file:/mnt/c/Users/msh/DataGripProjects/Qim/equal-cannonball-sephix.turso.io?cache=shared&mode=rwc&_journal_mode=WAL"

func GetDb() *sql.DB {

	db, err := sql.Open("sqlite", dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", dbUrl, err)
		os.Exit(1)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	return db
}
