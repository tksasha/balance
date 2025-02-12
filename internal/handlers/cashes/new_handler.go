package cashes

import (
	"net/http"

	components "github.com/tksasha/balance/internal/components/cash"
	"github.com/tksasha/balance/internal/handlers/utils"
	"github.com/tksasha/balance/internal/models"
)

type NewHandler struct{}

func NewNewHandler() *NewHandler {
	return &NewHandler{}
}

func (h *NewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cash := &models.Cash{}

	if err := components.CashNew(cash).Render(w); err != nil {
		utils.E(w, err)
	}
}
