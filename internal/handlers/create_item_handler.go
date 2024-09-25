package handlers

import (
	"html/template"
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/requests"
	"github.com/tksasha/balance/internal/server/app"
)

type CreateItemHandler struct {
	template       *template.Template
	itemRepository *repositories.ItemRepository
}

func NewCreateItemHandler(app *app.App) http.Handler {
	return &CreateItemHandler{
		template:       app.T,
		itemRepository: repositories.NewItemRepository(app.DB),
	}
}

func (h *CreateItemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	createItemRequest := requests.NewCreateItemRequest()

	createItemRequest.Parse(
		r.FormValue("date"),
		r.FormValue("formula"),
		"",
		r.FormValue("description"),
	)

	if err := h.template.ExecuteTemplate(w, "inline-item-form", createItemRequest); err != nil {
		slog.Error(err.Error())
	}
}
