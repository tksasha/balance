package routes

import (
	"embed"
	"net/http"

	"github.com/tksasha/balance/internal/handlers"
)

//go:embed assets
var assets embed.FS

func New(handlers *handlers.Handlers) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /assets/{$}", http.RedirectHandler("/", http.StatusMovedPermanently))

	mux.Handle("GET /assets/", http.FileServerFS(assets))

	mux.Handle("GET /", handlers.Index)

	mux.Handle("GET /items", handlers.GetItems)

	mux.Handle("POST /items", handlers.CreateItem)

	mux.Handle("GET /items/{id}/edit", handlers.EditItem)

	mux.Handle("PATCH /items/{id}", handlers.UpdateItem)

	mux.Handle("DELETE /items/{id}", handlers.DeleteItem)

	return mux
}
