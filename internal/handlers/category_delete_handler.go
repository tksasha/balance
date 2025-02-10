package handlers

import (
	"net/http"
	"strconv"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/responses"
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
		if response, ok := w.(*responses.Response); ok {
			response.Error = err

			return
		}

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
