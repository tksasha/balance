package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/core/common/tests"
	handlers "github.com/tksasha/balance/internal/core/index/handler"
	"gotest.tools/v3/assert"
)

func TestIndexPageHandler_ServeHTTP(t *testing.T) {
	ctx := t.Context()

	service, db := tests.NewIndexService(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	categoryService, db2 := tests.NewCategoryService(ctx, t)
	defer func() {
		_ = db2.Close()
	}()

	mux := tests.NewMux(t, "/", handlers.NewHandler(service, categoryService))

	t.Run("renders 200 when there no errors", func(t *testing.T) {
		request := tests.NewGetRequest(ctx, t, "/")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)
	})
}
