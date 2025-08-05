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

type CreateHandler struct {
	*handler.Handler

	categoryService category.Service
	component       *component.Component
}

func NewCreateHandler(
	categoryService category.Service,
) *CreateHandler {
	return &CreateHandler{
		Handler:         handler.New(),
		categoryService: categoryService,
		component:       component.New(),
	}
}

func (h *CreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	category, err := h.handle(r)
	if err != nil {
		h.errors(w, category, err)

		return
	}

	params := params.New().WithCurrency(category.Currency)

	h.ok(w, params)
}

func (h *CreateHandler) handle(r *http.Request) (*category.Category, error) {
	if err := r.ParseForm(); err != nil {
		return nil, common.ErrParsingForm
	}

	request := category.CreateRequest{
		Name:          r.FormValue("name"),
		Income:        r.FormValue("income"),
		Visible:       r.FormValue("visible"),
		Supercategory: r.FormValue("supercategory"),
		Number:        r.FormValue("number"),
		Currency:      r.FormValue("currency"),
	}

	return h.categoryService.Create(r.Context(), request)
}

func (h *CreateHandler) ok(w http.ResponseWriter, params params.Params) {
	writer := bytes.NewBuffer([]byte{})

	header := map[string]map[string]string{
		"backoffice.category.created": {
			"backofficeCategoriesPath": paths.BackofficeCategories(params),
		},
	}

	if err := json.NewEncoder(writer).Encode(header); err != nil {
		slog.Error("failed to encode", "error", err)

		writer.Reset()
	}

	w.Header().Add("Hx-Trigger-After-Swap", writer.String())

	w.WriteHeader(http.StatusOK)
}

func (h *CreateHandler) errors(w http.ResponseWriter, category *category.Category, err error) {
	var verrors validation.Errors
	if errors.As(err, &verrors) {
		err := h.component.Create(category, verrors).Render(w)

		h.SetError(w, err)

		return
	}

	h.SetError(w, err)
}
