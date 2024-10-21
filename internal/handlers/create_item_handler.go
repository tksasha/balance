package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/components"
	"github.com/tksasha/balance/internal/interfaces"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/services"
)

type CreateItemHandler struct {
	service *services.ItemService
}

func NewCreateItemHandler(itemCreator repositories.ItemCreator) interfaces.Route {
	service := services.NewItemServiceBuilder().WithItemCreator(itemCreator).Build()

	return &CreateItemHandler{service: service}
}

func (h *CreateItemHandler) Pattern() string {
	return "POST /items"
}

func (h *CreateItemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func (h *CreateItemHandler) handle(w http.ResponseWriter, r *http.Request) error {
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

	if err := h.service.CreateItem(r.Context(), item); err != nil {
		return err
	}

	return nil
}
