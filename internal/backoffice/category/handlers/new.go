package handlers

import (
	"net/http"
	"net/url"

	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/balance/internal/backoffice/category/component"
	"github.com/tksasha/balance/internal/common/component/path"
	"github.com/tksasha/balance/internal/common/currency"
	commonhandler "github.com/tksasha/balance/internal/common/handler"
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

	h.ok(w, r.URL.Query(), category)
}

func (h *NewHandler) ok(w http.ResponseWriter, values url.Values, category *category.Category) {
	params := path.Params{
		"currency": values.Get("currency"),
	}

	err := h.component.New(params, category, nil).Render(w)

	h.SetError(w, err)
}
