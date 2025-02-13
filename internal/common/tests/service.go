package tests

import (
	"context"
	"database/sql"
	"testing"

	"github.com/tksasha/balance/internal/cash"
	cashrepository "github.com/tksasha/balance/internal/cash/repository"
	cashservice "github.com/tksasha/balance/internal/cash/service"
	"github.com/tksasha/balance/internal/category"
	categoryrepository "github.com/tksasha/balance/internal/category/repository"
	categoryservice "github.com/tksasha/balance/internal/category/service"
)

func NewCashService(ctx context.Context, t *testing.T) (cash.Service, *sql.DB) {
	t.Helper()

	db := newDB(ctx, t)

	repository := cashrepository.New(db)

	return cashservice.New(repository), db
}

func NewCategoryService(ctx context.Context, t *testing.T) (category.Service, *sql.DB) {
	t.Helper()

	db := newDB(ctx, t)

	repository := categoryrepository.New(db)

	return categoryservice.New(repository), db
}
