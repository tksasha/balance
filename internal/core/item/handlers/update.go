package handlers

import (
	"errors"
	"net/http"

	"github.com/tksasha/balance/internal/core/category"
	apperrors "github.com/tksasha/balance/internal/core/common"
	"github.com/tksasha/balance/internal/core/common/handlers"
	"github.com/tksasha/balance/internal/core/item"
	"github.com/tksasha/balance/internal/core/item/components"
	"github.com/tksasha/balance/pkg/validation"
)

type UpdateHandler struct {
	service         item.Service
	categoryService category.Service
}

func NewUpdateHandler(service item.Service, categoryService category.Service) *UpdateHandler {
	return &UpdateHandler{
		service:         service,
		categoryService: categoryService,
	}
}

func (h *UpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	item, err := h.handle(r)
	if err == nil {
		w.WriteHeader(http.StatusNoContent)

		return
	}

	var verrors validation.Errors
	if errors.As(err, &verrors) {
		categories, err := h.categoryService.List(r.Context())
		if err != nil {
			handlers.SetError(w, err)

			return
		}

		err = components.Update(item, categories, verrors).Render(w)

		handlers.SetError(w, err)
	}

	handlers.SetError(w, err)
}

func (h *UpdateHandler) handle(r *http.Request) (*item.Item, error) {
	if err := r.ParseForm(); err != nil {
		return nil, apperrors.ErrParsingForm
	}

	request := item.UpdateRequest{
		ID:          r.PathValue("id"),
		Date:        r.FormValue("date"),
		Formula:     r.FormValue("formula"),
		CategoryID:  r.FormValue("category_id"),
		Description: r.FormValue("description"),
	}

	return h.service.Update(r.Context(), request)
}
