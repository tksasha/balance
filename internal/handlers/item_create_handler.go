package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/requests"
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
		if errors.Is(err, apperrors.ErrParsingForm) {
			slog.Error("invalid user input", "error", err)

			http.Error(w, "Invalid User Input", http.StatusBadRequest)

			return
		}

		var verrors validation.Errors
		if errors.As(err, &verrors) {
			_, _ = w.Write([]byte(verrors.Error()))

			return
		}

		slog.Error("create item handler error", "error", err)

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

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
