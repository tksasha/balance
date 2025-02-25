package handlers

import (
	"errors"
	"net/http"

	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/balance/internal/app/cash/component"
	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/balance/internal/common/handler"
	"github.com/tksasha/validator"
)

type UpdateHandler struct {
	*handler.Handler

	cashService   cash.Service
	cashComponent *component.CashComponent
}

func NewUpdateHandler(
	cashService cash.Service,
	cashComponent *component.CashComponent,
) *UpdateHandler {
	return &UpdateHandler{
		Handler:       handler.New(),
		cashService:   cashService,
		cashComponent: cashComponent,
	}
}

func (h *UpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cash, err := h.handle(r)
	if err == nil {
		w.WriteHeader(http.StatusNoContent)

		return
	}

	var verrors validator.Errors
	if errors.As(err, &verrors) {
		err := h.cashComponent.Update(cash, verrors).Render(w)

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
