package db

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tksasha/balance/internal/db/migrations"
)

type DB struct {
	Connection *sql.DB
}

func Open(dbNameProvider DBNameProvider) *DB {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", dbNameProvider.Provide())
	if err != nil {
		panic(err)
	}

	if err := db.PingContext(ctx); err != nil {
		if err := db.Close(); err != nil {
			panic(err)
		}

		panic(err)
	}

	if _, err := db.ExecContext(ctx, "PRAGMA foreign_keys=true"); err != nil {
		panic(err)
	}

	migrate(ctx, db)

	return &DB{
		Connection: db,
	}
}

func migrate(ctx context.Context, db *sql.DB) {
	migrations.NewCreateItemsMigration(db).Up(ctx)
	migrations.NewCreateIndexBalanceItemsOnDateAndCategoryID(db).Up(ctx)
	migrations.NewCreateIndexBalanceItemsOnDate(db).Up(ctx)
	migrations.NewCreateCategoriesMigration(db).Up(ctx)
	migrations.NewCreateIndexCategoriesOnNameAndCurrency(db).Up(ctx)
	migrations.NewAddItemsCategoryNameMigration(db).Up(ctx)
	migrations.NewChangeTypeOfDeletedAtInCategories(db).Up(ctx)
}
