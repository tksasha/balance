package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/core/common/helpers"
	"github.com/tksasha/balance/internal/core/common/providers"
	"github.com/tksasha/balance/internal/core/common/tests"
	"github.com/tksasha/balance/internal/core/index/components"
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

	timeProvider := providers.NewTimeProvider()

	helpers := helpers.New(timeProvider)

	monthsComonents := components.NewMonthsComponent(helpers)

	indexPageComponent := components.NewIndexPageComponent(helpers, monthsComonents)

	mux := tests.NewMux(t, "/", handlers.NewHandler(service, categoryService, indexPageComponent))

	t.Run("renders 200 when there no errors", func(t *testing.T) {
		request := tests.NewGetRequest(ctx, t, "/")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)
	})
}
