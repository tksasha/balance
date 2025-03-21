package handlers

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/common/handler"
	"github.com/tksasha/balance/internal/common/paths"
	"github.com/tksasha/balance/internal/common/paths/params"
)

type DeleteHandler struct {
	*handler.Handler

	itemService item.Service
}

func NewDeleteHandler(
	itemService item.Service,
) *DeleteHandler {
	return &DeleteHandler{
		Handler:     handler.New(),
		itemService: itemService,
	}
}

func (h *DeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	item, err := h.itemService.Delete(r.Context(), r.PathValue("id"))
	if err != nil {
		h.SetError(w, err)

		return
	}

	params := params.New(r.URL.Query())

	h.StatusOK(w, params, item)
}

func (h *DeleteHandler) StatusOK(w http.ResponseWriter, params params.Params, item *item.Item) {
	month, year := int(item.Date.Month()), item.Date.Year()

	params = params.WithMonth(month).WithYear(year)

	header := map[string]map[string]string{
		"balance.item.deleted": {
			"itemsPath":      paths.Items(params),
			"balancePath":    paths.Balance(params),
			"categoriesPath": paths.Categories(params),
		},
	}

	writer := bytes.NewBuffer([]byte{})

	if err := json.NewEncoder(writer).Encode(header); err != nil {
		slog.Error("failed to encode", "error", err)
	}

	w.Header().Add("Hx-Trigger-After-Swap", writer.String())

	w.WriteHeader(http.StatusOK)
}
