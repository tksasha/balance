package handlers

import (
	"errors"
	"net/http"

	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/balance/internal/backoffice/category/component"
	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/balance/internal/common/handler"
	"github.com/tksasha/validation"
)

type CreateHandler struct {
	*handler.Handler

	categoryService   category.Service
	categoryComponent *component.CategoryComponent
}

func NewCreateHandler(
	categoryService category.Service,
	categoryComponent *component.CategoryComponent,
) *CreateHandler {
	return &CreateHandler{
		Handler:           handler.New(),
		categoryService:   categoryService,
		categoryComponent: categoryComponent,
	}
}

func (h *CreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	category, err := h.handle(r)
	if err == nil {
		w.WriteHeader(http.StatusCreated)

		return
	}

	var verrors validation.Errors
	if errors.As(err, &verrors) {
		err := h.categoryComponent.Create(category, verrors).Render(w)

		h.SetError(w, err)

		return
	}

	h.SetError(w, err)
}

func (h *CreateHandler) handle(r *http.Request) (*category.Category, error) {
	if err := r.ParseForm(); err != nil {
		return nil, common.ErrParsingForm
	}

	request := category.CreateRequest{
		Name:          r.FormValue("name"),
		Income:        r.FormValue("income"),
		Visible:       r.FormValue("visible"),
		Supercategory: r.FormValue("supercategory"),
	}

	return h.categoryService.Create(r.Context(), request)
}
