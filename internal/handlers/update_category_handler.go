package handlers

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	internalerrors "github.com/tksasha/balance/internal/errors"
	"github.com/tksasha/balance/internal/models"
)

type UpdateCategoryHandler struct {
	categoryService CategoryService
}

func NewUpdateCategoryHandler(categoryService CategoryService) *UpdateCategoryHandler {
	return &UpdateCategoryHandler{
		categoryService: categoryService,
	}
}

func (h *UpdateCategoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	category, err := h.handle(r)
	if err != nil {
		if errors.Is(err, internalerrors.ErrResourceNotFound) {
			http.Error(w, "Resource Not Found", http.StatusNotFound)

			return
		}

		slog.Error("update category handler error", "error", err)

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

		return
	}

	_ = category

	_, _ = w.Write([]byte("update category page\n"))
}

func (h *UpdateCategoryHandler) handle(r *http.Request) (*models.Category, error) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		return nil, internalerrors.ErrResourceNotFound
	}

	category, err := h.categoryService.FindByID(r.Context(), id)
	if err != nil {
		return nil, err
	}

	if err := r.ParseForm(); err != nil {
		return category, err
	}

	category.Name = r.FormValue("name")

	if err := h.categoryService.Update(r.Context(), category); err != nil {
		return category, err
	}

	return category, nil
}
