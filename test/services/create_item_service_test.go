package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/requests"
	"github.com/tksasha/balance/internal/services"
	mocksforservices "github.com/tksasha/balance/mocks/services"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestCreateItemService(t *testing.T) { //nolint:funlen
	controller := gomock.NewController(t)

	itemRepository := mocksforservices.NewMockItemRepository(controller)
	categoryRepository := mocksforservices.NewMockCategoryRepository(controller)

	service := services.NewCreateItemService(itemRepository, categoryRepository)

	ctx := context.Background()

	t.Run("when date is blank, it should return error", func(t *testing.T) {
		request := requests.CreateItemRequest{
			Date:       "",
			Formula:    "2+2",
			CategoryID: "1321",
		}

		category := &models.Category{
			ID:   1321,
			Name: "Nanotechnology",
		}

		categoryRepository.
			EXPECT().
			FindByID(ctx, 1321).
			Return(category, nil)

		_, err := service.Create(ctx, request)

		assert.Error(t, err, "date: is required")
	})

	t.Run("when date is invalid, it should return error", func(t *testing.T) {
		request := requests.CreateItemRequest{
			Date:       "abc",
			Formula:    "2+2",
			CategoryID: "1320",
		}

		category := &models.Category{
			ID:   1320,
			Name: "Cybersecurity",
		}

		categoryRepository.
			EXPECT().
			FindByID(ctx, 1320).
			Return(category, nil)

		_, err := service.Create(ctx, request)

		assert.Error(t, err, "date: is invalid")
	})

	t.Run("when formula is blank, it should return error", func(t *testing.T) {
		request := requests.CreateItemRequest{
			Date:       "2025-01-23",
			Formula:    "",
			CategoryID: "1318",
		}

		category := &models.Category{
			ID:   1318,
			Name: "Biotechnology",
		}

		categoryRepository.
			EXPECT().
			FindByID(ctx, 1318).
			Return(category, nil)

		_, err := service.Create(ctx, request)

		assert.Error(t, err, "sum: is required")
	})

	t.Run("when formula is invalid, it should return error", func(t *testing.T) {
		request := requests.CreateItemRequest{
			Date:       "2025-01-23",
			Formula:    "abc",
			CategoryID: "1315",
		}

		category := &models.Category{
			ID:   1315,
			Name: "Pharmaceuticals",
		}

		categoryRepository.
			EXPECT().
			FindByID(ctx, 1315).
			Return(category, nil)

		_, err := service.Create(ctx, request)

		assert.Error(t, err, "sum: is invalid")
	})

	t.Run("when category_id is blank, it should return error", func(t *testing.T) {
		request := requests.CreateItemRequest{
			Date:       "2025-01-23",
			Formula:    "42.69+69.42",
			CategoryID: "",
		}

		_, err := service.Create(ctx, request)

		assert.Error(t, err, "category_id: is required")
	})

	t.Run("when category_id is invalid, it should return error", func(t *testing.T) {
		request := requests.CreateItemRequest{
			Date:       "2025-01-23",
			Formula:    "42.69+69.42",
			CategoryID: "abc",
		}

		_, err := service.Create(ctx, request)

		assert.Error(t, err, "category_id: is invalid")
	})

	t.Run("when find category by id returns error, it should return error", func(t *testing.T) {
		request := requests.CreateItemRequest{
			Date:       "2025-01-23",
			Formula:    "42.69+69.42",
			CategoryID: "1237",
		}

		categoryRepository.
			EXPECT().
			FindByID(ctx, 1237).
			Return(nil, errors.New("find category by id error"))

		_, err := service.Create(ctx, request)

		assert.Error(t, err, "category_id: internal error")
	})

	t.Run("when create item returns error, it should return error", func(t *testing.T) {
		request := requests.CreateItemRequest{
			Date:        "2025-01-23",
			Formula:     "42.69+69.42",
			CategoryID:  "1244",
			Description: "health, beauty & wellness",
		}

		category := &models.Category{
			ID:   1244,
			Name: "Entrepreneurship",
		}

		categoryRepository.
			EXPECT().
			FindByID(ctx, 1244).
			Return(category, nil)

		item := &models.Item{
			Date:         date(t, "2025-01-23"),
			Formula:      "42.69+69.42",
			Sum:          112.11,
			CategoryID:   1244,
			CategoryName: "Entrepreneurship",
			Description:  "health, beauty & wellness",
		}

		itemRepository.
			EXPECT().
			Create(ctx, item).
			Return(errors.New("create item error"))

		_, err := service.Create(ctx, request)

		assert.Error(t, err, "create item error")
	})

	t.Run("when create item doesn't return error, it should create item", func(t *testing.T) {
		request := requests.CreateItemRequest{
			Date:        "2025-01-23",
			Formula:     "42.69+69.42",
			CategoryID:  "1307",
			Description: "arts, crafts & hobbies",
		}

		category := &models.Category{
			ID:   1244,
			Name: "Telecommunications",
		}

		categoryRepository.
			EXPECT().
			FindByID(ctx, 1307).
			Return(category, nil)

		item := &models.Item{
			Date:         date(t, "2025-01-23"),
			Formula:      "42.69+69.42",
			Sum:          112.11,
			CategoryID:   1307,
			CategoryName: "Telecommunications",
			Description:  "arts, crafts & hobbies",
		}

		itemRepository.
			EXPECT().
			Create(ctx, item).
			Return(nil)

		_, err := service.Create(ctx, request)

		assert.NilError(t, err)
	})
}
