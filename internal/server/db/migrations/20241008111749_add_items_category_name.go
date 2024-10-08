package migrations

import (
	"context"
	"database/sql"
	"log"
)

type AddItemsCategoryNameMigration struct {
	db              *sql.DB
	schemaMigration *SchemaMigration
}

//nolint:ireturn
func NewAddItemsCategoryNameMigration(db *sql.DB) Migration {
	return &AddItemsCategoryNameMigration{
		db:              db,
		schemaMigration: NewSchemaMigration(db, "20241008111749"),
	}
}

func (m *AddItemsCategoryNameMigration) Up(ctx context.Context) {
	if m.schemaMigration.IsExist(ctx) {
		return
	}

	query := `
		ALTER TABLE
			items
		ADD COLUMN
			category_name VARCHAR
	`

	if _, err := m.db.ExecContext(ctx, query); err != nil {
		log.Fatalf("add items.category_name error: %v", err)
	}

	m.schemaMigration.Save(ctx)
}
