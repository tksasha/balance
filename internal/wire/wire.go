//go:build wireinject

package wire

import (
	"context"

	"github.com/google/wire"
	"github.com/tksasha/balance/internal/cash"
	cashhandlers "github.com/tksasha/balance/internal/cash/handlers"
	cashrepository "github.com/tksasha/balance/internal/cash/repository"
	cashservice "github.com/tksasha/balance/internal/cash/service"
	"github.com/tksasha/balance/internal/category"
	categoryhandlers "github.com/tksasha/balance/internal/category/handlers"
	categoryrepository "github.com/tksasha/balance/internal/category/repository"
	categoryservice "github.com/tksasha/balance/internal/category/service"
	"github.com/tksasha/balance/internal/config"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/middlewares"
	"github.com/tksasha/balance/internal/providers"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/routes"
	"github.com/tksasha/balance/internal/server"
	"github.com/tksasha/balance/internal/services"
)

func InitializeServer() *server.Server {
	wire.Build(
		cashhandlers.NewCreateHandler,
		cashhandlers.NewDeleteHandler,
		cashhandlers.NewEditHandler,
		cashhandlers.NewListHandler,
		cashhandlers.NewNewHandler,
		cashhandlers.NewUpdateHandler,
		cashrepository.New,
		cashservice.New,
		categoryhandlers.NewCreateHandler,
		categoryhandlers.NewDeleteHandler,
		categoryhandlers.NewEditHandler,
		categoryhandlers.NewListHandler,
		categoryhandlers.NewUpdateHandler,
		categoryrepository.New,
		categoryservice.New,
		config.New,
		context.Background,
		db.Open,
		handlers.NewIndexPageHandler,
		handlers.NewItemCreateHandler,
		handlers.NewItemEditHandler,
		handlers.NewItemListHandler,
		handlers.NewItemUpdateHandler,
		middlewares.New,
		providers.NewDBNameProvider,
		repositories.NewItemRepository,
		routes.New,
		server.New,
		services.NewItemService,
		wire.Bind(new(cash.Repository), new(*cashrepository.Repository)),
		wire.Bind(new(cash.Service), new(*cashservice.Service)),
		wire.Bind(new(category.Repository), new(*categoryrepository.Repository)),
		wire.Bind(new(category.Service), new(*categoryservice.Service)),
		wire.Bind(new(db.DBNameProvider), new(*providers.DBNameProvider)),
		wire.Bind(new(handlers.ItemService), new(*services.ItemService)),
		wire.Bind(new(services.ItemRepository), new(*repositories.ItemRepository)),
	)

	return nil
}
