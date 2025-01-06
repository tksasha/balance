package migrations

import (
	"context"
	"database/sql"
	"log"
)

type CreateIndexBalanceItemsOnDate struct {
	db              *sql.DB
	schemaMigration *SchemaMigration
}

func NewCreateIndexBalanceItemsOnDate(db *sql.DB) *CreateIndexBalanceItemsOnDate {
	return &CreateIndexBalanceItemsOnDate{
		db:              db,
		schemaMigration: NewSchemaMigration(db, "19700101000003"),
	}
}

func (m *CreateIndexBalanceItemsOnDate) Up(ctx context.Context) {
	if m.schemaMigration.IsExist(ctx) {
		return
	}

	query := `
		CREATE INDEX "index_balance_items_on_date" ON "items" ("date")
	`

	if _, err := m.db.ExecContext(ctx, query); err != nil {
		log.Fatalf("failed to create index_balance_items_on_date, error: %v", err)
	}

	m.schemaMigration.Save(ctx)
}
