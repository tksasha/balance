package handlers

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/models"
)

type CategoryEditHandler struct {
	categoryService CategoryService
}

func NewCategoryEditHandler(categoryService CategoryService) *CategoryEditHandler {
	return &CategoryEditHandler{
		categoryService: categoryService,
	}
}

func (h *CategoryEditHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	category, err := h.handle(r)
	if err != nil {
		if errors.Is(err, apperrors.ErrResourceNotFound) {
			http.Error(w, "Resource Not Found", http.StatusNotFound)

			return
		}

		slog.Error("edit category handler error", "error", err)

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

		return
	}

	_ = category

	_, _ = w.Write([]byte("edit category form\n"))
}

func (h *CategoryEditHandler) handle(r *http.Request) (*models.Category, error) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		return nil, apperrors.ErrResourceNotFound
	}

	category, err := h.categoryService.FindByID(r.Context(), id)
	if err != nil {
		return nil, err
	}

	return category, nil
}
