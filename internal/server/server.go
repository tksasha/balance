package server

import (
	"embed"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/tksasha/balance/internal/config"
	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/server/workdir"
)

//go:embed templates
var fs embed.FS

//go:embed assets
var assets embed.FS

type Server struct{}

func Run(config *config.Config) {
	tmpl, err := template.ParseFS(fs, "templates/*.html")
	if err != nil {
		log.Fatalf("parse templates/*.html error: %v", err)
	}

	workDir, err := workdir.New()
	if err != nil {
		log.Fatalf("workdir initialize error: %v", err)
	}

	_ = workDir

	http.Handle("GET /{$}", handlers.NewIndexHandler(tmpl))

	http.Handle("GET /assets/{$}", http.RedirectHandler("/", http.StatusMovedPermanently))

	http.Handle("GET /assets/", http.FileServerFS(assets))

	http.Handle("GET /items", handlers.NewGetItemsHandler(tmpl))

	http.Handle("GET /ping", handlers.NewPingHandler(tmpl))

	http.Handle("GET /pong", handlers.NewPongHandler(tmpl))

	slog.Info("starting server", "address", config.Address)

	server := http.Server{
		Addr:              config.Address,
		ReadHeaderTimeout: 1 * time.Second,
	}

	log.Fatal(
		server.ListenAndServe(),
	)
}
