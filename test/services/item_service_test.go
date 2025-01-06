package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/services"
	mocksforservices "github.com/tksasha/balance/mocks/services"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestItemService_GetItems(t *testing.T) {
	itemRepository := mocksforservices.NewMockItemRepository(gomock.NewController(t))

	service := services.NewItemService(itemRepository)

	ctx := context.Background()

	t.Run("when items getter returns an error it should return this error", func(t *testing.T) {
		itemRepository.EXPECT().GetItems(ctx).Return(nil, errors.New("get items error"))

		items, err := service.GetItems(ctx)

		assert.Assert(t, items == nil)
		assert.Error(t, err, "get items error")
	})

	t.Run("when items getter does not return any error it should return items", func(t *testing.T) {
		itemRepository.EXPECT().GetItems(ctx).Return(models.Items{}, nil)

		items, err := service.GetItems(ctx)

		assert.Assert(t, items != nil)
		assert.NilError(t, err)
	})
}

func TestItemService_GetItem(t *testing.T) {
	itemRepository := mocksforservices.NewMockItemRepository(gomock.NewController(t))

	service := services.NewItemService(itemRepository)

	ctx := context.Background()

	t.Run("when id is invalid it should return an error", func(t *testing.T) {
		_, err := service.GetItem(ctx, "abc")

		assert.Error(t, err, "not found")
	})

	t.Run("when id is zero it should return an error", func(t *testing.T) {
		item, err := service.GetItem(ctx, "0")

		assert.Assert(t, item == nil)
		assert.Error(t, err, "not found")
	})

	t.Run("when id is a negative number it should return an error", func(t *testing.T) {
		item, err := service.GetItem(ctx, "-1")

		assert.Assert(t, item == nil)
		assert.Error(t, err, "not found")
	})

	t.Run("when GetItem returns an error it should return this error", func(t *testing.T) {
		itemRepository.EXPECT().GetItem(ctx, 1314).Return(nil, errors.New("get item error"))

		item, err := service.GetItem(ctx, "1314")

		assert.Assert(t, item == nil)
		assert.Error(t, err, "get item error")
	})

	t.Run("when GetItems does not return any error it should return an item", func(t *testing.T) {
		itemRepository.EXPECT().GetItem(ctx, 1346).Return(&models.Item{}, nil)

		item, err := service.GetItem(ctx, "1346")

		assert.Assert(t, item != nil)
		assert.NilError(t, err)
	})
}

func TestItemService_UpdateItem(t *testing.T) {
	itemRepository := mocksforservices.NewMockItemRepository(gomock.NewController(t))

	service := services.NewItemService(itemRepository)

	ctx := context.Background()

	item := &models.Item{}

	t.Run("when UpdateItem returns an error it should return an error", func(t *testing.T) {
		itemRepository.EXPECT().UpdateItem(ctx, item).Return(errors.New("update item error"))

		err := service.UpdateItem(ctx, item)

		assert.Error(t, err, "update item error")
	})

	t.Run("when UpdateItem does not return an error it should return nil", func(t *testing.T) {
		itemRepository.EXPECT().UpdateItem(ctx, item).Return(nil)

		err := service.UpdateItem(ctx, item)

		assert.NilError(t, err)
	})
}

func TestItemService_CreateItem(t *testing.T) {
	itemRepository := mocksforservices.NewMockItemRepository(gomock.NewController(t))

	service := services.NewItemService(itemRepository)

	ctx := context.Background()

	item := &models.Item{}

	t.Run("when CreateItem returns an error it should return this error", func(t *testing.T) {
		itemRepository.EXPECT().CreateItem(ctx, item).Return(errors.New("create item error"))

		err := service.CreateItem(ctx, item)

		assert.Error(t, err, "create item error")
	})

	t.Run("when CreateItem doesn't return any error it should return nil", func(t *testing.T) {
		itemRepository.EXPECT().CreateItem(ctx, item).Return(nil)

		err := service.CreateItem(ctx, item)

		assert.NilError(t, err)
	})
}

func TestItemService_DeleteItem(t *testing.T) {
	itemRepository := mocksforservices.NewMockItemRepository(gomock.NewController(t))

	service := services.NewItemService(itemRepository)

	ctx := context.Background()

	t.Run("when id is not a digit it should return NotFoundError", func(t *testing.T) {
		err := service.DeleteItem(ctx, "abc")

		assert.Error(t, err, "not found")
	})

	t.Run("when id is zero it should return NotFoundError", func(t *testing.T) {
		err := service.DeleteItem(ctx, "0")

		assert.Error(t, err, "not found")
	})

	t.Run("when id is a negative number it should return NotFoundError", func(t *testing.T) {
		err := service.DeleteItem(ctx, "-1")

		assert.Error(t, err, "not found")
	})

	t.Run("when DeleteItem returns an error it should return this error", func(t *testing.T) {
		itemRepository.EXPECT().DeleteItem(ctx, 1203).Return(errors.New("delete item error"))

		err := service.DeleteItem(ctx, "1203")

		assert.Error(t, err, "delete item error")
	})

	t.Run("when item deleter doesn't return any error it should return nil", func(t *testing.T) {
		itemRepository.EXPECT().DeleteItem(ctx, 1204).Return(nil)

		err := service.DeleteItem(ctx, "1204")

		assert.NilError(t, err)
	})
}
