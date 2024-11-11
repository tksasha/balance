package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/components"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/services"
)

type createItemHandler struct {
	itemService services.ItemService
}

func NewCreateItemHandler(itemService services.ItemService) Handler {
	return &createItemHandler{itemService}
}

func (h *createItemHandler) Pattern() string {
	return "POST /items"
}

func (h *createItemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(w, r); err != nil {
		var formParsingError *FormParsingError

		if errors.As(err, &formParsingError) {
			slog.Error("invalid user input", "error", err)

			http.Error(w, "Invalid User Input", http.StatusBadRequest)

			return
		}

		slog.Error("create item handler error", "error", err)

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

		return
	}

	_, _ = w.Write([]byte("render create page\n"))
}

func (h *createItemHandler) handle(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return NewFormParsingError(err)
	}

	item, validationErrors := models.NewItemBuilder().
		WithDate(r.FormValue("date")).
		Build()

	if !validationErrors.IsEmpty() {
		if err := components.ItemForm(item, validationErrors).Render(w); err != nil {
			return err
		}

		return nil // TODO: return there NewValidationError()
	}

	if err := h.itemService.CreateItem(r.Context(), item); err != nil {
		return err
	}

	return nil
}
