package server

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tksasha/balance/internal/config"
	"github.com/tksasha/balance/internal/server/app"
)

const shutdownTimeout = 5 * time.Second

type Server struct {
	app     *app.App
	config  *config.Config
	routes  *http.ServeMux
	handler http.Handler
}

func New(app *app.App, config *config.Config, routes *http.ServeMux, handler http.Handler) *Server {
	return &Server{
		app:     app,
		config:  config,
		routes:  routes,
		handler: handler,
	}
}

func (s *Server) Run() {
	server := http.Server{
		Addr:              s.config.Address,
		ReadHeaderTimeout: 1 * time.Second,
		Handler:           s.handler,
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
