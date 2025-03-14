package handlers

import (
	"net/http"
	"net/url"

	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/backoffice/cash/component"
	"github.com/tksasha/balance/internal/common/component/path"
	"github.com/tksasha/balance/internal/common/handler"
)

type IndexHandler struct {
	*handler.Handler

	cashService cash.Service
	component   *component.Component
}

func NewIndexHandler(
	cashService cash.Service,
) *IndexHandler {
	return &IndexHandler{
		Handler:     handler.New(),
		cashService: cashService,
		component:   component.New(),
	}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cashes, err := h.handle(r)
	if err != nil {
		h.SetError(w, err)

		return
	}

	h.ok(w, r.URL.Query(), cashes)
}

func (h *IndexHandler) handle(r *http.Request) (cash.Cashes, error) {
	return h.cashService.List(r.Context())
}

func (h *IndexHandler) ok(w http.ResponseWriter, values url.Values, cashes cash.Cashes) {
	w.Header().Add("Hx-Trigger-After-Swap", "backoffice.cashes.shown")

	params := path.Params{
		"currency": values.Get("currency"),
	}

	err := h.component.Index(params, cashes).Render(w)

	h.SetError(w, err)
}
