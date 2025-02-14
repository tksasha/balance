package server

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tksasha/balance/internal/server/config"
	"github.com/tksasha/balance/internal/server/middlewares"
	"github.com/tksasha/balance/internal/server/routes"
)

type Server struct {
	httpServer            *http.Server
	shutDownServerTimeout time.Duration
}

func New(
	config *config.Config,
	routes *routes.Routes,
	middlewares []middlewares.Middleware,
) *Server {
	next := http.Handler(routes.Mux)

	for _, middleware := range middlewares {
		next = middleware.Wrap(next)
	}

	httpServer := &http.Server{
		Addr:              config.Address,
		ReadHeaderTimeout: config.ReadHeaderTimeout,
		Handler:           next,
	}

	return &Server{
		httpServer:            httpServer,
		shutDownServerTimeout: config.ShutDownServerTimeout,
	}
}

func (s *Server) Run() {
	slog.Info("Starting server...")

	go func() {
		if err := s.httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			slog.Error("Server start error", "error", err)
		}
	}()

	slog.Info("Server started")

	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sigChan)

	<-sigChan

	ctx, cancel := context.WithTimeout(context.Background(), s.shutDownServerTimeout)
	defer cancel()

	slog.Info("Shutting down server...")

	if err := s.httpServer.Shutdown(ctx); err != nil {
		slog.Error("Server shutdown error", "error", err)

		slog.Info("Forcing server close...")

		if err := s.httpServer.Close(); !errors.Is(err, http.ErrServerClosed) {
			slog.Error("Server close error", "error", err)
		}
	}

	slog.Info("Server stopped")
}
