package handlers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/app/category/handlers"
	"github.com/tksasha/balance/internal/app/category/repository"
	"github.com/tksasha/balance/internal/app/category/service"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/db/nameprovider"
	"github.com/tksasha/balance/internal/server/middlewares"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/golden"
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

		createCategory(t, db, &category.Category{Name: "Food", Slug: "food"})
		createCategory(t, db, &category.Category{Name: "Beverages", Slug: "beverages"})
		createCategory(t, db, &category.Category{Name: "Salary", Slug: "salary"})

		request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/categories", nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)

		response, err := io.ReadAll(recorder.Body)
		if err != nil {
			t.Fatal(err)
		}

		golden.Assert(t, string(response), "index.html")
	})
}
