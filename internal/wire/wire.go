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
		handlers.NewCreateCategoryHandler,
		handlers.NewCreateItemHandler,
		handlers.NewEditCategoryHandler,
		handlers.NewGetCategoriesHandler,
		handlers.NewGetItemHandler,
		handlers.NewGetItemsHandler,
		handlers.NewIndexPageHandler,
		handlers.NewUpdateCategoryHandler,
		providers.NewDBNameProvider,
		repositories.NewCategoryRepository,
		repositories.NewItemRepository,
		routes.New,
		server.New,
		services.NewCategoryService,
		services.NewCreateItemService,
		services.NewItemService,
		wire.Bind(new(db.DBNameProvider), new(*providers.DBNameProvider)),
		wire.Bind(new(handlers.CategoryService), new(*services.CategoryService)),
		wire.Bind(new(handlers.ItemService), new(*services.ItemService)),
		wire.Bind(new(handlers.ItemCreator), new(*services.CreateItemService)),
		wire.Bind(new(services.CategoryRepository), new(*repositories.CategoryRepository)),
		wire.Bind(new(services.ItemRepository), new(*repositories.ItemRepository)),
	)

	return nil
}
