package handlers

import (
	"errors"
	"net/http"

	"github.com/tksasha/balance/internal/core/category"
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
	if err := h.handle(r); err != nil {
		var verrors validation.Errors
		if errors.As(err, &verrors) {
			_, _ = w.Write([]byte(verrors.Error()))

			return
		}

		handlers.E(w, err)

		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *CreateHandler) handle(r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return common.ErrParsingForm
	}

	request := category.CreateRequest{
		Name:          r.FormValue("name"),
		Income:        r.FormValue("income"),
		Visible:       r.FormValue("visible"),
		Supercategory: r.FormValue("supercategory"),
	}

	return h.service.Create(r.Context(), request)
}
