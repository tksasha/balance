package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/backoffice/cash/component"
	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/balance/internal/common/handler"
	"github.com/tksasha/balance/internal/common/paths"
	"github.com/tksasha/balance/internal/common/paths/params"
	"github.com/tksasha/validation"
)

type CreateHandler struct {
	*handler.Handler

	cashService cash.Service
	component   *component.Component
}

func NewCreateHandler(
	cashService cash.Service,
) *CreateHandler {
	return &CreateHandler{
		Handler:     handler.New(),
		cashService: cashService,
		component:   component.New(),
	}
}

func (h *CreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cash, err := h.handle(r)
	if err == nil {
		params := params.New().SetCurrency(cash.Currency)

		h.ok(w, params)

		return
	}

	var verrors validation.Errors
	if errors.As(err, &verrors) {
		err := h.component.Create(cash, verrors).Render(w)

		h.SetError(w, err)

		return
	}

	h.SetError(w, err)
}

func (h *CreateHandler) handle(r *http.Request) (*cash.Cash, error) {
	if err := r.ParseForm(); err != nil {
		return nil, common.ErrParsingForm
	}

	request := cash.CreateRequest{
		Name:          r.FormValue("name"),
		Formula:       r.FormValue("formula"),
		Supercategory: r.FormValue("supercategory"),
		Currency:      r.FormValue("currency"),
	}

	return h.cashService.Create(r.Context(), request)
}

func (h *CreateHandler) ok(w http.ResponseWriter, params params.Params) {
	writer := bytes.NewBuffer([]byte{})

	header := map[string]map[string]string{
		"backoffice.cash.created": {
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
