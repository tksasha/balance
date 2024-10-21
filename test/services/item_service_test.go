package services_test

import (
	"context"
	"errors"
	"testing"

	internalerrors "github.com/tksasha/balance/internal/errors"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/services"
	mockedrepositories "github.com/tksasha/balance/mocks/repositories"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestItemService_GetItems(t *testing.T) {
	itemsGetter := mockedrepositories.NewMockItemsGetter(gomock.NewController(t))

	service := services.NewItemServiceBuilder().WithItemsGetter(itemsGetter).Build()

	ctx := context.Background()

	currency, _ := models.GetDefaultCurrency()

	t.Run("when items getter returns an error it should return this error", func(t *testing.T) {
		itemsGetter.
			EXPECT().
			GetItems(ctx, currency).
			Return(nil, errors.New("get items error"))

		items, err := service.GetItems(ctx, currency)

		assert.Assert(t, items == nil)
		assert.Error(t, err, "get items error")
	})

	t.Run("when items getter does not return any error it should return items", func(t *testing.T) {
		itemsGetter.
			EXPECT().
			GetItems(ctx, currency).
			Return(models.Items{}, nil)

		items, err := service.GetItems(ctx, currency)

		assert.Assert(t, items != nil)
		assert.NilError(t, err)
	})
}

func TestItemService_GetItem(t *testing.T) {
	itemGetter := mockedrepositories.NewMockItemGetter(gomock.NewController(t))

	service := services.NewItemServiceBuilder().WithItemGetter(itemGetter).Build()

	ctx := context.Background()

	t.Run("when id is invalid it should return an error", func(t *testing.T) {
		_, err := service.GetItem(ctx, "abc")

		assert.ErrorIs(t, err, internalerrors.ErrNotFound)
	})

	t.Run("when id is zero it should return an error", func(t *testing.T) {
		item, err := service.GetItem(ctx, "0")

		assert.Assert(t, item == nil)
		assert.ErrorIs(t, err, internalerrors.ErrNotFound)
	})

	t.Run("when id is a negative number it should return an error", func(t *testing.T) {
		item, err := service.GetItem(ctx, "-1")

		assert.Assert(t, item == nil)
		assert.ErrorIs(t, err, internalerrors.ErrNotFound)
	})

	t.Run("when item getter returns an error it should return this error", func(t *testing.T) {
		itemGetter.
			EXPECT().
			GetItem(ctx, 1314).
			Return(nil, errors.New("get item error"))

		item, err := service.GetItem(ctx, "1314")

		assert.Assert(t, item == nil)
		assert.Error(t, err, "get item error")
	})

	t.Run("when item getter does not return any error it should return an item", func(t *testing.T) {
		itemGetter.
			EXPECT().
			GetItem(ctx, 1346).
			Return(&models.Item{}, nil)

		item, err := service.GetItem(ctx, "1346")

		assert.Assert(t, item != nil)
		assert.NilError(t, err)
	})
}

func TestItemService_UpdateItem(t *testing.T) {
	itemUpdater := mockedrepositories.NewMockItemUpdater(gomock.NewController(t))

	service := services.NewItemServiceBuilder().WithItemUpdater(itemUpdater).Build()

	ctx := context.Background()

	item := &models.Item{}

	t.Run("when item updater returns an error it should return an error", func(t *testing.T) {
		itemUpdater.
			EXPECT().
			UpdateItem(ctx, item).
			Return(errors.New("update item error"))

		err := service.UpdateItem(ctx, item)

		assert.Error(t, err, "update item error")
	})

	t.Run("when item updater does not return an error it should return nil", func(t *testing.T) {
		itemUpdater.
			EXPECT().
			UpdateItem(ctx, item).
			Return(nil)

		err := service.UpdateItem(ctx, item)

		assert.NilError(t, err)
	})
}

func TestItemService_CreateItem(t *testing.T) {
	itemCreator := mockedrepositories.NewMockItemCreator(gomock.NewController(t))

	service := services.NewItemServiceBuilder().WithItemCreator(itemCreator).Build()

	ctx := context.Background()

	item := &models.Item{}

	t.Run("when item creator returns an error it should return this error", func(t *testing.T) {
		itemCreator.
			EXPECT().
			CreateItem(ctx, item).
			Return(errors.New("create item error"))

		err := service.CreateItem(ctx, item)

		assert.Error(t, err, "create item error")
	})

	t.Run("when item creator doesn't return any error it should return nil", func(t *testing.T) {
		itemCreator.
			EXPECT().
			CreateItem(ctx, item).
			Return(nil)

		err := service.CreateItem(ctx, item)

		assert.NilError(t, err)
	})
}

func TestItemService_DeleteItem(t *testing.T) {
	itemDeleter := mockedrepositories.NewMockItemDeleter(gomock.NewController(t))

	service := services.NewItemServiceBuilder().WithItemDeleter(itemDeleter).Build()

	ctx := context.Background()

	item := &models.Item{}

	t.Run("when item deleter returns an error it should return this error", func(t *testing.T) {
		itemDeleter.
			EXPECT().
			DeleteItem(ctx, item).
			Return(errors.New("delete item error"))

		err := service.DeleteItem(ctx, item)

		assert.Error(t, err, "delete item error")
	})

	t.Run("when item deleter doesn't return any error it should return nil", func(t *testing.T) {
		itemDeleter.
			EXPECT().
			DeleteItem(ctx, item).
			Return(nil)

		err := service.DeleteItem(ctx, item)

		assert.NilError(t, err)
	})
}
