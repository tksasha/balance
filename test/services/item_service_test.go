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

type itemServiceTestContext struct {
	service     *services.ItemService
	itemsGetter *mockedrepositories.MockItemsGetter
	itemGetter  *mockedrepositories.MockItemGetter
	itemUpdater *mockedrepositories.MockItemUpdater
	itemCreator *mockedrepositories.MockItemCreator
	itemDeleter *mockedrepositories.MockItemDeleter
}

func setupItemService(t *testing.T) *itemServiceTestContext {
	t.Helper()

	controller := gomock.NewController(t)

	itemsGetter := mockedrepositories.NewMockItemsGetter(controller)
	itemGetter := mockedrepositories.NewMockItemGetter(controller)
	itemUpdater := mockedrepositories.NewMockItemUpdater(controller)
	itemCreator := mockedrepositories.NewMockItemCreator(controller)
	itemDeleter := mockedrepositories.NewMockItemDeleter(controller)

	service := services.NewItemService(itemsGetter, itemGetter, itemUpdater, itemCreator, itemDeleter)

	return &itemServiceTestContext{
		service, itemsGetter, itemGetter, itemUpdater, itemCreator, itemDeleter,
	}
}

func TestItemService_GetItems(t *testing.T) {
	setup := setupItemService(t)

	itemsGetter, service := setup.itemsGetter, setup.service

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
	setup := setupItemService(t)

	service, itemGetter := setup.service, setup.itemGetter

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
	setup := setupItemService(t)

	service, itemUpdater := setup.service, setup.itemUpdater

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
	setup := setupItemService(t)

	service, itemCreator := setup.service, setup.itemCreator

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
	setup := setupItemService(t)

	service, itemDeleter := setup.service, setup.itemDeleter

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
