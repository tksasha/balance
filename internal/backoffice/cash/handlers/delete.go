package handlers

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/common/component/path"
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/common/handler"
)

type DeleteHandler struct {
	*handler.Handler

	cashService cash.Service
}

func NewDeleteHandler(cashService cash.Service) *DeleteHandler {
	return &DeleteHandler{
		Handler:     handler.New(),
		cashService: cashService,
	}
}

func (h *DeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(r); err != nil {
		h.SetError(w, err)

		return
	}

	h.ok(w, currency.Default)
}

func (h *DeleteHandler) handle(r *http.Request) error {
	return h.cashService.Delete(r.Context(), r.PathValue("id"))
}

func (h *DeleteHandler) ok(w http.ResponseWriter, currency currency.Currency) {
	writer := bytes.NewBuffer([]byte{})

	header := map[string]map[string]string{
		"backoffice.cash.deleted": {
			"backofficeCashesPath": path.BackofficeCashes(path.NewCurrency(currency)),
		},
	}

	if err := json.NewEncoder(writer).Encode(header); err != nil {
		slog.Error("failed to encode", "error", err)

		writer.Reset()
	}

	w.Header().Add("Hx-Trigger-After-Swap", writer.String())

	w.WriteHeader(http.StatusOK)
}
