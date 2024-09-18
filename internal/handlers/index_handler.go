package handlers

import (
	"html/template"
	"log/slog"
	"net/http"
)

type IndexHandler struct {
	tmpl *template.Template
}

func NewIndexHandler(tmpl *template.Template) http.Handler {
	return &IndexHandler{
		tmpl: tmpl,
	}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.tmpl.Execute(w, nil); err != nil {
		slog.Error(err.Error())
	}
}
