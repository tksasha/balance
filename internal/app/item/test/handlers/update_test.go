package handlers_test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/tksasha/balance/internal/app/category"
	categoryrepository "github.com/tksasha/balance/internal/app/category/repository"
	categoryservice "github.com/tksasha/balance/internal/app/category/service"
	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/app/item/handlers"
	"github.com/tksasha/balance/internal/app/item/repository"
	"github.com/tksasha/balance/internal/app/item/service"
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/db/nameprovider"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/golden"
)

func TestItemUpdateHandler(t *testing.T) { //nolint:funlen
	handler, db := newUpdateHandler(t)

	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	mux := mux(t, "PATCH /items/{id}", handler)

	ctx := t.Context()

	t.Run("responds 400 on invalid input", func(t *testing.T) {
		cleanup(t, db)

		request, err := http.NewRequestWithContext(ctx, http.MethodPatch, "/items/1138", nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusBadRequest)
	})

	t.Run("responds 404 on no item found", func(t *testing.T) {
		values := url.Values{}

		body := strings.NewReader(values.Encode())

		request, err := http.NewRequestWithContext(ctx, http.MethodPatch, "/items/1218", body)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("renders edit with errors when input is invalid", func(t *testing.T) {
		cleanup(t, db)

		categoryToCreate := &category.Category{
			ID:       1157,
			Currency: currency.UAH,
			Name:     "Food",
		}

		createCategory(t, db, categoryToCreate)

		itemToCreate := &item.Item{
			ID:           1158,
			Currency:     currency.UAH,
			CategoryID:   1157,
			CategoryName: "Food",
		}

		createItem(t, db, itemToCreate)

		values := url.Values{}

		body := strings.NewReader(values.Encode())

		request, err := http.NewRequestWithContext(ctx, http.MethodPatch, "/items/1158", body)
		if err != nil {
			t.Fatal(err)
		}

		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)
	})

	t.Run("renders updated item if it is valid", func(t *testing.T) {
		cleanup(t, db)

		categoryToCreate := &category.Category{
			ID:       1148,
			Name:     "Pharmaceutical",
			Currency: currency.EUR,
		}

		createCategory(t, db, categoryToCreate)

		itemToCreate := &item.Item{
			ID:         1143,
			CategoryID: 1148,
			Currency:   currency.EUR,
		}

		createItem(t, db, itemToCreate)

		values := url.Values{
			"date":        {"25.01.2025"},
			"formula":     {"24 + 11 + 49"},
			"category_id": {"1148"},
			"description": {"pizza, ninja and disco"},
		}

		body := strings.NewReader(values.Encode())

		request, err := http.NewRequestWithContext(ctx, http.MethodPatch, "/items/1143?currency=eur", body)
		if err != nil {
			t.Fatal(err)
		}

		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)

		golden.Assert(t, recorder.Header().Get("Hx-Trigger-After-Swap"),
			"update-hx-trigger-after-swap-header.json")

		item := findItemByDate(t, db, currency.EUR, "2025-01-25")

		assert.Equal(t, item.ID, 1143)
		assert.Equal(t, item.Currency, currency.EUR)
		assert.Equal(t, item.Date, time.Date(2025, 1, 25, 0, 0, 0, 0, time.UTC))
		assert.Equal(t, item.Formula, "24 + 11 + 49")
		assert.Equal(t, item.Sum, 84.0)
		assert.Equal(t, item.CategoryID, 1148)
		assert.Equal(t, item.CategoryName, "Pharmaceutical")
		assert.Equal(t, item.Description, "pizza, ninja and disco")
	})
}

func newUpdateHandler(t *testing.T) (*handlers.UpdateHandler, *sql.DB) {
	t.Helper()

	db := db.Open(t.Context(), nameprovider.NewTestProvider())

	itemRepository := repository.New(db)

	categoryRepository := categoryrepository.New(db)

	itemService := service.New(itemRepository, categoryRepository)

	categoryService := categoryservice.New(categoryRepository)

	handler := handlers.NewUpdateHandler(itemService, categoryService)

	return handler, db
}
