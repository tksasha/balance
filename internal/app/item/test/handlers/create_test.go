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
	"github.com/tksasha/balance/internal/app/item/components"
	"github.com/tksasha/balance/internal/app/item/handlers"
	"github.com/tksasha/balance/internal/app/item/repository"
	"github.com/tksasha/balance/internal/app/item/service"
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/db/nameprovider"
	"gotest.tools/v3/assert"
)

func TestItemCreateHandler(t *testing.T) { //nolint:funlen
	handler, db := newCreateHandler(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	mux := mux(t, "POST /items", handler)

	ctx := t.Context()

	t.Run("responds 400 on parse form fails", func(t *testing.T) {
		request, err := http.NewRequestWithContext(ctx, http.MethodPost, "/items", nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("renders errors on invalid input", func(t *testing.T) {
		cleanup(t, db)

		values := url.Values{"date": {""}}

		request, err := http.NewRequestWithContext(ctx, http.MethodPost, "/items", strings.NewReader(values.Encode()))
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("responds 204 when item created", func(t *testing.T) {
		cleanup(t, db)

		categoryToCreate := &category.Category{
			ID:       1101,
			Name:     "Accoutrements",
			Currency: currency.USD,
		}

		createCategory(t, db, categoryToCreate)

		values := url.Values{
			"date":        {"2024-10-16"},
			"formula":     {"42.69+69.42"},
			"category_id": {"1101"},
			"description": {"paper clips, notebooks, and pens"},
		}

		request, err := http.NewRequestWithContext(
			ctx,
			http.MethodPost,
			"/items?currency=usd",
			strings.NewReader(values.Encode()),
		)
		if err != nil {
			t.Fatal(err)
		}

		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNoContent)

		item := findItemByDate(t, db, currency.USD, "2024-10-16")

		assert.Equal(t, item.Date.Format(time.DateOnly), "2024-10-16")
		assert.Equal(t, item.CategoryID, 1101)
		assert.Equal(t, item.CategoryName.String, "Accoutrements")
		assert.Equal(t, item.Currency, currency.USD)
		assert.Equal(t, item.Formula, "42.69+69.42")
		assert.Equal(t, item.Sum, 112.11)
		assert.Equal(t, item.Description, "paper clips, notebooks, and pens")
	})
}

func newCreateHandler(t *testing.T) (*handlers.CreateHandler, *sql.DB) {
	t.Helper()

	db := db.Open(t.Context(), nameprovider.NewTestProvider())

	itemRepository := repository.New(db)

	categoryRepository := categoryrepository.New(db)

	itemService := service.New(itemRepository, categoryRepository)

	categoryService := categoryservice.New(categoryRepository)

	itemComponent := components.NewItemsComponent()

	handler := handlers.NewCreateHandler(itemService, categoryService, itemComponent)

	return handler, db
}
