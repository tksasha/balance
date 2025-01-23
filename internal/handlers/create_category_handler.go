package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	internalerrors "github.com/tksasha/balance/internal/errors"
	"github.com/tksasha/balance/internal/requests"
	"github.com/tksasha/balance/pkg/validationerror"
)

type CreateCategoryHandler struct {
	categoryCreator CategoryCreator
}

func NewCreateCategoryHandler(categoryCreator CategoryCreator) *CreateCategoryHandler {
	return &CreateCategoryHandler{
		categoryCreator: categoryCreator,
	}
}

func (h *CreateCategoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(r); err != nil {
		var validationError validationerror.ValidationError
		if errors.As(err, &validationError) {
			_, _ = w.Write([]byte(validationError.Error()))

			return
		}

		slog.Error("failed to create category", "error", err)

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

		return
	}

	_, _ = w.Write([]byte("create category page\n"))
}

func (h *CreateCategoryHandler) handle(r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return internalerrors.ErrParsingForm
	}

	request := requests.CreateCategoryRequest{
		Name:          r.FormValue("name"),
		Income:        r.FormValue("income"),
		Visible:       r.FormValue("visible"),
		Supercategory: r.FormValue("supercategory"),
	}

	return h.categoryCreator.Create(r.Context(), request)
}
