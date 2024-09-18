package handlers

import (
	"html/template"
	"log/slog"
	"net/http"
)

type PingHandler struct {
	tmpl *template.Template
}

func NewPingHandler(tmpl *template.Template) http.Handler {
	return &PingHandler{
		tmpl: tmpl,
	}
}

func (h *PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.tmpl.ExecuteTemplate(w, "pong-button", nil); err != nil {
		slog.Error(err.Error())
	}
}
