package handlers

import (
	"errors"
	"net/http"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/requests"
	"github.com/tksasha/balance/internal/responses"
	"github.com/tksasha/balance/pkg/validation"
)

type ItemCreateHandler struct {
	itemService ItemService
}

func NewItemCreateHandler(itemService ItemService) *ItemCreateHandler {
	return &ItemCreateHandler{
		itemService: itemService,
	}
}

func (h *ItemCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(r); err != nil {
		var verrors validation.Errors
		if errors.As(err, &verrors) {
			_, _ = w.Write([]byte(verrors.Error()))

			return
		}

		if response, ok := w.(*responses.Response); ok {
			response.Error = err

			return
		}

		return
	}

	_, _ = w.Write([]byte("render create page\n"))
}

func (h *ItemCreateHandler) handle(r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return apperrors.ErrParsingForm
	}

	if _, err := h.itemService.Create(
		r.Context(),
		requests.ItemCreateRequest{
			Date:        r.FormValue("date"),
			Formula:     r.FormValue("formula"),
			CategoryID:  r.FormValue("category_id"),
			Description: r.FormValue("description"),
		},
	); err != nil {
		return err
	}

	return nil
}
