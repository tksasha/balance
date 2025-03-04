package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tksasha/balance/internal/db/nameprovider"
	"github.com/tksasha/xstrings"
)

func setCategoriesSlug() error {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", nameprovider.New().Provide())
	if err != nil {
		return err
	}

	rows, err := db.QueryContext(ctx, `SELECT id, name FROM categories WHERE slug IS NULL`)
	if err != nil {
		return err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("failed to close rows: %v", err)
		}
	}()

	for rows.Next() {
		var id int

		var name string

		if err := rows.Scan(&id, &name); err != nil {
			return err
		}

		query := `UPDATE categories SET slug = ? WHERE id = ?`

		result, err := db.ExecContext(ctx, query, xstrings.Transliterate(name), id)
		if err != nil {
			return err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}

		if rowsAffected == 0 {
			return fmt.Errorf("failed to update category id: %d", id) //nolint:err113
		}
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}
