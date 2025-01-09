package handlers

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	internalerrors "github.com/tksasha/balance/internal/errors"
	"github.com/tksasha/balance/internal/models"
)

type EditCategoryHandler struct {
	categoryService CategoryService
}

func NewEditCategoryHandler(categoryService CategoryService) *EditCategoryHandler {
	return &EditCategoryHandler{
		categoryService: categoryService,
	}
}

func (h *EditCategoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	category, err := h.handle(r)
	if err != nil {
		if errors.Is(err, internalerrors.ErrResourceNotFound) {
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

func (h *EditCategoryHandler) handle(r *http.Request) (*models.Category, error) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		return nil, internalerrors.ErrResourceNotFound
	}

	category, err := h.categoryService.FindByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, internalerrors.ErrRecordNotFound) {
			return nil, internalerrors.ErrResourceNotFound
		}
	}

	return category, nil
}
