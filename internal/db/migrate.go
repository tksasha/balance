package db

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"log"
	"maps"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
)

//go:embed migrations/*.sql
var fs embed.FS

var ErrInvalidVersion = errors.New("invalid version")

type migration struct {
	db *sql.DB
}

func newMigration(db *sql.DB) migration {
	return migration{
		db: db,
	}
}

func (m migration) run(ctx context.Context) error {
	migrations, err := m.migrations()
	if err != nil {
		return err
	}

	versions := slices.Collect(
		maps.Keys(migrations),
	)

	slices.Sort(versions)

	for _, version := range versions {
		exists, err := m.exists(ctx, version)
		if err != nil {
			return err
		}

		if exists {
			continue
		}

		log.Printf("run migration: %s", version)

		if err := m.migrate(ctx, version, migrations[version]); err != nil {
			return err
		}
	}

	return nil
}

func (m migration) migrations() (map[string]string, error) {
	root := "migrations"

	entries, err := fs.ReadDir(root)
	if err != nil {
		return nil, err
	}

	migrations := map[string]string{}

	for _, entry := range entries {
		filename := filepath.Join(root, entry.Name())

		data, err := fs.ReadFile(filename)
		if err != nil {
			return nil, err
		}

		version, err := m.version(entry.Name())
		if err != nil {
			return nil, err
		}

		query := strings.TrimSpace(string(data))

		if query != "" {
			migrations[version] = query
		}
	}

	return migrations, nil
}

func (m migration) version(input string) (string, error) {
	version := strings.Split(input, "_")[0]

	match, err := regexp.MatchString(`\d{14}`, version)
	if err != nil {
		return "", err
	}

	if !match {
		return "", fmt.Errorf("%w: %s", ErrInvalidVersion, version)
	}

	return version, nil
}

func (m migration) exists(ctx context.Context, version string) (bool, error) {
	var exists int

	row := m.db.QueryRowContext(ctx, `SELECT 1 FROM schema_migrations WHERE version = ?`, version)

	if err := row.Scan(&exists); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, err
	}

	return exists == 1, nil
}

func (m migration) migrate(ctx context.Context, version, query string) error {
	if _, err := m.db.ExecContext(ctx, query); err != nil {
		return err
	}

	if _, err := m.db.ExecContext(ctx, `INSERT INTO schema_migrations(version) VALUES(?)`, version); err != nil {
		return err
	}

	return nil
}
