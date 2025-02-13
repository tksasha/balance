//go:build wireinject

package wire

import (
	"context"

	"github.com/google/wire"
	"github.com/tksasha/balance/internal/cash"
	cashhandlers "github.com/tksasha/balance/internal/cash/handlers"
	cashrepository "github.com/tksasha/balance/internal/cash/repository"
	cashservice "github.com/tksasha/balance/internal/cash/service"
	"github.com/tksasha/balance/internal/config"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/handlers/cashes"
	"github.com/tksasha/balance/internal/middlewares"
	"github.com/tksasha/balance/internal/providers"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/routes"
	"github.com/tksasha/balance/internal/server"
	"github.com/tksasha/balance/internal/services"
)

func InitializeServer() *server.Server {
	wire.Build(
		cashes.NewNewHandler,
		cashhandlers.NewCreateHandler,
		cashhandlers.NewListHandler,
		cashhandlers.NewUpdateHandler,
		cashrepository.New,
		cashservice.New,
		config.New,
		context.Background,
		db.Open,
		handlers.NewCashDeleteHandler,
		handlers.NewCashEditHandler,
		handlers.NewCategoryCreateHandler,
		handlers.NewCategoryDeleteHandler,
		handlers.NewCategoryEditHandler,
		handlers.NewCategoryListHandler,
		handlers.NewCategoryUpdateHandler,
		handlers.NewIndexPageHandler,
		handlers.NewItemCreateHandler,
		handlers.NewItemEditHandler,
		handlers.NewItemListHandler,
		handlers.NewItemUpdateHandler,
		middlewares.New,
		providers.NewDBNameProvider,
		repositories.NewCashRepository,
		repositories.NewCategoryRepository,
		repositories.NewItemRepository,
		routes.New,
		server.New,
		services.NewCashService,
		services.NewCategoryService,
		services.NewItemService,
		wire.Bind(new(cash.Repository), new(*cashrepository.Repository)),
		wire.Bind(new(cash.Service), new(*cashservice.Service)),
		wire.Bind(new(db.DBNameProvider), new(*providers.DBNameProvider)),
		wire.Bind(new(handlers.CashService), new(*services.CashService)),
		wire.Bind(new(handlers.CategoryService), new(*services.CategoryService)),
		wire.Bind(new(handlers.ItemService), new(*services.ItemService)),
		wire.Bind(new(services.CashRepository), new(*repositories.CashRepository)),
		wire.Bind(new(services.CategoryRepository), new(*repositories.CategoryRepository)),
		wire.Bind(new(services.ItemRepository), new(*repositories.ItemRepository)),
	)

	return nil
}
