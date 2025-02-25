package handler_test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	categoryrepository "github.com/tksasha/balance/internal/app/category/repository"
	categoryservice "github.com/tksasha/balance/internal/app/category/service"
	"github.com/tksasha/balance/internal/app/index/component"
	"github.com/tksasha/balance/internal/app/index/handler"
	"github.com/tksasha/balance/internal/app/index/repository"
	"github.com/tksasha/balance/internal/app/index/service"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/db/nameprovider"
	"gotest.tools/v3/assert"
)

func TestIndexHandler(t *testing.T) {
	handler, db := newIndexHandler(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	mux := mux(t, "/", handler)

	ctx := t.Context()

	t.Run("renders 200 when there no errors", func(t *testing.T) {
		request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)
	})
}

func newIndexHandler(t *testing.T) (*handler.Handler, *sql.DB) {
	t.Helper()

	db := db.Open(t.Context(), nameprovider.NewTestProvider())

	indexRepository := repository.New(db)

	indexService := service.New(indexRepository)

	categoryRepository := categoryrepository.New(db)

	categoryService := categoryservice.New(categoryRepository)

	indexComponent := component.NewIndexComponent()

	handler := handler.New(indexService, categoryService, indexComponent)

	return handler, db
}
