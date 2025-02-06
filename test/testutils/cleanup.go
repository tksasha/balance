package testutils

import (
	"context"
	"testing"

	"github.com/tksasha/balance/internal/db"
)

func Cleanup(ctx context.Context, t *testing.T, db *db.DB) {
	t.Helper()

	t.Cleanup(func() {
		tables := []string{"items", "categories", "cashes"}

		for _, table := range tables {
			if _, err := db.Connection.ExecContext(ctx, "DELETE FROM "+table); err != nil { //nolint:gosec
				t.Fatalf("failed to truncate %s, error: %v", table, err)
			}
		}
	})
}
