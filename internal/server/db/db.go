package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func Open(workdir, env string) (*sql.DB, error) {
	dbName := fmt.Sprintf(
		"%s%s%s.sqlite3",
		workdir,
		string(os.PathSeparator),
		env,
	)

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec("PRAGMA foreign_keys=true"); err != nil {
		return nil, err
	}

	return db, nil
}
