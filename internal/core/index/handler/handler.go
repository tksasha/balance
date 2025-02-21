package handler

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/common"
	"github.com/tksasha/balance/internal/core/index"
	"github.com/tksasha/balance/internal/core/index/components"
)

type Handler struct {
	*common.BaseHandler

	indexService    index.Service
	categoryService category.Service
	indexComponent  *components.IndexComponent
}

func New(
	baseHandler *common.BaseHandler,
	indexService index.Service,
	categoryService category.Service,
	indexComponent *components.IndexComponent,
) *Handler {
	return &Handler{
		BaseHandler:     baseHandler,
		indexService:    indexService,
		categoryService: categoryService,
		indexComponent:  indexComponent,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	categories, err := h.handle(r)
	if err != nil {
		h.SetError(w, err)

		return
	}

	err = h.indexComponent.Index(categories, r.URL.Query()).Render(w)

	h.SetError(w, err)
}

func (h *Handler) handle(r *http.Request) (category.Categories, error) {
	categories, err := h.categoryService.List(r.Context())
	if err != nil {
		return nil, err
	}

	return categories, nil
}
