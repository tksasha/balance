package cashes

import (
	"errors"
	"net/http"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/handlers/utils"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/requests"
	"github.com/tksasha/balance/pkg/validation"
)

type CreateHandler struct {
	cashService handlers.CashService
}

func NewCreateHandler(cashService handlers.CashService) *CreateHandler {
	return &CreateHandler{
		cashService: cashService,
	}
}

func (h *CreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := h.handle(r); err != nil {
		var verrors validation.Errors
		if errors.As(err, &verrors) {
			_, _ = w.Write([]byte(verrors.Error()))

			return
		}

		utils.E(w, err)

		return
	}

	_, _ = w.Write([]byte("cash"))
}

func (h *CreateHandler) handle(r *http.Request) (*models.Cash, error) {
	if err := r.ParseForm(); err != nil {
		return nil, apperrors.ErrParsingForm
	}

	request := requests.CashCreateRequest{
		Name:          r.FormValue("name"),
		Formula:       r.FormValue("formula"),
		Supercategory: r.FormValue("supercategory"),
		Favorite:      r.FormValue("favorite"),
	}

	if err := h.cashService.Create(r.Context(), request); err != nil {
		return nil, err
	}

	return &models.Cash{}, nil
}
