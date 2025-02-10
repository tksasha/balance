package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/requests"
	"github.com/tksasha/balance/internal/services"
	mocksforservices "github.com/tksasha/balance/mocks/services"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestItemService_Create(t *testing.T) { //nolint:funlen
	controller := gomock.NewController(t)

	itemRepository := mocksforservices.NewMockItemRepository(controller)
	categoryRepository := mocksforservices.NewMockCategoryRepository(controller)

	service := services.NewItemService(itemRepository, categoryRepository)

	ctx := context.Background()

	t.Run("returns error when date is blank", func(t *testing.T) {
		request := requests.ItemCreateRequest{
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

	t.Run("returns error when date is invalid", func(t *testing.T) {
		request := requests.ItemCreateRequest{
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

	t.Run("returns error when formula is blank", func(t *testing.T) {
		request := requests.ItemCreateRequest{
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

	t.Run("returns error when formula is invalid", func(t *testing.T) {
		request := requests.ItemCreateRequest{
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

	t.Run("returns error when category_id is blank", func(t *testing.T) {
		request := requests.ItemCreateRequest{
			Date:       "2025-01-23",
			Formula:    "42.69+69.42",
			CategoryID: "",
		}

		_, err := service.Create(ctx, request)

		assert.Error(t, err, "category: is required")
	})

	t.Run("returns error when category_id is invalid", func(t *testing.T) {
		request := requests.ItemCreateRequest{
			Date:       "2025-01-23",
			Formula:    "42.69+69.42",
			CategoryID: "abc",
		}

		_, err := service.Create(ctx, request)

		assert.Error(t, err, "category: is invalid")
	})

	t.Run("returns error when find by id failed", func(t *testing.T) {
		request := requests.ItemCreateRequest{
			Date:       "2025-01-23",
			Formula:    "42.69+69.42",
			CategoryID: "1237",
		}

		categoryRepository.
			EXPECT().
			FindByID(ctx, 1237).
			Return(nil, errors.New("find category by id error"))

		_, err := service.Create(ctx, request)

		assert.Error(t, err, "find category by id error")
	})

	t.Run("returns error when create failed", func(t *testing.T) {
		request := requests.ItemCreateRequest{
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

	t.Run("returns item when create succeeded", func(t *testing.T) {
		request := requests.ItemCreateRequest{
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

func TestItemService_Update(t *testing.T) { //nolint:funlen,maintidx
	controller := gomock.NewController(t)

	itemRepository := mocksforservices.NewMockItemRepository(controller)
	categoryRepository := mocksforservices.NewMockCategoryRepository(controller)

	service := services.NewItemService(itemRepository, categoryRepository)

	ctx := context.Background()

	t.Run("returns error when id is invalid", func(t *testing.T) {
		request := requests.ItemUpdateRequest{
			ID: "abc",
		}

		err := service.Update(ctx, request)

		assert.Error(t, err, "resource not found")
	})

	t.Run("returns error when item doesn't exist", func(t *testing.T) {
		request := requests.ItemUpdateRequest{
			ID: "1027",
		}

		itemRepository.
			EXPECT().
			FindByID(ctx, 1027).
			Return(nil, apperrors.ErrRecordNotFound)

		err := service.Update(ctx, request)

		assert.Error(t, err, "resource not found")
	})

	t.Run("returns error when find by id failed", func(t *testing.T) {
		request := requests.ItemUpdateRequest{
			ID: "1049",
		}

		itemRepository.
			EXPECT().
			FindByID(ctx, 1049).
			Return(nil, errors.New("find item by id error"))

		err := service.Update(ctx, request)

		assert.Error(t, err, "resource not found")
	})

	t.Run("returns error when date is blank", func(t *testing.T) {
		request := requests.ItemUpdateRequest{
			ID:         "1051",
			Date:       "",
			Formula:    "2+2",
			CategoryID: "1052",
		}

		item := &models.Item{}

		itemRepository.
			EXPECT().
			FindByID(ctx, 1051).
			Return(item, nil)

		category := &models.Category{
			ID:   1052,
			Name: "Tools",
		}

		categoryRepository.
			EXPECT().
			FindByID(ctx, 1052).
			Return(category, nil)

		err := service.Update(ctx, request)

		assert.Error(t, err, "date: is required")
	})

	t.Run("returns error when date is invalid", func(t *testing.T) {
		request := requests.ItemUpdateRequest{
			ID:         "1051",
			Date:       "abc",
			Formula:    "2+2",
			CategoryID: "1052",
		}

		item := &models.Item{}

		itemRepository.
			EXPECT().
			FindByID(ctx, 1051).
			Return(item, nil)

		category := &models.Category{
			ID:   1052,
			Name: "Tools",
		}

		categoryRepository.
			EXPECT().
			FindByID(ctx, 1052).
			Return(category, nil)

		err := service.Update(ctx, request)

		assert.Error(t, err, "date: is invalid")
	})

	t.Run("returns error when formula is blank", func(t *testing.T) {
		request := requests.ItemUpdateRequest{
			ID:         "1051",
			Date:       "2025-01-25",
			Formula:    "",
			CategoryID: "1052",
		}

		item := &models.Item{}

		itemRepository.
			EXPECT().
			FindByID(ctx, 1051).
			Return(item, nil)

		category := &models.Category{
			ID:   1052,
			Name: "Tools",
		}

		categoryRepository.
			EXPECT().
			FindByID(ctx, 1052).
			Return(category, nil)

		err := service.Update(ctx, request)

		assert.Error(t, err, "sum: is required")
	})

	t.Run("returns error when formula is invalid", func(t *testing.T) {
		request := requests.ItemUpdateRequest{
			ID:         "1051",
			Date:       "2025-01-25",
			Formula:    "abc",
			CategoryID: "1052",
		}

		item := &models.Item{}

		itemRepository.
			EXPECT().
			FindByID(ctx, 1051).
			Return(item, nil)

		category := &models.Category{
			ID:   1052,
			Name: "Tools",
		}

		categoryRepository.
			EXPECT().
			FindByID(ctx, 1052).
			Return(category, nil)

		err := service.Update(ctx, request)

		assert.Error(t, err, "sum: is invalid")
	})

	t.Run("returns error when category_id is blank", func(t *testing.T) {
		request := requests.ItemUpdateRequest{
			ID:         "1051",
			Date:       "2025-01-25",
			Formula:    "2+2",
			CategoryID: "",
		}

		item := &models.Item{}

		itemRepository.
			EXPECT().
			FindByID(ctx, 1051).
			Return(item, nil)

		err := service.Update(ctx, request)

		assert.Error(t, err, "category: is required")
	})

	t.Run("returns error when category_id is invalid", func(t *testing.T) {
		request := requests.ItemUpdateRequest{
			ID:         "1051",
			Date:       "2025-01-25",
			Formula:    "2+2",
			CategoryID: "abc",
		}

		item := &models.Item{}

		itemRepository.
			EXPECT().
			FindByID(ctx, 1051).
			Return(item, nil)

		err := service.Update(ctx, request)

		assert.Error(t, err, "category: is invalid")
	})

	t.Run("returns error when find category by id failed", func(t *testing.T) {
		request := requests.ItemUpdateRequest{
			ID:         "1051",
			Date:       "2025-01-25",
			Formula:    "2+2",
			CategoryID: "1100",
		}

		item := &models.Item{}

		itemRepository.
			EXPECT().
			FindByID(ctx, 1051).
			Return(item, nil)

		categoryRepository.
			EXPECT().
			FindByID(ctx, 1100).
			Return(nil, errors.New("find category by id error"))

		err := service.Update(ctx, request)

		assert.Error(t, err, "find category by id error")
	})

	t.Run("returns error when update failed", func(t *testing.T) {
		request := requests.ItemUpdateRequest{
			ID:          "1051",
			Date:        "2025-01-25",
			Formula:     "2+2",
			CategoryID:  "1100",
			Description: "food, wine and flowers",
		}

		itemFound := &models.Item{
			ID: 1051,
		}

		itemRepository.
			EXPECT().
			FindByID(ctx, 1051).
			Return(itemFound, nil)

		category := &models.Category{
			ID:   1100,
			Name: "Beverages",
		}

		categoryRepository.
			EXPECT().
			FindByID(ctx, 1100).
			Return(category, nil)

		itemToUpdate := &models.Item{
			ID:           1051,
			Date:         date(t, "2025-01-25"),
			Formula:      "2+2",
			Sum:          4,
			CategoryID:   1100,
			CategoryName: "Beverages",
			Description:  "food, wine and flowers",
		}

		itemRepository.
			EXPECT().
			Update(ctx, itemToUpdate).
			Return(errors.New("update category error"))

		err := service.Update(ctx, request)

		assert.Error(t, err, "update category error")
	})

	t.Run("returns nil when update succeeded", func(t *testing.T) {
		request := requests.ItemUpdateRequest{
			ID:          "1051",
			Date:        "2025-01-25",
			Formula:     "2+2",
			CategoryID:  "1100",
			Description: "food, wine and flowers",
		}

		itemFound := &models.Item{
			ID: 1051,
		}

		itemRepository.
			EXPECT().
			FindByID(ctx, 1051).
			Return(itemFound, nil)

		category := &models.Category{
			ID:   1100,
			Name: "Beverages",
		}

		categoryRepository.
			EXPECT().
			FindByID(ctx, 1100).
			Return(category, nil)

		itemToUpdate := &models.Item{
			ID:           1051,
			Date:         date(t, "2025-01-25"),
			Formula:      "2+2",
			Sum:          4,
			CategoryID:   1100,
			CategoryName: "Beverages",
			Description:  "food, wine and flowers",
		}

		itemRepository.
			EXPECT().
			Update(ctx, itemToUpdate).
			Return(nil)

		err := service.Update(ctx, request)

		assert.NilError(t, err)
	})
}

func TestItemService_Delete(t *testing.T) {
	controller := gomock.NewController(t)

	itemRepository := mocksforservices.NewMockItemRepository(controller)
	categoryRepository := mocksforservices.NewMockCategoryRepository(controller)

	service := services.NewItemService(itemRepository, categoryRepository)

	ctx := context.Background()

	t.Run("returns error when id is blank", func(t *testing.T) {
		err := service.Delete(ctx, "")

		assert.Error(t, err, "resource not found")
	})

	t.Run("returns error when id is zero", func(t *testing.T) {
		err := service.Delete(ctx, "0")

		assert.Error(t, err, "resource not found")
	})

	t.Run("returns error when delete failed", func(t *testing.T) {
		itemRepository.
			EXPECT().
			Delete(ctx, 2847).
			Return(errors.New("delete category error"))

		err := service.Delete(ctx, "2847")

		assert.Error(t, err, "delete category error")
	})

	t.Run("returns nil when delete succeeded", func(t *testing.T) {
		itemRepository.
			EXPECT().
			Delete(ctx, 2847).
			Return(nil)

		err := service.Delete(ctx, "2847")

		assert.NilError(t, err)
	})
}
