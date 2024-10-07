package handlers

import (
	"log/slog"
	"net/http"

	itemcomponents "github.com/tksasha/balance/internal/components/items"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/repositories"
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
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	item := models.NewItem().
		SetDate(r.FormValue("date")).
		SetFormula(r.FormValue("formula")).
		SetCategoryID(r.FormValue("category_id")).
		SetDescription(r.FormValue("description"))

	if !item.IsValid() {
		if err := itemcomponents.Form(item).Render(r.Context(), w); err != nil {
			slog.Error(err.Error())
		}

		return
	}

	if err := services.CreateItemService(r.Context(), h.itemRepository, item); err != nil {
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

	if err := itemcomponents.CreatePage(items).Render(r.Context(), w); err != nil {
		slog.Error(err.Error())
	}
}
