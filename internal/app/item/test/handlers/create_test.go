package handlers_test

import (
	"database/sql"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/tksasha/balance/internal/app/category"
	categoryrepository "github.com/tksasha/balance/internal/app/category/repository"
	categoryservice "github.com/tksasha/balance/internal/app/category/service"
	"github.com/tksasha/balance/internal/app/item/handlers"
	"github.com/tksasha/balance/internal/app/item/repository"
	"github.com/tksasha/balance/internal/app/item/service"
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/db/nameprovider"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/golden"
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

	t.Run("renders edit with errors when input is invalid", func(t *testing.T) {
		values := url.Values{"date": {""}}

		body := strings.NewReader(values.Encode())

		request, err := http.NewRequestWithContext(ctx, http.MethodPost, "/items", body)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)

		response, err := io.ReadAll(recorder.Body)
		if err != nil {
			t.Fatal(err)
		}

		golden.Assert(t, string(response), "create-with-errors.html")
	})

	t.Run("renders created item", func(t *testing.T) {
		cleanup(t, db)

		categoryToCreate := &category.Category{
			ID:       1101,
			Name:     "Accoutrements",
			Slug:     "accoutrements",
			Currency: currency.USD,
		}

		createCategory(t, db, categoryToCreate)

		values := url.Values{
			"date":        {"2024-10-16"},
			"formula":     {"42.69+69.42"},
			"category_id": {"1101"},
			"description": {"paper clips, notebooks, and pens"},
		}

		body := strings.NewReader(values.Encode())

		request, err := http.NewRequestWithContext(ctx, http.MethodPost, "/items?currency=usd", body)
		if err != nil {
			t.Fatal(err)
		}

		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusCreated)

		item := findItemByDate(t, db, currency.USD, "2024-10-16")

		assert.Equal(t, item.Currency, currency.USD)
		assert.Equal(t, item.Date.Format(time.DateOnly), "2024-10-16")
		assert.Equal(t, item.Formula, "42.69+69.42")
		assert.Equal(t, item.Sum, 112.11)
		assert.Equal(t, item.CategoryID, 1101)
		assert.Equal(t, item.CategoryName, "Accoutrements")
		assert.Equal(t, item.CategorySlug, "accoutrements")
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

	handler := handlers.NewCreateHandler(itemService, categoryService)

	return handler, db
}
