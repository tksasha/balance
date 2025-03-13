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
		h.StatusOK(w)

		return
	}

	var verrors validation.Errors
	if errors.As(err, &verrors) {
		err := h.component.Update(cash, verrors).Render(w)

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
		ID:            r.PathValue("id"),
		Formula:       r.FormValue("formula"),
		Name:          r.FormValue("name"),
		Supercategory: r.FormValue("supercategory"),
		Favorite:      r.FormValue("favorite"),
	}

	return h.cashService.Update(r.Context(), request)
}

func (h *UpdateHandler) StatusOK(w http.ResponseWriter) {
	writer := bytes.NewBuffer([]byte{})

	values := map[string]map[string]string{
		"backoffice.cash.updated": {
			"backofficeCashesPath": path.BackofficeCashes(nil),
		},
	}

	if err := json.NewEncoder(writer).Encode(values); err != nil {
		slog.Error("failed to encode", "error", err)
	}

	w.Header().Set("Hx-Trigger-After-Swap", writer.String())

	w.WriteHeader(http.StatusOK)
}
