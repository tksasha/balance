package handlers

import (
	"errors"
	"net/http"

	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/balance/internal/app/cash/components"
	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/validator"
)

type CreateHandler struct {
	*common.BaseHandler

	cashService   cash.Service
	cashComponent *components.CashComponent
}

func NewCreateHandler(
	baseHandler *common.BaseHandler,
	cashService cash.Service,
	cashComponent *components.CashComponent,
) *CreateHandler {
	return &CreateHandler{
		BaseHandler:   baseHandler,
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

	var verrors validator.Errors
	if errors.As(err, &verrors) {
		err := h.cashComponent.Create(cash, verrors).Render(w)

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
		Favorite:      r.FormValue("favorite"),
	}

	return h.cashService.Create(r.Context(), request)
}
