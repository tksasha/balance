package server

import (
	"embed"
	"html/template"
	"log"
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/config"
	"github.com/tksasha/balance/internal/handlers"
)

//go:embed application.html
var applicationHTML embed.FS

type Server struct{}

func Run(config *config.Config) {
	applicationTemplate, err := template.ParseFS(applicationHTML, "application.html")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("GET /{$}", handlers.NewIndexHandler(applicationTemplate))

	slog.Info("starting server", "address", config.Address)

	log.Fatal(
		http.ListenAndServe(config.Address, nil),
	)
}
