package service_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/tksasha/balance/internal/core/category"
	categorymocks "github.com/tksasha/balance/internal/core/category/test/mocks"
	"github.com/tksasha/balance/internal/core/common"
	"github.com/tksasha/balance/internal/core/common/tests"
	"github.com/tksasha/balance/internal/core/item"
	"github.com/tksasha/balance/internal/core/item/service"
	"github.com/tksasha/balance/internal/core/item/test/mocks"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestUpdate(t *testing.T) { //nolint:funlen,maintidx
	ctrl := gomock.NewController(t)

	itemRepository := mocks.NewMockRepository(ctrl)
	categoryRepository := categorymocks.NewMockRepository(ctrl)

	service := service.New(common.NewBaseService(), itemRepository, categoryRepository)

	ctx := t.Context()

	t.Run("returns error when id is invalid", func(t *testing.T) {
		request := item.UpdateRequest{
			ID: "abc",
		}

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "resource not found")
	})

	t.Run("returns error when item doesn't exist", func(t *testing.T) {
		request := item.UpdateRequest{
			ID: "1027",
		}

		itemRepository.
			EXPECT().
			FindByID(ctx, 1027).
			Return(nil, common.ErrRecordNotFound)

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "resource not found")
	})

	t.Run("returns error when find by id failed", func(t *testing.T) {
		request := item.UpdateRequest{
			ID: "1049",
		}

		itemRepository.
			EXPECT().
			FindByID(ctx, 1049).
			Return(nil, errors.New("find item by id error"))

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "find item by id error")
	})

	t.Run("returns error when date is blank", func(t *testing.T) {
		request := item.UpdateRequest{
			ID:         "1051",
			Date:       "",
			Formula:    "2+2",
			CategoryID: "1052",
		}

		item := &item.Item{}

		itemRepository.
			EXPECT().
			FindByID(ctx, 1051).
			Return(item, nil)

		category := &category.Category{
			ID:   1052,
			Name: "Tools",
		}

		categoryRepository.
			EXPECT().
			FindByID(ctx, 1052).
			Return(category, nil)

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "date: is required")
	})

	t.Run("returns error when date is invalid", func(t *testing.T) {
		request := item.UpdateRequest{
			ID:         "1051",
			Date:       "abc",
			Formula:    "2+2",
			CategoryID: "1052",
		}

		item := &item.Item{}

		itemRepository.
			EXPECT().
			FindByID(ctx, 1051).
			Return(item, nil)

		category := &category.Category{
			ID:   1052,
			Name: "Tools",
		}

		categoryRepository.
			EXPECT().
			FindByID(ctx, 1052).
			Return(category, nil)

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "date: is invalid")
	})

	t.Run("returns error when formula is blank", func(t *testing.T) {
		request := item.UpdateRequest{
			ID:         "1051",
			Date:       "2025-01-25",
			Formula:    "",
			CategoryID: "1052",
		}

		item := &item.Item{}

		itemRepository.
			EXPECT().
			FindByID(ctx, 1051).
			Return(item, nil)

		category := &category.Category{
			ID:   1052,
			Name: "Tools",
		}

		categoryRepository.
			EXPECT().
			FindByID(ctx, 1052).
			Return(category, nil)

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "sum: is required")
	})

	t.Run("returns error when formula is invalid", func(t *testing.T) {
		request := item.UpdateRequest{
			ID:         "1051",
			Date:       "2025-01-25",
			Formula:    "abc",
			CategoryID: "1052",
		}

		item := &item.Item{}

		itemRepository.
			EXPECT().
			FindByID(ctx, 1051).
			Return(item, nil)

		category := &category.Category{
			ID:   1052,
			Name: "Tools",
		}

		categoryRepository.
			EXPECT().
			FindByID(ctx, 1052).
			Return(category, nil)

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "sum: is invalid")
	})

	t.Run("returns error when category_id is blank", func(t *testing.T) {
		request := item.UpdateRequest{
			ID:         "1051",
			Date:       "2025-01-25",
			Formula:    "2+2",
			CategoryID: "",
		}

		item := &item.Item{}

		itemRepository.
			EXPECT().
			FindByID(ctx, 1051).
			Return(item, nil)

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "category: is required")
	})

	t.Run("returns error when category_id is invalid", func(t *testing.T) {
		request := item.UpdateRequest{
			ID:         "1051",
			Date:       "2025-01-25",
			Formula:    "2+2",
			CategoryID: "abc",
		}

		item := &item.Item{}

		itemRepository.
			EXPECT().
			FindByID(ctx, 1051).
			Return(item, nil)

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "category: is invalid")
	})

	t.Run("returns error when find category by id failed", func(t *testing.T) {
		request := item.UpdateRequest{
			ID:         "1051",
			Date:       "2025-01-25",
			Formula:    "2+2",
			CategoryID: "1100",
		}

		item := &item.Item{}

		itemRepository.
			EXPECT().
			FindByID(ctx, 1051).
			Return(item, nil)

		categoryRepository.
			EXPECT().
			FindByID(ctx, 1100).
			Return(nil, errors.New("find category by id error"))

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "find category by id error")
	})

	t.Run("returns error when update failed", func(t *testing.T) {
		request := item.UpdateRequest{
			ID:          "1051",
			Date:        "2025-01-25",
			Formula:     "2+2",
			CategoryID:  "1100",
			Description: "food, wine and flowers",
		}

		itemFound := &item.Item{
			ID: 1051,
		}

		itemRepository.
			EXPECT().
			FindByID(ctx, 1051).
			Return(itemFound, nil)

		category := &category.Category{
			ID:   1100,
			Name: "Beverages",
		}

		categoryRepository.
			EXPECT().
			FindByID(ctx, 1100).
			Return(category, nil)

		itemToUpdate := &item.Item{
			ID:           1051,
			Date:         tests.Date(t, "2025-01-25"),
			Formula:      "2+2",
			Sum:          4,
			CategoryID:   1100,
			CategoryName: sql.NullString{String: "Beverages"},
			Description:  "food, wine and flowers",
		}

		itemRepository.
			EXPECT().
			Update(ctx, itemToUpdate).
			Return(errors.New("update category error"))

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "update category error")
	})

	t.Run("returns nil when update succeeded", func(t *testing.T) {
		request := item.UpdateRequest{
			ID:          "1051",
			Date:        "2025-01-25",
			Formula:     "2+2",
			CategoryID:  "1100",
			Description: "food, wine and flowers",
		}

		itemFound := &item.Item{
			ID: 1051,
		}

		itemRepository.
			EXPECT().
			FindByID(ctx, 1051).
			Return(itemFound, nil)

		category := &category.Category{
			ID:   1100,
			Name: "Beverages",
		}

		categoryRepository.
			EXPECT().
			FindByID(ctx, 1100).
			Return(category, nil)

		itemToUpdate := &item.Item{
			ID:           1051,
			Date:         tests.Date(t, "2025-01-25"),
			Formula:      "2+2",
			Sum:          4,
			CategoryID:   1100,
			CategoryName: sql.NullString{String: "Beverages"},
			Description:  "food, wine and flowers",
		}

		itemRepository.
			EXPECT().
			Update(ctx, itemToUpdate).
			Return(nil)

		_, err := service.Update(ctx, request)

		assert.NilError(t, err)
	})
}
