package server

import (
	"context"
	"embed"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tksasha/balance/internal/config"
	"github.com/tksasha/balance/internal/server/app"
	"github.com/tksasha/balance/internal/server/db"
	"github.com/tksasha/balance/internal/server/env"
	"github.com/tksasha/balance/internal/server/middlewares"
	"github.com/tksasha/balance/internal/server/routes"
	"github.com/tksasha/balance/internal/server/workdir"
)

const shutdownTimeout = 5 * time.Second

//go:embed assets
var assets embed.FS

type Server struct{}

func Run(config *config.Config) {
	ctx := context.Background()

	db, err := db.Open(ctx, workdir.New(), env.Get())
	if err != nil {
		log.Fatalf("open db error: %v", err)
	}

	app := app.New(db)

	router := routes.New(config, app, assets)

	handler := middlewares.RecoveryMiddleware(
		middlewares.NewCurrencyMiddleware(app).Wrap(router),
	)

	server := http.Server{
		Addr:              config.Address,
		ReadHeaderTimeout: 1 * time.Second,
		Handler:           handler,
	}

	startingServerErrors := make(chan error, 1)

	go func() {
		slog.Info("server started")

		startingServerErrors <- server.ListenAndServe()
	}()

	signals := make(chan os.Signal, 1)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-startingServerErrors:
		slog.Error("failed to start server", "error", err)
	case <-signals:
		slog.Info("trying to stop server")

		ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			slog.Error("failed to shutdown server", "error", err)

			if err := server.Close(); err != nil {
				slog.Error("failed to close server", "error", err)
			}
		}
	}

	slog.Info("server stopped")

	time.Sleep(1 * time.Second) // to flush logs
}
