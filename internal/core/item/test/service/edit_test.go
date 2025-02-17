package service_test

import (
	"testing"

	categorymocks "github.com/tksasha/balance/internal/core/category/test/mocks"
	"github.com/tksasha/balance/internal/core/common"
	"github.com/tksasha/balance/internal/core/item"
	"github.com/tksasha/balance/internal/core/item/service"
	"github.com/tksasha/balance/internal/core/item/test/mocks"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestEdit(t *testing.T) {
	controller := gomock.NewController(t)

	itemRepository := mocks.NewMockRepository(controller)

	categoryRepository := categorymocks.NewMockRepository(controller)

	service := service.New(itemRepository, categoryRepository)

	ctx := t.Context()

	t.Run("returns error when id is invalid", func(t *testing.T) {
		_, err := service.Edit(ctx, "abc")

		assert.Error(t, err, "resource not found")
	})

	t.Run("returns error when item is not found", func(t *testing.T) {
		itemRepository.EXPECT().FindByID(ctx, 1101).Return(nil, common.ErrRecordNotFound)

		_, err := service.Edit(ctx, "1101")

		assert.Error(t, err, "resource not found")
	})

	t.Run("returns item when it is found", func(t *testing.T) {
		foundItem := &item.Item{}

		itemRepository.EXPECT().FindByID(ctx, 1102).Return(foundItem, nil)

		item, err := service.Edit(ctx, "1102")

		assert.NilError(t, err)
		assert.Equal(t, item, foundItem)
	})
}
