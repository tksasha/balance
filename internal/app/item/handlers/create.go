package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/app/item/component"
	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/balance/internal/common/handler"
	"github.com/tksasha/balance/internal/common/paths"
	"github.com/tksasha/balance/internal/common/paths/params"
	"github.com/tksasha/validation"
)

type CreateHandler struct {
	*handler.Handler

	itemService     item.Service
	categoryService item.CategoryService
	component       *component.Component
}

func NewCreateHandler(
	itemService item.Service,
	categoryService item.CategoryService,
) *CreateHandler {
	return &CreateHandler{
		Handler:         handler.New(),
		itemService:     itemService,
		categoryService: categoryService,
		component:       component.New(),
	}
}

func (h *CreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	categories, err := h.categoryService.List(r.Context())
	if err != nil {
		h.SetError(w, err)

		return
	}

	params := params.New(r.URL.Query())

	item, err := h.handle(r)
	if err != nil {
		h.errors(w, params, item, categories, err)

		return
	}

	h.ok(w, params, item, categories)
}

func (h *CreateHandler) handle(r *http.Request) (*item.Item, error) {
	if err := r.ParseForm(); err != nil {
		return nil, common.ErrParsingForm
	}

	request := item.CreateRequest{
		Date:        r.FormValue("date"),
		Formula:     r.FormValue("formula"),
		CategoryID:  r.FormValue("category_id"),
		Description: r.FormValue("description"),
	}

	return h.itemService.Create(r.Context(), request)
}

func (h *CreateHandler) ok(
	w http.ResponseWriter,
	params params.Params,
	item *item.Item,
	categories category.Categories,
) {
	month, year := int(item.Date.Month()), item.Date.Year()

	params = params.WithMonth(month).WithYear(year)

	header := map[string]map[string]string{
		"balance.item.created": {
			"itemsPath":      paths.Items(params),
			"categoriesPath": paths.Categories(params),
			"balancePath":    paths.Balance(params),
		},
	}

	writer := bytes.NewBuffer([]byte{})

	if err := json.NewEncoder(writer).Encode(header); err != nil {
		slog.Error("failed to encode", "error", err)

		writer.Reset()
	}

	w.Header().Add("Hx-Trigger-After-Swap", writer.String())

	w.WriteHeader(http.StatusOK)

	if err := h.component.New(params, categories).Render(w); err != nil {
		h.SetError(w, err)
	}
}

func (h *CreateHandler) errors(
	w http.ResponseWriter,
	params params.Params,
	item *item.Item,
	categories category.Categories,
	err error,
) {
	var verrors validation.Errors
	if errors.As(err, &verrors) {
		w.Header().Add("Hx-Trigger-After-Swap", "balance.item.create.error")

		w.Header().Add("Hx-Retarget", "#modal-body")

		err = h.component.Create(params, item, categories, verrors).Render(w)

		h.SetError(w, err)

		return
	}

	h.SetError(w, err)
}
