package handlers

import (
	"errors"
	"net/http"

	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/balance/internal/app/cash/component"
	"github.com/tksasha/balance/internal/common"
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
		err := h.component.Update(cash).Render(w)

		h.SetError(w, err)

		return
	}

	var verrors validation.Errors
	if errors.As(err, &verrors) {
		err := h.component.Edit(cash, verrors).Render(w)

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
