package handlers

import (
	"html/template"
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/server/app"
)

type IndexHandler struct {
	tmpl *template.Template
}

func NewIndexHandler(app *app.App) http.Handler {
	return &IndexHandler{
		tmpl: app.T,
	}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.tmpl.ExecuteTemplate(w, "index-page", nil); err != nil {
		slog.Error(err.Error())
	}
}
