package handlers

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/app/item/component"
	"github.com/tksasha/balance/internal/common/handler"
	"github.com/tksasha/balance/internal/common/paths/params"
)

type IndexHandler struct {
	*handler.Handler

	itemService item.Service
	component   *component.Component
}

func NewIndexHandler(
	itemService item.Service,
) *IndexHandler {
	return &IndexHandler{
		Handler:     handler.New(),
		itemService: itemService,
		component:   component.New(),
	}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	items, err := h.handle(r)
	if err != nil {
		h.SetError(w, err)

		return
	}

	params := params.New(r.URL.Query())

	h.ok(w, params, items)
}

func (h *IndexHandler) handle(r *http.Request) (item.Items, error) {
	request := item.ListRequest{
		Month:    r.URL.Query().Get("month"),
		Year:     r.URL.Query().Get("year"),
		Category: r.URL.Query().Get("category"),
	}

	return h.itemService.List(r.Context(), request)
}

func (h *IndexHandler) ok(w http.ResponseWriter, params params.Params, items item.Items) {
	month, year := params.Get("month"), params.Get("year")

	header := map[string]map[string]string{
		"balance.items.shown": {
			"month": month,
			"year":  year,
		},
	}

	writer := bytes.NewBuffer([]byte{})

	if err := json.NewEncoder(writer).Encode(header); err != nil {
		slog.Error("failed to encode", "error", err)

		writer.Reset()
	}

	w.Header().Add("Hx-Trigger-After-Swap", writer.String())

	err := h.component.Index(params, items).Render(w)

	h.SetError(w, err)
}
