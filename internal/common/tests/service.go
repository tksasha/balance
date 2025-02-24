package tests

import (
	"context"
	"database/sql"
	"testing"

	"github.com/tksasha/balance/internal/app/cash"
	cashrepository "github.com/tksasha/balance/internal/app/cash/repository"
	cashservice "github.com/tksasha/balance/internal/app/cash/service"
	categoryrepository "github.com/tksasha/balance/internal/app/category/repository"
	"github.com/tksasha/balance/internal/app/index"
	indexrepository "github.com/tksasha/balance/internal/app/index/repository"
	indexservice "github.com/tksasha/balance/internal/app/index/service"
	"github.com/tksasha/balance/internal/app/item"
	itemrepository "github.com/tksasha/balance/internal/app/item/repository"
	itemservice "github.com/tksasha/balance/internal/app/item/service"
	"github.com/tksasha/balance/internal/common"
)

func NewCashService(ctx context.Context, t *testing.T) (cash.Service, *sql.DB) {
	t.Helper()

	baseRepository := common.NewBaseRepository()

	db := newDB(ctx, t)

	repository := cashrepository.New(baseRepository, db)

	return cashservice.New(repository), db
}

func NewItemService(ctx context.Context, t *testing.T) (item.Service, *sql.DB) {
	t.Helper()

	baseRepository := common.NewBaseRepository()

	db := newDB(ctx, t)

	itemRepository := itemrepository.New(baseRepository, db)

	categoryRepository := categoryrepository.New(baseRepository, db)

	return itemservice.New(common.NewBaseService(), itemRepository, categoryRepository), db
}

func NewIndexPageService(ctx context.Context, t *testing.T) (index.Service, *sql.DB) {
	t.Helper()

	baseRepository := common.NewBaseRepository()

	db := newDB(ctx, t)

	repository := indexrepository.New(baseRepository, db)

	return indexservice.New(repository), db
}
