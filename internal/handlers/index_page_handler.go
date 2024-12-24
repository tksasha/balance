package handlers

import (
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/components"
)

type indexPageHandler struct {
	categoryService CategoryService
}

func NewIndexPageHandler(categoryService CategoryService) Handler {
	return &indexPageHandler{
		categoryService: categoryService,
	}
}

func (h *indexPageHandler) Pattern() string {
	return "GET /"
}

func (h *indexPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(w, r); err != nil {
		slog.Error("index page handler error", "error", err)

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *indexPageHandler) handle(w http.ResponseWriter, r *http.Request) error {
	categories, err := h.categoryService.GetCategories(r.Context())
	if err != nil {
		return err
	}

	return components.IndexPage(categories).Render(w)
}
