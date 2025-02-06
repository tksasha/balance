package cash

import (
	"errors"
	"net/http"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/models"
)

type EditHandler struct {
	cashService handlers.CashService
}

func NewEditHandler(cashService handlers.CashService) *EditHandler {
	return &EditHandler{
		cashService: cashService,
	}
}

func (h *EditHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cash, err := h.handle(r)
	if err != nil {
		if errors.Is(err, apperrors.ErrResourceNotFound) {
			http.Error(w, "Resource Not Found", http.StatusNotFound)

			return
		}
	}

	_, _ = w.Write([]byte(cash.Name))
}

func (h *EditHandler) handle(r *http.Request) (*models.Cash, error) {
	return h.cashService.FindByID(r.Context(), r.PathValue("id"))
}
