package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	internalerrors "github.com/tksasha/balance/internal/errors"
	"github.com/tksasha/balance/internal/requests"
	"github.com/tksasha/balance/pkg/validationerror"
)

type CreateItemHandler struct {
	itemCreator ItemCreator
}

func NewCreateItemHandler(itemCreator ItemCreator) *CreateItemHandler {
	return &CreateItemHandler{
		itemCreator: itemCreator,
	}
}

func (h *CreateItemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(r); err != nil {
		if errors.Is(err, internalerrors.ErrParsingForm) {
			slog.Error("invalid user input", "error", err)

			http.Error(w, "Invalid User Input", http.StatusBadRequest)

			return
		}

		var validationErrors validationerror.ValidationError

		if errors.As(err, &validationErrors) {
			_, _ = w.Write([]byte("render form with errors"))

			return
		}

		slog.Error("create item handler error", "error", err)

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

		return
	}

	_, _ = w.Write([]byte("render create page\n"))
}

func (h *CreateItemHandler) handle(r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return internalerrors.ErrParsingForm
	}

	request := requests.CreateItemRequest{
		Date:        r.FormValue("date"),
		Formula:     r.FormValue("formula"),
		CategoryID:  r.FormValue("category_id"),
		Description: r.FormValue("description"),
	}

	if _, err := h.itemCreator.Create(r.Context(), request); err != nil {
		return err
	}

	return nil
}
