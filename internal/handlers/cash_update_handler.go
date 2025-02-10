package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/requests"
	"github.com/tksasha/balance/internal/responses"
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
		if response, ok := w.(*responses.Response); ok {
			response.Error = err

			return
		}

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
