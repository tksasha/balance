package handlers

import (
	"errors"
	"net/http"

	"github.com/tksasha/balance/internal/core/cash"
	"github.com/tksasha/balance/internal/core/cash/components"
	"github.com/tksasha/balance/internal/core/common"
	"github.com/tksasha/balance/internal/core/common/handlers"
	"github.com/tksasha/balance/pkg/validation"
)

type CreateHandler struct {
	cashService   cash.Service
	cashComponent *components.CashComponent
}

func NewCreateHandler(
	cashService cash.Service,
	cashComponent *components.CashComponent,
) *CreateHandler {
	return &CreateHandler{
		cashService:   cashService,
		cashComponent: cashComponent,
	}
}

func (h *CreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cash, err := h.handle(r)
	if err == nil {
		w.WriteHeader(http.StatusCreated)

		return
	}

	var verrors validation.Errors
	if errors.As(err, &verrors) {
		err := h.cashComponent.Create(cash, verrors).Render(w)

		handlers.SetError(w, err)

		return
	}

	handlers.SetError(w, err)
}

func (h *CreateHandler) handle(r *http.Request) (*cash.Cash, error) {
	if err := r.ParseForm(); err != nil {
		return nil, common.ErrParsingForm
	}

	request := cash.CreateRequest{
		Name:          r.FormValue("name"),
		Formula:       r.FormValue("formula"),
		Supercategory: r.FormValue("supercategory"),
		Favorite:      r.FormValue("favorite"),
	}

	return h.cashService.Create(r.Context(), request)
}
