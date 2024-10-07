package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/services"
	mockedrepositories "github.com/tksasha/balance/test/mocks/repositories"
	mockedservices "github.com/tksasha/balance/test/mocks/services"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

//nolint:err113
func TestUpdateItem(t *testing.T) {
	controller := gomock.NewController(t)

	itemGetter := mockedservices.NewMockItemGetter(controller)

	itemUpdater := mockedrepositories.NewMockItemUpdater(controller)

	service := services.NewUpdateItemService(itemGetter, itemUpdater)

	ctx := context.Background()

	item := models.NewItem()

	t.Run("when item getter returns an error it should return an error", func(t *testing.T) {
		itemGetter.
			EXPECT().
			GetItem(ctx, "1440").
			Return(nil, errors.New("get item error"))

		err := service.UpdateItem(ctx, "1440")

		assert.Error(t, err, "get item error")
	})

	t.Run("when item updater returns an error it should return an error", func(t *testing.T) {
		itemGetter.
			EXPECT().
			GetItem(ctx, "1444").
			Return(item, nil)

		itemUpdater.
			EXPECT().
			UpdateItem(ctx, item).
			Return(errors.New("update item error"))

		err := service.UpdateItem(ctx, "1444")

		assert.Error(t, err, "update item error")
	})

	t.Run("when item updater does not return an error it should return nil", func(t *testing.T) {
		itemGetter.
			EXPECT().
			GetItem(ctx, "1445").
			Return(item, nil)

		itemUpdater.
			EXPECT().
			UpdateItem(ctx, item).
			Return(nil)

		err := service.UpdateItem(ctx, "1445")

		assert.NilError(t, err)
	})
}
