package handlers

import (
	"errors"
	"net/http"

	"github.com/tksasha/balance/internal/common/handlers"
	"github.com/tksasha/balance/internal/core/cash"
	"github.com/tksasha/balance/internal/core/common"
	"github.com/tksasha/balance/pkg/validation"
)

type CreateHandler struct {
	service cash.Service
}

func NewCreateHandler(service cash.Service) *CreateHandler {
	return &CreateHandler{
		service: service,
	}
}

func (h *CreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := h.handle(r); err != nil {
		var verrors validation.Errors
		if errors.As(err, &verrors) {
			_, _ = w.Write([]byte(verrors.Error()))

			return
		}

		handlers.E(w, err)

		return
	}

	_, _ = w.Write([]byte("cash"))
}

func (h *CreateHandler) handle(r *http.Request) (*cash.Cash, error) {
	if err := r.ParseForm(); err != nil {
		return nil, common.ErrParsingForm
	}

	request := cash.CreateRequest{
		Name:          r.FormValue("name"),
		Formula:       r.FormValue("formula"),
		Supercategory: r.FormValue("supercategory"),
		Favorite:      r.FormValue("favorite"),
	}

	if err := h.service.Create(r.Context(), request); err != nil {
		return nil, err
	}

	return &cash.Cash{}, nil
}
