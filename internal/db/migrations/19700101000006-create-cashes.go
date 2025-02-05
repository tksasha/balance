package migrations

import (
	"context"
	"database/sql"
	"log"
)

type CreateCashesMigration struct {
	db              *sql.DB
	schemaMigration *SchemaMigration
}

func NewCreateCashesMigration(db *sql.DB) *CreateCashesMigration {
	return &CreateCashesMigration{
		db:              db,
		schemaMigration: NewSchemaMigration(db, "19700101000006"),
	}
}

func (m *CreateCashesMigration) Up(ctx context.Context) {
	if m.schemaMigration.IsExist(ctx) {
		return
	}

	query := `
		BEGIN TRANSACTION;

		CREATE TABLE IF NOT EXISTS "cashes"(
			"id" integer NOT NULL PRIMARY KEY,
			"sum" decimal(10,2) DEFAULT NULL,
			"name" varchar DEFAULT NULL,
			"deleted_at" time DEFAULT NULL,
			"formula" varchar DEFAULT NULL,
			"currency" integer DEFAULT 0,
			"supercategory" integer DEFAULT 1 NOT NULL,
			"favorite" boolean DEFAULT 0
		);

		CREATE UNIQUE INDEX IF NOT EXISTS "index_cashes_on_name_and_currency" ON "cashes"("name", "currency");

		COMMIT;
		`

	if _, err := m.db.ExecContext(ctx, query); err != nil {
		log.Fatalf("create table cashes error: %v", err)
	}

	m.schemaMigration.Save(ctx)
}
