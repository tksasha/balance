package handlers

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	internalerrors "github.com/tksasha/balance/internal/errors"
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

func (h *CreateItemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(r); err != nil {
		if errors.Is(err, internalerrors.ErrParsingForm) {
			slog.Error("invalid user input", "error", err)

			http.Error(w, "Invalid User Input", http.StatusBadRequest)

			return
		}

		if errors.Is(err, internalerrors.ErrResourceInvalid) {
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

	item := &models.Item{
		Formula:     r.FormValue("formula"),
		Description: r.FormValue("description"),
	}

	if r.FormValue("date") != "" {
		date, err := time.Parse(time.DateOnly, r.FormValue("date")) // TODO: test me
		if err != nil {
			return internalerrors.ErrParsingForm
		}

		item.Date = date
	}

	if r.FormValue("category_id") != "" {
		categoryID, err := strconv.Atoi(r.FormValue("category_id")) // TODO: test me
		if err != nil {
			return internalerrors.ErrParsingForm
		}

		item.CategoryID = categoryID
	}

	if err := h.itemService.CreateItem(r.Context(), item); err != nil {
		return err
	}

	return nil
}
