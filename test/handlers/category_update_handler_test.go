package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/pkg/currencies"
	"gotest.tools/v3/assert"
)

func TestCategoryUpdateHandler(t *testing.T) { //nolint:funlen
	ctx := t.Context()

	service, db := newCategoryService(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	mux := newMux(t, "PATCH /categories/{id}", handlers.NewCategoryUpdateHandler(service))

	t.Run("responds 404 on no category found", func(t *testing.T) {
		cleanup(ctx, t)

		request := newPatchRequest(ctx, t, "/categories/1141", nil)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("responds 200 on duplicates name", func(t *testing.T) {
		cleanup(ctx, t)

		for id, name := range map[int]string{1151: "Heterogeneous", 11654: "Paraphernalia"} {
			categoryToCreate := &models.Category{
				ID:       id,
				Name:     name,
				Currency: currencies.USD,
			}

			createCategory(ctx, t, categoryToCreate)
		}

		request := newPatchRequest(ctx, t,
			"/categories/1151?currency=usd",
			Params{"name": "Paraphernalia"},
		)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)
	})

	t.Run("responds 200 on successful update", func(t *testing.T) {
		cleanup(ctx, t)

		categoryToCreate := &models.Category{
			ID:            1208,
			Name:          "Paraphernalia",
			Income:        false,
			Visible:       false,
			Currency:      currencies.USD,
			Supercategory: 5,
		}

		createCategory(ctx, t, categoryToCreate)

		request := newPatchRequest(ctx, t, "/categories/1208?currency=usd",
			Params{
				"name":          "Heterogeneous",
				"income":        "true",
				"visible":       "true",
				"supercategory": "4",
			},
		)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)

		category := findCategoryByID(ctx, t, currencies.USD, 1208)

		assert.Equal(t, category.ID, 1208)
		assert.Equal(t, category.Name, "Heterogeneous")
		assert.Equal(t, category.Income, true)
		assert.Equal(t, category.Visible, true)
		assert.Equal(t, category.Currency, currencies.USD)
		assert.Equal(t, category.Supercategory, 4)
	})
}
