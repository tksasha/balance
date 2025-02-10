package handlers

import (
	"errors"
	"net/http"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/requests"
)

type CashUpdateHandler struct {
	cashService CashService
}

func NewCashUpdateHandler(cashService CashService) *CashUpdateHandler {
	return &CashUpdateHandler{
		cashService: cashService,
	}
}

func (h *CashUpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cash, err := h.handle(r)
	if err != nil {
		if errors.Is(err, apperrors.ErrParsingForm) {
			http.Error(w, "Bad Request", http.StatusBadRequest)

			return
		}

		if errors.Is(err, apperrors.ErrResourceNotFound) {
			http.Error(w, "Resource Not Found", http.StatusNotFound)

			return
		}

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

		return
	}

	_, _ = w.Write([]byte(cash.Name))
}

func (h *CashUpdateHandler) handle(r *http.Request) (*models.Cash, error) {
	if err := r.ParseForm(); err != nil {
		return nil, apperrors.ErrParsingForm
	}

	request := requests.CashUpdateRequest{
		ID:            r.PathValue("id"),
		Formula:       r.FormValue("formula"),
		Name:          r.FormValue("name"),
		Supercategory: r.FormValue("supercategory"),
		Favorite:      r.FormValue("favorite"),
	}

	return h.cashService.Update(r.Context(), request)
}
