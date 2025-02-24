package handlers

import (
	"errors"
	"net/http"

	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/balance/internal/core/cash"
	"github.com/tksasha/balance/internal/core/cash/components"
	"github.com/tksasha/balance/pkg/validation"
)

type UpdateHandler struct {
	*common.BaseHandler

	cashService   cash.Service
	cashComponent *components.CashComponent
}

func NewUpdateHandler(
	baseHandler *common.BaseHandler,
	cashService cash.Service,
	cashComponent *components.CashComponent,
) *UpdateHandler {
	return &UpdateHandler{
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

	var verrors validation.Errors
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
		ID:            r.PathValue("id"),
		Formula:       r.FormValue("formula"),
		Name:          r.FormValue("name"),
		Supercategory: r.FormValue("supercategory"),
		Favorite:      r.FormValue("favorite"),
	}

	return h.cashService.Update(r.Context(), request)
}
