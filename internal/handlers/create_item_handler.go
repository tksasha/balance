package handlers

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/tksasha/balance/internal/components"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/requests"
	"github.com/tksasha/balance/internal/server/app"
	"github.com/tksasha/balance/internal/services"
)

type CreateItemHandler struct {
	itemRepository *repositories.ItemRepository
}

func NewCreateItemHandler(app *app.App) http.Handler {
	return &CreateItemHandler{
		itemRepository: repositories.NewItemRepository(app.DB),
	}
}

func (h *CreateItemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second)

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	createItemRequest := requests.NewCreateItemRequest()

	createItemRequest.Parse(
		r.FormValue("date"),
		r.FormValue("formula"),
		r.FormValue("category_id"),
		r.FormValue("description"),
	)

	if !createItemRequest.Valid() {
		if err := components.ItemForm(createItemRequest).Render(r.Context(), w); err != nil {
			slog.Error(err.Error())
		}

		return
	}

	if err := services.CreateItemService(
		r.Context(),
		h.itemRepository,
		createItemRequest.ToItem(),
	); err != nil {
		slog.Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	items, err := h.itemRepository.GetItems(r.Context())
	if err != nil {
		slog.Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := components.ItemCreated(items).Render(r.Context(), w); err != nil {
		slog.Error(err.Error())
	}
}
