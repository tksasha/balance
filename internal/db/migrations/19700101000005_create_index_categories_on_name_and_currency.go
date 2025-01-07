package migrations

import (
	"context"
	"database/sql"
	"log"
)

type CreateIndexCategoriesOnNameAndCurrency struct {
	db              *sql.DB
	schemaMigration *SchemaMigration
}

func NewCreateIndexCategoriesOnNameAndCurrency(db *sql.DB) *CreateIndexCategoriesOnNameAndCurrency {
	return &CreateIndexCategoriesOnNameAndCurrency{
		db:              db,
		schemaMigration: NewSchemaMigration(db, "19700101000005"),
	}
}

func (m *CreateIndexCategoriesOnNameAndCurrency) Up(ctx context.Context) {
	if m.schemaMigration.IsExist(ctx) {
		return
	}

	query := `
		CREATE UNIQUE INDEX IF NOT EXISTS "index_categories_on_name_and_currency" ON "categories" ("name", "currency")
	`

	if _, err := m.db.ExecContext(ctx, query); err != nil {
		log.Fatalf("failed to create index index_categories_on_name_and_currency, error: %v", err)
	}

	m.schemaMigration.Save(ctx)
}
