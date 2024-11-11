package handlers

import (
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/components"
	"github.com/tksasha/balance/internal/services"
)

type getCategoriesHandler struct {
	categoryService services.CategoryService
}

func NewGetCategoriesHandler(categoryService services.CategoryService) Handler {
	return &getCategoriesHandler{
		categoryService: categoryService,
	}
}

func (h *getCategoriesHandler) Pattern() string {
	return "GET /categories"
}

func (h *getCategoriesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(w, r); err != nil {
		slog.Error("get categories handler error", "error", err)

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *getCategoriesHandler) handle(w http.ResponseWriter, r *http.Request) error {
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
