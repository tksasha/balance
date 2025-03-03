package handler

import (
	"net/http"

	"github.com/tksasha/balance/internal/app/index"
	"github.com/tksasha/balance/internal/app/index/component"
	"github.com/tksasha/balance/internal/common/handler"
)

type Handler struct {
	*handler.Handler

	balanceService  index.BalanceService
	cashService     index.CashService
	categoryService index.CategoryService
	component       *component.Component
}

func New(
	balanceService index.BalanceService,
	cashService index.CashService,
	categoryService index.CategoryService,
) *Handler {
	return &Handler{
		Handler:         handler.New(),
		balanceService:  balanceService,
		cashService:     cashService,
		categoryService: categoryService,
		component:       component.New(),
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	balance, err := h.balanceService.Balance(r.Context())
	if err != nil {
		h.SetError(w, err)

		return
	}

	cashes, err := h.cashService.List(r.Context())
	if err != nil {
		h.SetError(w, err)

		return
	}

	categories, err := h.categoryService.List(r.Context())
	if err != nil {
		h.SetError(w, err)

		return
	}

	err = h.component.Index(balance, cashes, categories, r.URL.Query()).Render(w)

	h.SetError(w, err)
}
