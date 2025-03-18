package handlers

import (
	"net/http"
	"net/url"

	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/balance/internal/backoffice/category/component"
	"github.com/tksasha/balance/internal/common/component/path"
	"github.com/tksasha/balance/internal/common/handler"
)

type IndexHandler struct {
	*handler.Handler

	service   category.Service
	component *component.Component
}

func NewIndexHandler(
	service category.Service,
) *IndexHandler {
	return &IndexHandler{
		Handler:   handler.New(),
		service:   service,
		component: component.New(),
	}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	categories, err := h.handle(r)
	if err != nil {
		h.SetError(w, err)

		return
	}

	h.ok(w, r.URL.Query(), categories)
}

func (h *IndexHandler) handle(r *http.Request) (category.Categories, error) {
	return h.service.List(r.Context())
}

func (h *IndexHandler) ok(w http.ResponseWriter, values url.Values, categories category.Categories) {
	params := path.Params{
		"currency": values.Get("currency"),
	}

	w.Header().Add("Hx-Trigger-After-Swap", "backoffice.categories.shown")

	err := h.component.Index(params, categories).Render(w)

	h.SetError(w, err)
}
