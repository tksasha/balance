package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/balance/internal/app/cash/component"
	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/balance/internal/common/component/path"
	"github.com/tksasha/balance/internal/common/handler"
	"github.com/tksasha/validation"
)

type UpdateHandler struct {
	*handler.Handler

	cashService cash.Service
	component   *component.Component
}

func NewUpdateHandler(
	cashService cash.Service,
) *UpdateHandler {
	return &UpdateHandler{
		Handler:     handler.New(),
		cashService: cashService,
		component:   component.New(),
	}
}

func (h *UpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cash, err := h.handle(r)
	if err == nil {
		h.ok(w, r.URL.Query(), cash)

		return
	}

	var verrors validation.Errors
	if errors.As(err, &verrors) {
		err := h.component.Edit(r.URL.Query(), cash, verrors).Render(w)

		h.SetError(w, err)

		return
	}

	h.SetError(w, err)
}

func (h *UpdateHandler) handle(r *http.Request) (*cash.Cash, error) {
	if err := r.ParseForm(); err != nil {
		return nil, common.ErrParsingForm
	}

	request := cash.UpdateRequest{
		ID:      r.PathValue("id"),
		Formula: r.FormValue("formula"),
		Name:    r.FormValue("name"),
	}

	return h.cashService.Update(r.Context(), request)
}

func (h *UpdateHandler) ok(w http.ResponseWriter, values url.Values, cash *cash.Cash) {
	writer := bytes.NewBuffer([]byte{})

	header := map[string]map[string]string{
		"balance.cash.updated": {
			"balancePath": path.Balance(values),
		},
	}

	if err := json.NewEncoder(writer).Encode(header); err != nil {
		slog.Error("failed to encode", "error", err)
	}

	w.Header().Add("Hx-Trigger-After-Swap", writer.String())

	err := h.component.Update(values, cash).Render(w)

	h.SetError(w, err)
}
