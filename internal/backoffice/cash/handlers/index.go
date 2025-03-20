package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/backoffice/cash/component"
	"github.com/tksasha/balance/internal/common/handler"
	"github.com/tksasha/balance/internal/common/paths/params"
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

	params := params.New(r.URL.Query())

	h.ok(w, params, cashes)
}

func (h *IndexHandler) handle(r *http.Request) (cash.Cashes, error) {
	return h.cashService.List(r.Context())
}

func (h *IndexHandler) ok(w http.ResponseWriter, params params.Params, cashes cash.Cashes) {
	w.Header().Add("Hx-Trigger-After-Swap", "backoffice.cashes.shown")

	err := h.component.Index(params, cashes).Render(w)

	h.SetError(w, err)
}
