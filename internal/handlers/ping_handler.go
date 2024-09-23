package handlers

import (
	"html/template"
	"log/slog"
	"net/http"
	"time"

	"github.com/tksasha/balance/internal/server/app"
)

type PingHandler struct {
	tmpl *template.Template
}

func NewPingHandler(app *app.App) http.Handler {
	return &PingHandler{
		tmpl: app.T,
	}
}

func (h *PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)

	if err := h.tmpl.ExecuteTemplate(w, "pong-button", nil); err != nil {
		slog.Error(err.Error())
	}
}
