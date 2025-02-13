package handlers

import (
	"errors"
	"net/http"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/category"
	"github.com/tksasha/balance/internal/category/components"
	"github.com/tksasha/balance/internal/common/handlers"
	"github.com/tksasha/balance/pkg/validation"
)

type UpdateHandler struct {
	service category.Service
}

func NewUpdateHandler(service category.Service) *UpdateHandler {
	return &UpdateHandler{
		service: service,
	}
}

func (h *UpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	category, err := h.handle(r)
	if err != nil {
		var verrors validation.Errors
		if errors.As(err, &verrors) {
			_, _ = w.Write([]byte(verrors.Error()))

			return
		}

		handlers.E(w, err)

		return
	}

	if err := components.Update(category).Render(w); err != nil {
		handlers.E(w, err)
	}
}

func (h *UpdateHandler) handle(r *http.Request) (*category.Category, error) {
	if err := r.ParseForm(); err != nil {
		return nil, apperrors.ErrParsingForm
	}

	request := category.UpdateRequest{
		ID:            r.PathValue("id"),
		Name:          r.FormValue("name"),
		Income:        r.FormValue("income"),
		Visible:       r.FormValue("visible"),
		Supercategory: r.FormValue("supercategory"),
	}

	return h.service.Update(r.Context(), request)
}
