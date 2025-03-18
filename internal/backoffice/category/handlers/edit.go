package handlers

import (
	"net/http"
	"net/url"

	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/balance/internal/backoffice/category/component"
	"github.com/tksasha/balance/internal/common/component/path"
	"github.com/tksasha/balance/internal/common/handler"
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

	h.ok(w, r.URL.Query(), category)
}

func (h *EditHandler) handle(r *http.Request) (*category.Category, error) {
	return h.categoryService.Edit(r.Context(), r.PathValue("id"))
}

func (h *EditHandler) ok(w http.ResponseWriter, values url.Values, category *category.Category) {
	params := path.Params{
		"currency": values.Get("currency"),
	}

	err := h.component.Edit(params, category).Render(w)

	h.SetError(w, err)
}
