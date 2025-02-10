package handlers

import (
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/components"
	"github.com/tksasha/balance/internal/responses"
)

type IndexPageHandler struct {
	categoryService CategoryService
}

func NewIndexPageHandler(categoryService CategoryService) *IndexPageHandler {
	return &IndexPageHandler{
		categoryService: categoryService,
	}
}

func (h *IndexPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(w, r); err != nil {
		slog.Info("debug", "error", err)

		if response, ok := w.(*responses.Response); ok {
			response.Error = err

			return
		}

		return
	}
}

func (h *IndexPageHandler) handle(w http.ResponseWriter, r *http.Request) error {
	categories, err := h.categoryService.GetAll(r.Context())
	if err != nil {
		return err
	}

	return components.IndexPage(categories).Render(w)
}
