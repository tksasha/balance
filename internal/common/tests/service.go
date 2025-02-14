package tests

import (
	"context"
	"database/sql"
	"testing"

	"github.com/tksasha/balance/internal/category"
	categoryrepository "github.com/tksasha/balance/internal/category/repository"
	categoryservice "github.com/tksasha/balance/internal/category/service"
	"github.com/tksasha/balance/internal/core/cash"
	cashrepository "github.com/tksasha/balance/internal/core/cash/repository"
	cashservice "github.com/tksasha/balance/internal/core/cash/service"
	"github.com/tksasha/balance/internal/item"
	itemrepository "github.com/tksasha/balance/internal/item/repository"
	itemservice "github.com/tksasha/balance/internal/item/service"
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

func NewItemService(ctx context.Context, t *testing.T) (item.Service, *sql.DB) {
	t.Helper()

	db := newDB(ctx, t)

	itemRepository := itemrepository.New(db)

	categoryRepository := categoryrepository.New(db)

	return itemservice.New(itemRepository, categoryRepository), db
}
