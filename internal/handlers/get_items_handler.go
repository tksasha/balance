package handlers

import (
	"html/template"
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/models"
)

type GetItemsHandler struct {
	tmpl *template.Template
}

func NewGetItemsHandler(tmpl *template.Template) http.Handler {
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
