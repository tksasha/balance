package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/app/item/component"
	commonhandler "github.com/tksasha/balance/internal/common/handler"
	"github.com/tksasha/balance/internal/common/paths/params"
)

type NewHandler struct {
	*commonhandler.Handler

	categoryService item.CategoryService
	component       *component.Component
}

func NewNewHandler(categoryService item.CategoryService) *NewHandler {
	return &NewHandler{
		Handler:         commonhandler.New(),
		categoryService: categoryService,
		component:       component.New(),
	}
}

func (h *NewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	categories, err := h.categoryService.List(r.Context())
	if err != nil {
		h.SetError(w, err)

		return
	}

	w.Header().Set("Hx-Trigger-After-Swap", "balance.item.initialized")

	params := params.New(r.URL.Query())

	if err := h.component.New(params, categories).Render(w); err != nil {
		h.SetError(w, err)

		return
	}
}
