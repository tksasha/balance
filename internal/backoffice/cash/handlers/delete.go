package handlers

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/common/handler"
	"github.com/tksasha/balance/internal/common/paths"
	"github.com/tksasha/balance/internal/common/paths/params"
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

	params := params.New().WithCurrency(currency.Default)

	h.ok(w, params)
}

func (h *DeleteHandler) handle(r *http.Request) error {
	return h.cashService.Delete(r.Context(), r.PathValue("id"))
}

func (h *DeleteHandler) ok(w http.ResponseWriter, params params.Params) {
	writer := bytes.NewBuffer([]byte{})

	header := map[string]map[string]string{
		"backoffice.cash.deleted": {
			"backofficeCashesPath": paths.BackofficeCashes(params),
		},
	}

	if err := json.NewEncoder(writer).Encode(header); err != nil {
		slog.Error("failed to encode", "error", err)

		writer.Reset()
	}

	w.Header().Add("Hx-Trigger-After-Swap", writer.String())

	w.WriteHeader(http.StatusOK)
}
