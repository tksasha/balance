package migrations

import (
	"context"
	"database/sql"
	"log"
)

type ChangeTypeOfDeletedAtInCategories struct {
	db              *sql.DB
	schemaMigration SchemaMigration
}

func NewChangeTypeOfDeletedAtInCategories(db *sql.DB) *ChangeTypeOfDeletedAtInCategories {
	return &ChangeTypeOfDeletedAtInCategories{
		db:              db,
		schemaMigration: *NewSchemaMigration(db, "20250117100105"),
	}
}

func (m *ChangeTypeOfDeletedAtInCategories) Up(ctx context.Context) {
	if m.schemaMigration.IsExist(ctx) {
		return
	}

	query := `
		PRAGMA foreign_keys=OFF;

		BEGIN TRANSACTION;

		CREATE TABLE
			"categories2" (
				"id" integer NOT NULL PRIMARY KEY,
				"name" varchar DEFAULT NULL,
				"income" boolean DEFAULT 0,
				"visible" boolean DEFAULT 1,
				"currency" integer DEFAULT 0,
				"supercategory" integer DEFAULT 1 NOT NULL,
				"deleted_at" datetime
			);

		INSERT INTO categories2
		SELECT id, name, income, visible, currency, supercategory, deleted_at FROM categories;

		DROP TABLE categories;

		ALTER TABLE categories2 RENAME TO categories;

		CREATE UNIQUE INDEX "index_categories_on_name_and_currency" ON "categories" ("name", "currency");

		COMMIT;

		PRAGMA foreign_keys=OFF;
	`

	if _, err := m.db.ExecContext(ctx, query); err != nil {
		log.Fatalf("failed to change type of deleted_at in categories, err: %v", err)
	}

	m.schemaMigration.Save(ctx)
}
