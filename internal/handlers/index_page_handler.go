package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/category"
	"github.com/tksasha/balance/internal/components"
	"github.com/tksasha/balance/internal/responses"
)

type IndexPageHandler struct {
	categoryService category.Service
}

func NewIndexPageHandler(categoryService category.Service) *IndexPageHandler {
	return &IndexPageHandler{
		categoryService: categoryService,
	}
}

func (h *IndexPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(w, r); err != nil {
		if response, ok := w.(*responses.Response); ok {
			response.Error = err

			return
		}

		return
	}
}

func (h *IndexPageHandler) handle(w http.ResponseWriter, r *http.Request) error {
	categories, err := h.categoryService.List(r.Context())
	if err != nil {
		return err
	}

	return components.IndexPage(categories).Render(w)
}
