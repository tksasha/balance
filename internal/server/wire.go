//go:build wireinject

package server

import (
	"github.com/google/wire"
	"github.com/tksasha/balance/internal/config"
	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/middlewares"
)

func InitializeServer() *Server {
	wire.Build(
		New,
		config.New,
		handlers.NewIndexPageHandler,
		middlewares.NewCurrencyMiddleware,
		middlewares.NewRecoveryMiddleware,
		NewRouter,
	)

	return &Server{}
}
