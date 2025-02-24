package service_test

import (
	"errors"
	"testing"

	"github.com/tksasha/balance/internal/app/item/service"
	"github.com/tksasha/balance/internal/app/item/test/mocks"
	"github.com/tksasha/balance/internal/common"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)

	itemRepository := mocks.NewMockRepository(ctrl)
	categoryRepository := mocks.NewMockCategoryRepository(ctrl)

	service := service.New(common.NewBaseService(), itemRepository, categoryRepository)

	ctx := t.Context()

	t.Run("returns error when id is blank", func(t *testing.T) {
		err := service.Delete(ctx, "")

		assert.Error(t, err, "resource not found")
	})

	t.Run("returns error when id is zero", func(t *testing.T) {
		err := service.Delete(ctx, "0")

		assert.Error(t, err, "resource not found")
	})

	t.Run("returns error when item is not found", func(t *testing.T) {
		itemRepository.EXPECT().Delete(ctx, 1038).Return(common.ErrRecordNotFound)

		err := service.Delete(ctx, "1038")

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
