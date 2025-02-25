package handlers

import (
	"errors"
	"net/http"

	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/balance/internal/backoffice/category/component"
	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/balance/internal/common/handler"
	"github.com/tksasha/validator"
)

type UpdateHandler struct {
	*handler.Handler

	categoryService   category.Service
	categoryComponent *component.CategoryComponent
}

func NewUpdateHandler(
	categoryService category.Service,
	categoryComponent *component.CategoryComponent,
) *UpdateHandler {
	return &UpdateHandler{
		Handler:           handler.New(),
		categoryService:   categoryService,
		categoryComponent: categoryComponent,
	}
}

func (h *UpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	category, err := h.handle(r)
	if err == nil {
		w.WriteHeader(http.StatusNoContent)

		return
	}

	var verrors validator.Errors
	if errors.As(err, &verrors) {
		err := h.categoryComponent.Update(category, verrors).Render(w)

		h.SetError(w, err)
	}

	h.SetError(w, err)
}

func (h *UpdateHandler) handle(r *http.Request) (*category.Category, error) {
	if err := r.ParseForm(); err != nil {
		return nil, common.ErrParsingForm
	}

	request := category.UpdateRequest{
		ID:            r.PathValue("id"),
		Name:          r.FormValue("name"),
		Income:        r.FormValue("income"),
		Visible:       r.FormValue("visible"),
		Supercategory: r.FormValue("supercategory"),
	}

	return h.categoryService.Update(r.Context(), request)
}
