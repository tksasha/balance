package handlers

import (
	"embed"
	"html/template"
	"log"
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/models"
)

//go:embed templates/items.html
var itemsFS embed.FS

type GetItemsHandler struct {
	tmpl *template.Template
}

func NewGetItemsHandler() http.Handler {
	tmpl, err := template.ParseFS(itemsFS, "templates/items.html")
	if err != nil {
		log.Fatalf("parse items.html error: %v", err)
	}

	return &GetItemsHandler{
		tmpl: tmpl,
	}
}

func (h *GetItemsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var items []models.Item

	items = append(items, models.Item{ID: 1, Name: "John McClane", Age: 42})
	items = append(items, models.Item{ID: 2, Name: "Bruce Wayne", Age: 69})
	items = append(items, models.Item{ID: 3, Name: "Peter Parker", Age: 18})

	if err := h.tmpl.ExecuteTemplate(w, "item-list", items); err != nil {
		slog.Error(err.Error())
	}
}
