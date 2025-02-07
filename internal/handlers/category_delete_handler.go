package handlers

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/models"
)

type CategoryDeleteHandler struct {
	categoryService CategoryService
}

func NewCategoryDeleteHandler(categoryService CategoryService) *CategoryDeleteHandler {
	return &CategoryDeleteHandler{
		categoryService: categoryService,
	}
}

func (h *CategoryDeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	category, err := h.handle(r)
	if err != nil {
		if errors.Is(err, apperrors.ErrResourceNotFound) {
			http.Error(w, "Resource Not Found", http.StatusNotFound)

			return
		}

		slog.Error("delete category handler error", "error", err)

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

		return
	}

	_ = category

	_, _ = w.Write([]byte("delete category page\n"))
}

func (h *CategoryDeleteHandler) handle(r *http.Request) (*models.Category, error) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		return nil, apperrors.ErrResourceNotFound
	}

	category, err := h.categoryService.FindByID(r.Context(), id)
	if err != nil {
		return nil, err
	}

	if err := h.categoryService.Delete(r.Context(), category); err != nil {
		return category, err
	}

	return category, nil
}
