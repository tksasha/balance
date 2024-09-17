package handlers

import (
	"html/template"
	"log/slog"
	"net/http"
)

type IndexHandler struct {
	applicationTemplate *template.Template
}

func NewIndexHandler(applicationTemplate *template.Template) http.Handler {
	return &IndexHandler{
		applicationTemplate: applicationTemplate,
	}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.applicationTemplate.Execute(w, nil); err != nil {
		slog.Error(err.Error())
	}
}
