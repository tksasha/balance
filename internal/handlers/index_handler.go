package handlers

import (
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/components"
	"github.com/tksasha/balance/internal/server/app"
)

type IndexHandler struct{}

func NewIndexHandler(app *app.App) http.Handler {
	return &IndexHandler{}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := components.IndexPage().Render(r.Context(), w); err != nil {
		slog.Error(err.Error())
	}
}
