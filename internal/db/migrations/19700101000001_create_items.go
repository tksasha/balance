package migrations

import (
	"context"
	"database/sql"
	"log"
)

type CreateItemsMigration struct {
	db              *sql.DB
	schemaMigration *SchemaMigration
}

func NewCreateItemsMigration(db *sql.DB) *CreateItemsMigration {
	return &CreateItemsMigration{
		db:              db,
		schemaMigration: NewSchemaMigration(db, "19700101000001"),
	}
}

func (m *CreateItemsMigration) Up(ctx context.Context) {
	if m.schemaMigration.IsExist(ctx) {
		return
	}

	query := `
		CREATE TABLE IF NOT EXISTS
			"items" (
				"id" integer NOT NULL PRIMARY KEY,
				"date" date DEFAULT NULL,
				"sum" decimal(10,2) NOT NULL,
				"category_id" integer DEFAULT NULL,
				"description" varchar DEFAULT NULL,
				"created_at" datetime DEFAULT NULL,
				"updated_at" datetime DEFAULT NULL,
				"formula" text DEFAULT NULL,
				"deleted_at" time DEFAULT NULL,
				"currency" integer DEFAULT 0,
				CONSTRAINT "fk_rails_89fb86dc8b" FOREIGN KEY ("category_id") REFERENCES "categories" ("id")
			)
	`

	if _, err := m.db.ExecContext(ctx, query); err != nil {
		log.Fatalf("create table items error: %v", err)
	}

	m.schemaMigration.Save(ctx)
}
