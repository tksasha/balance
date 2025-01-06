package migrations

import (
	"context"
	"database/sql"
	"log"
)

type CreateCategoriesMigration struct {
	db              *sql.DB
	schemaMigration *SchemaMigration
}

func NewCreateCategoriesMigration(db *sql.DB) *CreateCategoriesMigration {
	return &CreateCategoriesMigration{
		db:              db,
		schemaMigration: NewSchemaMigration(db, "19700101000004"),
	}
}

func (m *CreateCategoriesMigration) Up(ctx context.Context) {
	if m.schemaMigration.IsExist(ctx) {
		return
	}

	query := `
		CREATE TABLE IF NOT EXISTS
			"categories" (
				"id" integer NOT NULL PRIMARY KEY,
				"name" varchar DEFAULT NULL,
				"income" boolean DEFAULT 0,
				"visible" boolean DEFAULT 1,
				"currency" integer DEFAULT 0,
				"supercategory" integer DEFAULT 1 NOT NULL,
				"deleted_at" datetime(6)
			)
	`

	if _, err := m.db.ExecContext(ctx, query); err != nil {
		log.Fatalf("failed to create table categories, error: %v", err)
	}

	m.schemaMigration.Save(ctx)
}
