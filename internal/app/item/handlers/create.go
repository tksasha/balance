package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"

	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/app/item/component"
	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/balance/internal/common/component/path"
	"github.com/tksasha/balance/internal/common/handler"
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
	if err := r.ParseForm(); err != nil {
		h.SetError(w, common.ErrParsingForm)

		return
	}

	item, err := h.itemService.Create(r.Context(), h.parseRequest(r))
	if err == nil {
		h.StatusOK(w, r.URL.Query(), item)

		return
	}

	var verrors validation.Errors
	if errors.As(err, &verrors) {
		categories, err := h.categoryService.List(r.Context())
		if err != nil {
			h.SetError(w, err)

			return
		}

		w.Header().Add("Hx-Trigger-After-Swap", "balance.item.create.error")

		err = h.component.Create(item, categories, verrors).Render(w)

		h.SetError(w, err)

		return
	}

	h.SetError(w, err)
}

func (h *CreateHandler) parseRequest(r *http.Request) item.CreateRequest {
	return item.CreateRequest{
		Date:        r.FormValue("date"),
		Formula:     r.FormValue("formula"),
		CategoryID:  r.FormValue("category_id"),
		Description: r.FormValue("description"),
	}
}

func (h *CreateHandler) StatusOK(w http.ResponseWriter, values url.Values, item *item.Item) {
	params := path.Params{
		"month": strconv.Itoa(int(item.Date.Month())),
		"year":  strconv.Itoa(item.Date.Year()),
	}

	header := map[string]map[string]string{
		"balance.item.created": {
			"itemsPath":      path.Items(params, values),
			"categoriesPath": path.Categories(params),
			"balancePath":    path.Balance(values),
		},
	}

	writer := bytes.NewBuffer([]byte{})

	if err := json.NewEncoder(writer).Encode(header); err != nil {
		slog.Error("failed to encode", "error", err)
	}

	w.Header().Add("Hx-Trigger-After-Swap", writer.String())

	w.WriteHeader(http.StatusOK)
}
