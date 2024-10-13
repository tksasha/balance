package server

import (
	"context"
	"embed"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/tksasha/balance/internal/config"
	"github.com/tksasha/balance/internal/server/app"
	"github.com/tksasha/balance/internal/server/db"
	"github.com/tksasha/balance/internal/server/env"
	"github.com/tksasha/balance/internal/server/middlewares"
	"github.com/tksasha/balance/internal/server/routes"
	"github.com/tksasha/balance/internal/server/workdir"
)

//go:embed assets
var assets embed.FS

type Server struct{}

func Run(config *config.Config) {
	ctx := context.Background()

	workdir, err := workdir.New()
	if err != nil {
		log.Fatalf("workdir initialize error: %v", err)
	}

	db, err := db.Open(ctx, workdir, env.Get())
	if err != nil {
		log.Fatalf("open db error: %v", err)
	}

	app := app.New(db)

	router := routes.New(config, app, assets)

	handler := middlewares.NewCurrencyMiddleware(app).Wrap(router)

	slog.Info("starting server", "address", config.Address)

	server := http.Server{
		Addr:              config.Address,
		ReadHeaderTimeout: 1 * time.Second,
		Handler:           handler,
	}

	log.Fatal(
		server.ListenAndServe(),
	)
}
