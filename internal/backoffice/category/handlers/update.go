package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/balance/internal/backoffice/category/component"
	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/balance/internal/common/handler"
	"github.com/tksasha/balance/internal/common/paths"
	"github.com/tksasha/balance/internal/common/paths/params"
	"github.com/tksasha/validation"
)

type UpdateHandler struct {
	*handler.Handler

	categoryService category.Service
	component       *component.Component
}

func NewUpdateHandler(
	categoryService category.Service,
) *UpdateHandler {
	return &UpdateHandler{
		Handler:         handler.New(),
		categoryService: categoryService,
		component:       component.New(),
	}
}

func (h *UpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	category, err := h.handle(r)
	if err == nil {
		params := params.New().WithCurrency(category.Currency)

		h.ok(w, params)

		return
	}

	var verrors validation.Errors
	if errors.As(err, &verrors) {
		err := h.component.Update(category, verrors).Render(w)

		h.SetError(w, err)

		return
	}

	h.SetError(w, err)
}

func (h *UpdateHandler) handle(r *http.Request) (*category.Category, error) {
	if err := r.ParseForm(); err != nil {
		return nil, common.ErrParsingForm
	}

	request := category.UpdateRequest{
		ID:            r.PathValue("id"),
		Name:          r.FormValue("name"),
		Income:        r.FormValue("income"),
		Visible:       r.FormValue("visible"),
		Supercategory: r.FormValue("supercategory"),
		Number:        r.FormValue("number"),
		Currency:      r.FormValue("currency"),
	}

	return h.categoryService.Update(r.Context(), request)
}

func (h *UpdateHandler) ok(w http.ResponseWriter, params params.Params) {
	header := map[string]map[string]string{
		"backoffice.category.updated": {
			"backofficeCategoriesPath": paths.BackofficeCategories(params),
		},
	}

	writer := bytes.NewBuffer([]byte{})

	if err := json.NewEncoder(writer).Encode(header); err != nil {
		slog.Error("failed to encode", "error", err)

		writer.Reset()
	}

	w.Header().Add("Hx-Trigger-After-Swap", writer.String())

	w.WriteHeader(http.StatusOK)
}
