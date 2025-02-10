package handlers

import (
	"errors"
	"net/http"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/requests"
	"github.com/tksasha/balance/pkg/validation"
)

type CategoryCreateHandler struct {
	categoryService CategoryService
}

func NewCategoryCreateHandler(categoryService CategoryService) *CategoryCreateHandler {
	return &CategoryCreateHandler{
		categoryService: categoryService,
	}
}

func (h *CategoryCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(r); err != nil {
		var verrors validation.Errors
		if errors.As(err, &verrors) {
			_, _ = w.Write([]byte(verrors.Error()))

			return
		}

		return
	}

	_, _ = w.Write([]byte("create category page\n"))
}

func (h *CategoryCreateHandler) handle(r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return apperrors.ErrParsingForm
	}

	request := requests.CategoryCreateRequest{
		Name:          r.FormValue("name"),
		Income:        r.FormValue("income"),
		Visible:       r.FormValue("visible"),
		Supercategory: r.FormValue("supercategory"),
	}

	return h.categoryService.Create(r.Context(), request)
}
