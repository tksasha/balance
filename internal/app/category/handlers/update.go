package handlers

import (
	"errors"
	"net/http"

	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/app/category/components"
	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/validator"
)

type UpdateHandler struct {
	*common.BaseHandler

	categoryService   category.Service
	categoryComponent *components.CategoryComponent
}

func NewUpdateHandler(
	baseHandler *common.BaseHandler,
	categoryService category.Service,
	categoryComponent *components.CategoryComponent,
) *UpdateHandler {
	return &UpdateHandler{
		BaseHandler:       baseHandler,
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
