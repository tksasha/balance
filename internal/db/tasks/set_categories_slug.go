package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tksasha/balance/internal/db/nameprovider"
	"github.com/tksasha/xstrings"
)

func setCategoriesSlug() {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", nameprovider.New().Provide())
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.QueryContext(ctx, `SELECT id, name FROM categories WHERE slug IS NULL`)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var id int

		var name string

		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}

		query := `UPDATE categories SET slug = ? WHERE id = ?`

		result, err := db.ExecContext(ctx, query, xstrings.Transliterate(name), id)
		if err != nil {
			log.Fatal(err)
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}

		if rowsAffected == 0 {
			log.Fatalf("failed to update category id: %d", id)
		}
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
