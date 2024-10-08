package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tksasha/balance/internal/server/db/migrations"
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

	migrate(context.Background(), db)

	return db, nil
}

func migrate(ctx context.Context, db *sql.DB) {
	migrations.
		NewAddItemsCategoryNameMigration(db).Up(ctx)
}
