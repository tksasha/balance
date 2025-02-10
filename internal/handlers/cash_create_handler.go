package handlers

import (
	"errors"
	"net/http"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/requests"
	"github.com/tksasha/balance/pkg/validation"
)

type CashCreateHandler struct {
	cashService CashService
}

func NewCashCreateHandler(cashService CashService) *CashCreateHandler {
	return &CashCreateHandler{
		cashService: cashService,
	}
}

func (h *CashCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := h.handle(r); err != nil {
		if errors.Is(err, apperrors.ErrParsingForm) {
			http.Error(w, "Bad Request", http.StatusBadRequest)

			return
		}

		var verrors validation.Errors
		if errors.As(err, &verrors) {
			_, _ = w.Write([]byte(verrors.Error()))

			return
		}
	}

	_, _ = w.Write([]byte("cash"))
}

func (h *CashCreateHandler) handle(r *http.Request) (*models.Cash, error) {
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
