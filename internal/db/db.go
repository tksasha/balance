package db

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Open(ctx context.Context, dbNameProvider DBNameProvider) *sql.DB {
	db, err := sql.Open("sqlite3", dbNameProvider.Provide())
	if err != nil {
		log.Fatal(err)
	}

	if err := db.PingContext(ctx); err != nil {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}

		log.Fatal(err)
	}

	if _, err := db.ExecContext(ctx, "PRAGMA foreign_keys=true"); err != nil {
		log.Fatal(err)
	}

	if err := newMigration(db).run(ctx); err != nil {
		log.Fatal(err)
	}

	return db
}
