package handlers_test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/balance/internal/backoffice/category/component"
	"github.com/tksasha/balance/internal/backoffice/category/handlers"
	"github.com/tksasha/balance/internal/backoffice/category/repository"
	"github.com/tksasha/balance/internal/backoffice/category/service"
	"github.com/tksasha/balance/internal/common"
	commonComponent "github.com/tksasha/balance/internal/common/component"
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/db"
	nameprovider "github.com/tksasha/balance/internal/db/nameprovider/test"
	"gotest.tools/v3/assert"
)

func TestCategoryEditHandler(t *testing.T) {
	handler, db := newEditHandler(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	mux := mux(t, "GET /categories/{id}/edit", handler)

	ctx := t.Context()

	t.Run("responds 404 on category not found", func(t *testing.T) {
		request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/categories/1004/edit", nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("responds 200 on category found", func(t *testing.T) {
		cleanup(t, db)

		categoryToCreate := &category.Category{
			ID:       1010,
			Name:     "Xenomorphic",
			Currency: currency.EUR,
		}

		createCategory(t, db, categoryToCreate)

		request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/categories/1010/edit?currency=eur", nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)
	})
}

func newEditHandler(t *testing.T) (*handlers.EditHandler, *sql.DB) {
	t.Helper()

	db := db.Open(t.Context(), nameprovider.New())

	categoryRepository := repository.New(db)

	categoryService := service.New(common.NewBaseService(), categoryRepository)

	categoryComponent := component.New(commonComponent.New())

	handler := handlers.NewEditHandler(categoryService, categoryComponent)

	return handler, db
}
