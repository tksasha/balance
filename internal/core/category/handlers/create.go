package handlers

import (
	"errors"
	"net/http"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/category/components"
	"github.com/tksasha/balance/internal/core/common"
	"github.com/tksasha/balance/internal/core/common/handlers"
	"github.com/tksasha/balance/pkg/validation"
)

type CreateHandler struct {
	service category.Service
}

func NewCreateHandler(service category.Service) *CreateHandler {
	return &CreateHandler{
		service: service,
	}
}

func (h *CreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	category, err := h.handle(r)
	if err == nil {
		w.WriteHeader(http.StatusCreated)

		return
	}

	var verrors validation.Errors
	if errors.As(err, &verrors) {
		err := components.Create(category, verrors).Render(w)

		handlers.SetError(w, err)

		return
	}

	handlers.SetError(w, err)
}

func (h *CreateHandler) handle(r *http.Request) (*category.Category, error) {
	if err := r.ParseForm(); err != nil {
		return nil, common.ErrParsingForm
	}

	request := category.CreateRequest{
		Name:          r.FormValue("name"),
		Income:        r.FormValue("income"),
		Visible:       r.FormValue("visible"),
		Supercategory: r.FormValue("supercategory"),
	}

	return h.service.Create(r.Context(), request)
}
