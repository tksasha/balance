package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/components"
	"github.com/tksasha/balance/internal/models"
)

type CreateItemHandler struct {
	itemService     ItemService
	categoryService CategoryService
}

func NewCreateItemHandler(itemService ItemService, categoryService CategoryService) *CreateItemHandler {
	return &CreateItemHandler{
		itemService:     itemService,
		categoryService: categoryService,
	}
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

	item, validationErrors := models.
		NewItemBuilder().
		WithDate(r.FormValue("date")).
		Build()

	categories, err := h.categoryService.GetCategories(r.Context())
	if err != nil {
		return err
	}

	if !validationErrors.IsEmpty() {
		if err := components.ItemForm(item, categories, validationErrors).Render(w); err != nil {
			return nil //nolint:nilerr // TODO: render form with errors here
		}

		return nil // TODO: return there NewValidationError()
	}

	if err := h.itemService.CreateItem(r.Context(), item); err != nil {
		return err
	}

	return nil
}
