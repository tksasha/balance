package handlers_test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/balance/internal/backoffice/category/handlers"
	"github.com/tksasha/balance/internal/backoffice/category/repository"
	"github.com/tksasha/balance/internal/backoffice/category/service"
	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/db"
	nameprovider "github.com/tksasha/balance/internal/db/nameprovider/test"
	"gotest.tools/v3/assert"
)

func TestCategoryDeleteHandler(t *testing.T) {
	handler, db := newDeleteHandler(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	mux := mux(t, "DELETE /backoffice/categories/{id}", handler)

	ctx := t.Context()

	t.Run("responds 404 when category not found", func(t *testing.T) {
		request, err := http.NewRequestWithContext(ctx, http.MethodDelete, "/backoffice/categories/1348", nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("responds 204 when category deleted", func(t *testing.T) {
		cleanup(t, db)

		categoryToCreate := &category.Category{
			ID:       1411,
			Name:     "Miscellaneous",
			Currency: currency.EUR,
		}

		createCategory(t, db, categoryToCreate)

		request, err := http.NewRequestWithContext(
			ctx,
			http.MethodDelete,
			"/backoffice/categories/1411?currency=eur",
			nil,
		)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNoContent)

		category := findCategoryByID(t, db, currency.EUR, 1411)

		assert.Equal(t, category.ID, 1411)
		assert.Equal(t, category.Name, "Miscellaneous")
		assert.Equal(t, category.Currency, currency.EUR)
		assert.Assert(t, !category.DeletedAt.Time.IsZero())
	})
}

func newDeleteHandler(t *testing.T) (*handlers.DeleteHandler, *sql.DB) {
	t.Helper()

	db := db.Open(t.Context(), nameprovider.New())

	categoryRepository := repository.New(common.NewBaseRepository(), db)

	categoryService := service.New(common.NewBaseService(), categoryRepository)

	handler := handlers.NewDeleteHandler(categoryService)

	return handler, db
}
