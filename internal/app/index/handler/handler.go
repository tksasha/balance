package handler

import (
	"net/http"

	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/app/index"
	"github.com/tksasha/balance/internal/app/index/components"
	"github.com/tksasha/balance/internal/common"
)

type Handler struct {
	*common.BaseHandler

	indexService    index.Service
	categoryService index.CategoryService
	indexComponent  *components.IndexComponent
}

func New(
	baseHandler *common.BaseHandler,
	indexService index.Service,
	categoryService index.CategoryService,
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
