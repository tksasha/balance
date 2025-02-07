//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/tksasha/balance/internal/config"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/providers"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/routes"
	"github.com/tksasha/balance/internal/server"
	"github.com/tksasha/balance/internal/services"
)

func InitializeServer() *server.Server {
	wire.Build(
		config.New,
		db.Open,
		handlers.NewCashCreateHandler,
		handlers.NewCashEditHandler,
		handlers.NewCategoryCreateHandler,
		handlers.NewEditCategoryHandler,
		handlers.NewGetCategoriesHandler,
		handlers.NewGetItemHandler,
		handlers.NewGetItemsHandler,
		handlers.NewIndexPageHandler,
		handlers.NewItemCreateHandler,
		handlers.NewUpdateCategoryHandler,
		providers.NewDBNameProvider,
		repositories.NewCashRepository,
		repositories.NewCategoryRepository,
		repositories.NewItemRepository,
		routes.New,
		server.New,
		services.NewCashService,
		services.NewCategoryService,
		services.NewItemService,
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
