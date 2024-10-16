package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/repositories"
)

type CreateItemHandler struct {
	itemCreator repositories.ItemCreator
}

func NewCreateItemHandler(itemCreator repositories.ItemCreator) http.Handler {
	return &CreateItemHandler{
		itemCreator: itemCreator,
	}
}

func (h *CreateItemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	item, validationErrors := models.NewItemBuilder().Build()

	if !validationErrors.IsEmpty() {
		_, _ = w.Write([]byte("form with validation errors\n"))
	}

	if err := h.itemCreator.CreateItem(r.Context(), item); err != nil {
		// render internal server error
		_, _ = w.Write([]byte("internal server error\n"))
	}

	// render create page

	_, _ = w.Write([]byte("create page\n"))
}
