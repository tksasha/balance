package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/balance/internal/backoffice/category/component"
	"github.com/tksasha/balance/internal/common/handler"
	"github.com/tksasha/balance/internal/common/paths/params"
)

type EditHandler struct {
	*handler.Handler

	categoryService category.Service
	component       *component.Component
}

func NewEditHandler(
	categoryService category.Service,
) *EditHandler {
	return &EditHandler{
		Handler:         handler.New(),
		categoryService: categoryService,
		component:       component.New(),
	}
}

func (h *EditHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	category, err := h.handle(r)
	if err != nil {
		h.SetError(w, err)

		return
	}

	h.ok(w, category)
}

func (h *EditHandler) handle(r *http.Request) (*category.Category, error) {
	return h.categoryService.Edit(r.Context(), r.PathValue("id"))
}

func (h *EditHandler) ok(w http.ResponseWriter, category *category.Category) {
	params := params.New().WithCurrency(category.Currency)

	err := h.component.Edit(params, category, nil).Render(w)

	h.SetError(w, err)
}
