package server

import (
	"context"
	"embed"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tksasha/balance/internal/server/config"
)

//go:embed assets
var assets embed.FS

type Server struct {
	httpServer            *http.Server
	shutDownServerTimeout time.Duration
}

func New() Server {
	mux := http.NewServeMux()

	mux.Handle("GET /assets/{$}", http.RedirectHandler("/", http.StatusMovedPermanently))
	mux.Handle("GET /assets/", http.FileServerFS(assets))

	config := config.New()

	httpServer := &http.Server{
		Addr:              config.Address,
		ReadHeaderTimeout: config.ReadHeaderTimeout,
		Handler:           mux,
	}

	return Server{
		httpServer:            httpServer,
		shutDownServerTimeout: config.ShutDownServerTimeout,
	}
}

func (s Server) Run() {
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
