package handlers

import (
	"log/slog"
	"net/http"

	itemcomponents "github.com/tksasha/balance/internal/components/items"
)

type UpdateItemHandler struct{}

func NewUpdateItemHandler() http.Handler {
	return &UpdateItemHandler{}
}

func (h *UpdateItemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := itemcomponents.UpdatePage().Render(r.Context(), w); err != nil {
		slog.Error(err.Error())
	}
}
