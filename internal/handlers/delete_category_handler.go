package handlers

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	internalerrors "github.com/tksasha/balance/internal/errors"
	"github.com/tksasha/balance/internal/models"
)

type DeleteCategoryHandler struct {
	categoryService CategoryService
}

func NewDeleteCategoryHandler(categoryService CategoryService) *DeleteCategoryHandler {
	return &DeleteCategoryHandler{
		categoryService: categoryService,
	}
}

func (h *DeleteCategoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	category, err := h.handle(r)
	if err != nil {
		if errors.Is(err, internalerrors.ErrResourceNotFound) {
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

func (h *DeleteCategoryHandler) handle(r *http.Request) (*models.Category, error) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		return nil, internalerrors.ErrResourceNotFound
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
