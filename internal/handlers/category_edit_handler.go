package handlers

import (
	"net/http"
	"strconv"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/responses"
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
		if response, ok := w.(*responses.Response); ok {
			response.Error = err

			return
		}

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
