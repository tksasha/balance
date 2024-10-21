package server

import (
	"embed"
	"net/http"

	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/middlewares"
)

//go:embed assets
var assets embed.FS

type Router struct {
	indexPageHandler   *handlers.IndexPageHandler
	createItemHandler  *handlers.CreateItemHandler
	recoveryMiddleware *middlewares.RecoveryMiddleware
	currencyMiddleware *middlewares.CurrencyMiddleware
}

func NewRouter(
	indexPageHandler *handlers.IndexPageHandler,
	createItemHandler *handlers.CreateItemHandler,
	recoveryMiddleware *middlewares.RecoveryMiddleware,
	currencyMiddleware *middlewares.CurrencyMiddleware,
) *Router {
	return &Router{
		indexPageHandler:   indexPageHandler,
		createItemHandler:  createItemHandler,
		recoveryMiddleware: recoveryMiddleware,
		currencyMiddleware: currencyMiddleware,
	}
}

func (r *Router) GetHandler() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /assets/{$}", http.RedirectHandler("/", http.StatusMovedPermanently))

	mux.Handle("GET /assets/", http.FileServerFS(assets))

	mux.Handle("GET /", r.indexPageHandler)

	mux.Handle("POST /items", r.createItemHandler)

	handler := r.recoveryMiddleware.Wrap(
		r.currencyMiddleware.Wrap(mux),
	)

	return handler
}
