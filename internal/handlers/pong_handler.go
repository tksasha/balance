package handlers

import (
	"html/template"
	"log/slog"
	"net/http"
	"time"

	"github.com/tksasha/balance/internal/server/app"
)

type PongHandler struct {
	tmpl *template.Template
}

func NewPongHandler(app *app.App) http.Handler {
	return &PongHandler{
		tmpl: app.T,
	}
}

func (h *PongHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)

	if err := h.tmpl.ExecuteTemplate(w, "ping-button", nil); err != nil {
		slog.Error(err.Error())
	}
}
