package handlers

import (
	"log/slog"
	"net/http"

	indexcomponents "github.com/tksasha/balance/internal/components/index"
	"github.com/tksasha/balance/internal/server/app"
)

type IndexHandler struct{}

func NewIndexHandler(app *app.App) http.Handler {
	return &IndexHandler{}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := indexcomponents.IndexPage().Render(r.Context(), w); err != nil {
		slog.Error(err.Error())
	}
}
