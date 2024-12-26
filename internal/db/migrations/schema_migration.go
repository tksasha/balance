package migrations

import (
	"context"
	"database/sql"
	"log"
)

type SchemaMigration struct {
	db      *sql.DB
	version string
}

func NewSchemaMigration(db *sql.DB, version string) *SchemaMigration {
	return &SchemaMigration{
		db:      db,
		version: version,
	}
}

func (m *SchemaMigration) IsExist(ctx context.Context) bool {
	query := `
		SELECT
			COUNT(*)
		FROM
			schema_migrations
		WHERE
			version=?
	`

	row := m.db.QueryRowContext(ctx, query, m.version)

	var count int

	if err := row.Scan(&count); err != nil {
		log.Fatalf("count schema_migrations error: %v", err)
	}

	switch count {
	case 1:
		return true
	case 0:
		return false
	default:
		log.Fatalf("broken schema_migrations, count = %d", count)

		return false
	}
}

func (m *SchemaMigration) Save(ctx context.Context) {
	query := `
		INSERT INTO
			schema_migrations(version)
		VALUES(?)
	`

	if _, err := m.db.ExecContext(ctx, query, m.version); err != nil {
		log.Fatalf("save schema_migrations error: %v", err)
	}
}
