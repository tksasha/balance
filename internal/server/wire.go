//go:build wireinject

package server

import (
	"github.com/google/wire"
	"github.com/tksasha/balance/internal/config"
	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/middlewares"
	"github.com/tksasha/balance/internal/repositories"
)

func InitializeServer() *Server {
	wire.Build(
		New,
		config.New,
		handlers.NewIndexPageHandler,
		handlers.NewCreateItemHandler,
		middlewares.NewCurrencyMiddleware,
		middlewares.NewRecoveryMiddleware,
		NewRouter,
		repositories.NewItemRepository,
		wire.Bind(new(repositories.ItemCreator), new(*repositories.ItemRepository)),
	)

	return &Server{}
}
