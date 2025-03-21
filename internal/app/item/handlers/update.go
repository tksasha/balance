package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/app/item/component"
	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/balance/internal/common/handler"
	"github.com/tksasha/balance/internal/common/paths"
	"github.com/tksasha/balance/internal/common/paths/params"
	"github.com/tksasha/validation"
)

type UpdateHandler struct {
	*handler.Handler

	itemService     item.Service
	categoryService item.CategoryService
	component       *component.Component
}

func NewUpdateHandler(
	itemService item.Service,
	categoryService item.CategoryService,
) *UpdateHandler {
	return &UpdateHandler{
		Handler:         handler.New(),
		itemService:     itemService,
		categoryService: categoryService,
		component:       component.New(),
	}
}

func (h *UpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := params.New(r.URL.Query())

	if err := r.ParseForm(); err != nil {
		h.SetError(w, common.ErrParsingForm)

		return
	}

	request := item.UpdateRequest{
		ID:          r.PathValue("id"),
		Date:        r.FormValue("date"),
		Formula:     r.FormValue("formula"),
		CategoryID:  r.FormValue("category_id"),
		Description: r.FormValue("description"),
	}

	item, err := h.itemService.Update(r.Context(), request)
	if err == nil {
		h.ok(w, params, item)

		return
	}

	var verrors validation.Errors
	if errors.As(err, &verrors) {
		categories, err := h.categoryService.List(r.Context())
		if err != nil {
			h.SetError(w, err)

			return
		}

		err = h.component.Edit(params, item, categories, verrors).Render(w)

		h.SetError(w, err)

		return
	}

	h.SetError(w, err)
}

func (h *UpdateHandler) ok(w http.ResponseWriter, params params.Params, item *item.Item) {
	month, year := int(item.Date.Month()), item.Date.Year()

	params = params.WithMonth(month).WithYear(year)

	headers := map[string]map[string]string{
		"balance.item.updated": {
			"categoriesPath": paths.Categories(params),
			"balancePath":    paths.Balance(params),
		},
	}

	writer := bytes.NewBuffer([]byte{})

	if err := json.NewEncoder(writer).Encode(headers); err != nil {
		slog.Error("failed to encode", "error", err)
	}

	w.Header().Add("Hx-Trigger-After-Swap", writer.String())

	err := h.component.Update(params, item).Render(w)

	h.SetError(w, err)
}
