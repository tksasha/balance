package tests

import (
	"context"
	"database/sql"
	"testing"

	"github.com/tksasha/balance/internal/core/cash"
	cashrepository "github.com/tksasha/balance/internal/core/cash/repository"
	cashservice "github.com/tksasha/balance/internal/core/cash/service"
	"github.com/tksasha/balance/internal/core/category"
	categoryrepository "github.com/tksasha/balance/internal/core/category/repository"
	categoryservice "github.com/tksasha/balance/internal/core/category/service"
	"github.com/tksasha/balance/internal/core/common"
	"github.com/tksasha/balance/internal/core/indexpage"
	indexpagerepository "github.com/tksasha/balance/internal/core/indexpage/repository"
	indexpageservice "github.com/tksasha/balance/internal/core/indexpage/service"
	"github.com/tksasha/balance/internal/core/item"
	itemrepository "github.com/tksasha/balance/internal/core/item/repository"
	itemservice "github.com/tksasha/balance/internal/core/item/service"
)

func NewCashService(ctx context.Context, t *testing.T) (cash.Service, *sql.DB) {
	t.Helper()

	baseRepository := common.NewBaseRepository()

	db := newDB(ctx, t)

	repository := cashrepository.New(baseRepository, db)

	return cashservice.New(repository), db
}

func NewCategoryService(ctx context.Context, t *testing.T) (category.Service, *sql.DB) {
	t.Helper()

	baseService := common.NewBaseService()

	baseRepository := common.NewBaseRepository()

	db := newDB(ctx, t)

	repository := categoryrepository.New(baseRepository, db)

	return categoryservice.New(baseService, repository), db
}

func NewItemService(ctx context.Context, t *testing.T) (item.Service, *sql.DB) {
	t.Helper()

	baseRepository := common.NewBaseRepository()

	db := newDB(ctx, t)

	itemRepository := itemrepository.New(baseRepository, db)

	categoryRepository := categoryrepository.New(baseRepository, db)

	return itemservice.New(common.NewBaseService(), itemRepository, categoryRepository), db
}

func NewIndexPageService(ctx context.Context, t *testing.T) (indexpage.Service, *sql.DB) {
	t.Helper()

	baseRepository := common.NewBaseRepository()

	db := newDB(ctx, t)

	repository := indexpagerepository.New(baseRepository, db)

	return indexpageservice.New(repository), db
}
