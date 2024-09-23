package server

import (
	"embed"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/tksasha/balance/internal/config"
	"github.com/tksasha/balance/internal/server/app"
	"github.com/tksasha/balance/internal/server/db"
	"github.com/tksasha/balance/internal/server/env"
	"github.com/tksasha/balance/internal/server/routes"
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

	workdir, err := workdir.New()
	if err != nil {
		log.Fatalf("workdir initialize error: %v", err)
	}

	db, err := db.Open(workdir, env.Get())
	if err != nil {
		log.Fatalf("open db error: %v", err)
	}

	router := routes.New(
		&app.App{
			T:  tmpl,
			DB: db,
		},
		assets,
	)

	slog.Info("starting server", "address", config.Address)

	server := http.Server{
		Addr:              config.Address,
		ReadHeaderTimeout: 1 * time.Second,
		Handler:           router,
	}

	log.Fatal(
		server.ListenAndServe(),
	)
}
