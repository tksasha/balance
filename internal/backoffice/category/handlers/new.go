package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/balance/internal/backoffice/category/component"
	"github.com/tksasha/balance/internal/common/currency"
	commonhandler "github.com/tksasha/balance/internal/common/handler"
	"github.com/tksasha/balance/internal/common/paths/params"
)

type NewHandler struct {
	*commonhandler.Handler

	component *component.Component
}

func NewNewHandler() *NewHandler {
	return &NewHandler{
		Handler:   commonhandler.New(),
		component: component.New(),
	}
}

func (h *NewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()

	category := &category.Category{
		Currency: currency.GetByCode(values.Get("currency")),
	}

	params := params.New(r.URL.Query())

	h.ok(w, params, category)
}

func (h *NewHandler) ok(w http.ResponseWriter, params params.Params, category *category.Category) {
	err := h.component.New(params, category, nil).Render(w)

	h.SetError(w, err)
}
