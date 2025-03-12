package handlers

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"

	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/common/component/path"
	"github.com/tksasha/balance/internal/common/handler"
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

	h.StatusOK(w, r.URL.Query(), item)
}

func (h *DeleteHandler) StatusOK(w http.ResponseWriter, values url.Values, item *item.Item) {
	params := path.Params{
		"month": strconv.Itoa(int(item.Date.Month())),
		"year":  strconv.Itoa(item.Date.Year()),
	}

	header := map[string]map[string]string{
		"balance.item.deleted": {
			"itemsPath":      path.Items(values, params),
			"balancePath":    path.Balance(values),
			"categoriesPath": path.Categories(values, params),
		},
	}

	writer := bytes.NewBuffer([]byte{})

	if err := json.NewEncoder(writer).Encode(header); err != nil {
		slog.Error("failed to encode", "error", err)
	}

	w.Header().Add("Hx-Trigger-After-Swap", writer.String())

	w.WriteHeader(http.StatusOK)
}
