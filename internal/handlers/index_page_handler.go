package handlers

import (
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/components"
)

type IndexPageHandler struct{}

func NewIndexPageHandler() *IndexPageHandler {
	return &IndexPageHandler{}
}

func (h *IndexPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(w, r); err != nil {
		slog.Error("index page handler error", "error", err)

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *IndexPageHandler) handle(w http.ResponseWriter, _ *http.Request) error {
	return components.IndexPageComponent().Render(w)
}
