package handlers

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
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

	w.Header().Add("Hx-Trigger-After-Swap", h.header(int(item.Date.Month()), item.Date.Year()))

	w.WriteHeader(http.StatusOK)
}

func (h *DeleteHandler) header(month, year int) string {
	params := path.Params{
		"month": strconv.Itoa(month),
		"year":  strconv.Itoa(year),
	}

	values := map[string]map[string]string{
		"balance.item.deleted": {
			"itemsPath":      path.Items(params, nil),
			"balancePath":    path.Balance(),
			"categoriesPath": path.Categories(params),
		},
	}

	w := bytes.NewBuffer([]byte{})

	if err := json.NewEncoder(w).Encode(values); err != nil {
		slog.Error("failed to encode", "error", err)
	}

	return w.String()
}
