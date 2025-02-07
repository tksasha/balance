package handlers

import (
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/components"
)

type CategoryListHandler struct {
	categoryService CategoryService
}

func NewCategoryListHandler(categoryService CategoryService) *CategoryListHandler {
	return &CategoryListHandler{
		categoryService: categoryService,
	}
}

func (h *CategoryListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(w, r); err != nil {
		slog.Error("get categories handler error", "error", err)

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *CategoryListHandler) handle(w http.ResponseWriter, r *http.Request) error {
	categories, err := h.categoryService.GetAll(r.Context())
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
