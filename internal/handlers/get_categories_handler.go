package handlers

import (
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/components"
)

type GetCategoriesHandler struct {
	categoryService CategoryService
}

func NewGetCategoriesHandler(categoryService CategoryService) *GetCategoriesHandler {
	return &GetCategoriesHandler{
		categoryService: categoryService,
	}
}

func (h *GetCategoriesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(w, r); err != nil {
		slog.Error("get categories handler error", "error", err)

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *GetCategoriesHandler) handle(w http.ResponseWriter, r *http.Request) error {
	categories, err := h.categoryService.GetCategories(r.Context())
	if err != nil {
		slog.Error("get categories error", "error", err)

		return err
	}

	if err := components.Categories(categories).Render(w); err != nil {
		slog.Error("render categories component error", "error", err)

		return err
	}

	return nil
}
