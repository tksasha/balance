package server

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/config"
)

type Server struct{}

func Run(config *config.Config) {
	slog.Info("starting server", "address", config.Address)

	log.Fatal(
		http.ListenAndServe(config.Address, nil),
	)
}
