package migrations

import (
	"context"
	"database/sql"
	"log"
)

type CreateIndexBalanceItemsOnDateAndCategoryID struct {
	db              *sql.DB
	schemaMigration *SchemaMigration
}

func NewCreateIndexBalanceItemsOnDateAndCategoryID(db *sql.DB) *CreateIndexBalanceItemsOnDateAndCategoryID {
	return &CreateIndexBalanceItemsOnDateAndCategoryID{
		db:              db,
		schemaMigration: NewSchemaMigration(db, "19700101000002"),
	}
}

func (m *CreateIndexBalanceItemsOnDateAndCategoryID) Up(ctx context.Context) {
	if m.schemaMigration.IsExist(ctx) {
		return
	}

	query := `
		CREATE INDEX IF NOT EXISTS "index_balance_items_on_date_and_category_id" ON "items" ("date", "category_id")
	`

	if _, err := m.db.ExecContext(ctx, query); err != nil {
		log.Fatalf("failed to add index_balance_items_on_date_and_category_id, error: %v", err)
	}

	m.schemaMigration.Save(ctx)
}
