package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/app/category/handlers"
	"github.com/tksasha/balance/internal/app/category/repository"
	"github.com/tksasha/balance/internal/app/category/service"
	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/db/nameprovider"
	"github.com/tksasha/balance/internal/server/middlewares"
	"gotest.tools/v3/assert"
)

func TestIndexHandler(t *testing.T) {
	ctx := t.Context()

	db := db.Open(ctx, nameprovider.NewTestProvider())
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("failed to close db: %v", err)
		}
	}()

	categoryRepository := repository.New(db)

	categoryService := service.New(categoryRepository)

	indexHandler := handlers.NewIndexHandler(categoryService)

	next := http.Handler(indexHandler)

	for _, middleware := range middlewares.New() {
		next = middleware.Wrap(next)
	}

	mux := http.NewServeMux()

	mux.Handle("GET /categories", next)

	t.Run("render index.html", func(t *testing.T) {
		cleanup(t, db)

		createCategory(t, db, &category.Category{ID: 1, Name: "Food"})
		createCategory(t, db, &category.Category{ID: 2, Name: "Beverages"})
		createCategory(t, db, &category.Category{ID: 3, Name: "Salary", Income: true})

		createItem(t, db, &item.Item{Date: date(t, "2025-03-01"), Sum: 11.11, CategoryID: 1})
		createItem(t, db, &item.Item{Date: date(t, "2025-03-02"), Sum: 22.22, CategoryID: 2})
		createItem(t, db, &item.Item{Date: date(t, "2025-03-03"), Sum: 33.33, CategoryID: 3})

		url := "/categories?year=2025&month=03"

		request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)
	})
}
