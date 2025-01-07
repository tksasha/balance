package handlers_test

import (
	"context"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/middlewares"
	"github.com/tksasha/balance/internal/models"
	providers "github.com/tksasha/balance/internal/providers/test"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/services"
	"gotest.tools/v3/assert"
)

func TestCreateItemHandler_ServeHTTP(t *testing.T) { //nolint:funlen
	dbNameProvider := providers.NewDBNameProvider()

	dbConnection := db.Open(dbNameProvider)

	itemRepository := repositories.NewItemRepository(dbConnection)
	categoryRepository := repositories.NewCategoryRepository(dbConnection)

	itemService := services.NewItemService(itemRepository)
	categoryService := services.NewCategoryService(categoryRepository)

	middleware := middlewares.NewCurrencyMiddleware().Wrap(
		handlers.NewCreateItemHandler(itemService, categoryService),
	)

	route := http.NewServeMux()
	route.Handle("POST /items", middleware)

	ctx := context.Background()

	t.Run("when form parsing error is happened, it should respond with 400", func(t *testing.T) {
		body := strings.NewReader("%")

		request, err := http.NewRequestWithContext(ctx, http.MethodPost, "/items", body)
		assert.NilError(t, err)

		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		recorder := httptest.NewRecorder()

		route.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("when input data is invalid it should render form", func(t *testing.T) {
		formData := url.Values{}
		formData.Set("date", "")

		body := strings.NewReader(formData.Encode())

		request, err := http.NewRequestWithContext(ctx, http.MethodPost, "/items", body)
		assert.NilError(t, err)

		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		recorder := httptest.NewRecorder()

		route.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("when input data is valid, it should respond with 200", func(t *testing.T) {
		t.Skip("TODO: will fix later")

		formData := url.Values{}
		formData.Set("date", "2024-10-16")

		body := strings.NewReader(formData.Encode())

		request, err := http.NewRequestWithContext(ctx, http.MethodPost, "/items", body)
		if err != nil {
			t.Fatalf("failed to build request with context, error: %v", err)
		}

		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		recorder := httptest.NewRecorder()

		route.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)

		item, err := getItemByDate(ctx, dbConnection.Connection, "2024-10-16")
		if err != nil {
			t.Fatalf("failed to get item by date, error: %v", err)
		}

		assert.Equal(t, "2024-10-16", item.Date.String())
	})
}

func getItemByDate(ctx context.Context, db *sql.DB, date string) (*models.Item, error) {
	var item *models.Item

	query := `SELECT id, date FROM items WHERE date="?" ORDER BY created_at LIMIT 1`

	if err := db.QueryRowContext(ctx, query, date).Scan(item); err != nil {
		return nil, err
	}

	return item, nil
}
