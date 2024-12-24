//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/routes"
	"github.com/tksasha/balance/internal/server"
	"github.com/tksasha/balance/internal/server/config"
	"github.com/tksasha/balance/internal/server/db"
	"github.com/tksasha/balance/internal/services"
)

func InitializeServer() *server.Server {
	wire.Build(
		config.New,
		db.Open,
		repositories.NewCategoryRepository,
		repositories.NewItemRepository,
		server.New,
		services.NewItemService,
		services.NewCategoryService,
		routes.New,
		handlers.New,
		wire.Bind(new(handlers.ItemService), new(*services.ItemService)),
		wire.Bind(new(handlers.CategoryService), new(*services.CategoryService)),
	)

	return nil
}
