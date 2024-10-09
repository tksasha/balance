package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/tksasha/balance/internal/decorators"
	internalerrors "github.com/tksasha/balance/internal/errors"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/services"
	mockedrepositories "github.com/tksasha/balance/mocks/repositories"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

//nolint:err113
func TestGetItem(t *testing.T) {
	controller := gomock.NewController(t)

	itemGetter := mockedrepositories.NewMockItemGetter(controller)

	service := services.NewGetItemService(itemGetter)

	ctx := context.Background()

	t.Run("when id is invalid it should return an error", func(t *testing.T) {
		_, err := service.GetItem(ctx, "abc")

		assert.ErrorIs(t, err, internalerrors.ErrNotFound)
	})

	t.Run("when item getter returns an error it should return this error", func(t *testing.T) {
		itemGetter.
			EXPECT().
			GetItem(ctx, 1314).
			Return(nil, errors.New("get item error"))

		_, err := service.GetItem(ctx, "1314")

		assert.Error(t, err, "get item error")
	})

	t.Run("when item getter does not return any error it should return an item", func(t *testing.T) {
		decorator := decorators.NewItemDecorator(
			models.NewItem(),
		)

		itemGetter.
			EXPECT().
			GetItem(ctx, 1346).
			Return(decorator.Item, nil)

		item, err := service.GetItem(ctx, "1346")

		assert.Equal(t, *item, *decorator)
		assert.NilError(t, err)
	})
}
