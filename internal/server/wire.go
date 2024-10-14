//go:build wireinject

package server

import (
	"github.com/google/wire"
	"github.com/tksasha/balance/internal/config"
	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/server/app"
	"github.com/tksasha/balance/internal/server/db"
	"github.com/tksasha/balance/internal/server/middlewares"
	"github.com/tksasha/balance/internal/server/routes"
)

func Initialize() *Server {
	wire.Build(
		New,
		app.New,
		config.New,
		db.Open,
		handlers.New,
		middlewares.New,
		routes.New,
	)

	return &Server{}
}
