package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/services"
	mockedrepositories "github.com/tksasha/balance/test/mocks/repositories"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

//nolint:err113
func TestUpdateItem(t *testing.T) {
	controller := gomock.NewController(t)

	itemUpdater := mockedrepositories.NewMockItemUpdater(controller)

	service := services.NewUpdateItemService(itemUpdater)

	ctx := context.Background()

	item := models.NewItem()

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
