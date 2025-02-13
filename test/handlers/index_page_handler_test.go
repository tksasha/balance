package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/common/testutils"
	"github.com/tksasha/balance/internal/handlers"
	"gotest.tools/v3/assert"
)

func TestIndexPageHandler_ServeHTTP(t *testing.T) {
	ctx := t.Context()

	service, db := testutils.NewCategoryService(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	mux := newMux(t, "/", handlers.NewIndexPageHandler(service))

	t.Run("renders 200 when there no errors", func(t *testing.T) {
		request := newGetRequest(ctx, t, "/")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)
	})
}
