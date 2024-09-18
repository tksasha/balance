package handlers

import (
	"html/template"
	"log/slog"
	"net/http"
)

type PongHandler struct {
	tmpl *template.Template
}

func NewPongHandler(tmpl *template.Template) http.Handler {
	return &PongHandler{
		tmpl: tmpl,
	}
}

func (h *PongHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.tmpl.ExecuteTemplate(w, "ping-button", nil); err != nil {
		slog.Error(err.Error())
	}
}
