package utils

import (
	"context"
	"database/sql"
	"testing"

	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/providers"
)

func newDB(ctx context.Context, t *testing.T) *sql.DB {
	t.Helper()

	dbNameProvider := providers.NewDBNameProvider()

	return db.Open(ctx, dbNameProvider)
}

func Cleanup(ctx context.Context, t *testing.T) {
	t.Helper()

	t.Cleanup(func() {
		db := newDB(ctx, t)
		defer func() {
			_ = db.Close()
		}()

		tables := []string{"items", "categories", "cashes"}

		for _, table := range tables {
			if _, err := db.ExecContext(ctx, "DELETE FROM "+table); err != nil { //nolint:gosec
				t.Fatalf("failed to truncate %s, error: %v", table, err)
			}
		}
	})
}
