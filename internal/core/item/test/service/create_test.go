package service_test

import (
	"errors"
	"testing"

	"github.com/tksasha/balance/internal/common/tests"
	"github.com/tksasha/balance/internal/core/category"
	categorymocks "github.com/tksasha/balance/internal/core/category/test/mocks"
	"github.com/tksasha/balance/internal/core/item"
	"github.com/tksasha/balance/internal/core/item/service"
	"github.com/tksasha/balance/internal/core/item/test/mocks"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestCreate(t *testing.T) { //nolint:funlen
	ctrl := gomock.NewController(t)

	itemRepository := mocks.NewMockRepository(ctrl)
	categoryRepository := categorymocks.NewMockRepository(ctrl)

	service := service.New(itemRepository, categoryRepository)

	ctx := t.Context()

	t.Run("returns error when date is blank", func(t *testing.T) {
		request := item.CreateRequest{
			Date:       "",
			Formula:    "2+2",
			CategoryID: "1321",
		}

		category := &category.Category{
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
		request := item.CreateRequest{
			Date:       "abc",
			Formula:    "2+2",
			CategoryID: "1320",
		}

		category := &category.Category{
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
		request := item.CreateRequest{
			Date:       "2025-01-23",
			Formula:    "",
			CategoryID: "1318",
		}

		category := &category.Category{
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
		request := item.CreateRequest{
			Date:       "2025-01-23",
			Formula:    "abc",
			CategoryID: "1315",
		}

		category := &category.Category{
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
		request := item.CreateRequest{
			Date:       "2025-01-23",
			Formula:    "42.69+69.42",
			CategoryID: "",
		}

		_, err := service.Create(ctx, request)

		assert.Error(t, err, "category: is required")
	})

	t.Run("returns error when category_id is invalid", func(t *testing.T) {
		request := item.CreateRequest{
			Date:       "2025-01-23",
			Formula:    "42.69+69.42",
			CategoryID: "abc",
		}

		_, err := service.Create(ctx, request)

		assert.Error(t, err, "category: is invalid")
	})

	t.Run("returns error when find by id failed", func(t *testing.T) {
		request := item.CreateRequest{
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
		request := item.CreateRequest{
			Date:        "2025-01-23",
			Formula:     "42.69+69.42",
			CategoryID:  "1244",
			Description: "health, beauty & wellness",
		}

		category := &category.Category{
			ID:   1244,
			Name: "Entrepreneurship",
		}

		categoryRepository.
			EXPECT().
			FindByID(ctx, 1244).
			Return(category, nil)

		item := &item.Item{
			Date:         tests.Date(t, "2025-01-23"),
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
		request := item.CreateRequest{
			Date:        "2025-01-23",
			Formula:     "42.69+69.42",
			CategoryID:  "1307",
			Description: "arts, crafts & hobbies",
		}

		category := &category.Category{
			ID:   1244,
			Name: "Telecommunications",
		}

		categoryRepository.
			EXPECT().
			FindByID(ctx, 1307).
			Return(category, nil)

		item := &item.Item{
			Date:         tests.Date(t, "2025-01-23"),
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
