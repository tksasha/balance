package handlers_test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/backoffice/category/component"
	"github.com/tksasha/balance/internal/backoffice/category/handlers"
	"github.com/tksasha/balance/internal/backoffice/category/repository"
	"github.com/tksasha/balance/internal/backoffice/category/service"
	"github.com/tksasha/balance/internal/common"
	commoncomponent "github.com/tksasha/balance/internal/common/component"
	"github.com/tksasha/balance/internal/common/tests"
	"github.com/tksasha/balance/internal/db"
	nameprovider "github.com/tksasha/balance/internal/db/nameprovider/test"
	"gotest.tools/v3/assert"
)

func TestListCategories(t *testing.T) {
	handler, db := newListHandler(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	mux := tests.NewMux(t, "GET /backoffice/categories", handler)

	ctx := t.Context()

	t.Run("responds with 200 when there no errors", func(t *testing.T) {
		request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/backoffice/categories", nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)
	})
}

func newListHandler(t *testing.T) (*handlers.ListHandler, *sql.DB) {
	t.Helper()

	db := db.Open(t.Context(), nameprovider.New())

	categoryRepository := repository.New(common.NewBaseRepository(), db)

	categoryService := service.New(categoryRepository)

	categoryComponent := component.New(
		commoncomponent.New(),
	)

	handler := handlers.NewListHandler(common.NewBaseHandler(), categoryService, categoryComponent)

	return handler, db
}
