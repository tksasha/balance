package handler

import (
	"net/http"

	"github.com/tksasha/balance/internal/app/index"
	"github.com/tksasha/balance/internal/app/index/component"
	"github.com/tksasha/balance/internal/common/handler"
)

type Handler struct {
	*handler.Handler

	indexService    index.Service
	cashService     index.CashService
	categoryService index.CategoryService
	component       *component.Component
}

func New(
	indexService index.Service,
	cashService index.CashService,
	categoryService index.CategoryService,
) *Handler {
	return &Handler{
		Handler:         handler.New(),
		indexService:    indexService,
		cashService:     cashService,
		categoryService: categoryService,
		component:       component.New(),
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	err = h.component.Index(cashes, categories, r.URL.Query()).Render(w)

	h.SetError(w, err)
}
