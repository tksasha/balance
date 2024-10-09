package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/tksasha/balance/internal/decorators"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/services"
	mockedrepositories "github.com/tksasha/balance/mocks/repositories"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

//nolint:err113
func TestGetItems(t *testing.T) {
	controller := gomock.NewController(t)

	itemsGetter := mockedrepositories.NewMockItemsGetter(controller)

	service := services.NewGetItemsService(itemsGetter)

	ctx := context.Background()

	currency := models.Currency{}

	t.Run("when items getter returns an error it should return an error", func(t *testing.T) {
		itemsGetter.
			EXPECT().
			GetItems(ctx, currency).
			Return(nil, errors.New("get items error"))

		_, err := service.GetItems(ctx, currency)

		assert.Error(t, err, "get items error")
	})

	t.Run("when items getter does not return an error it should return decorated items", func(t *testing.T) {
		item := decorators.NewItemDecorator(
			models.NewItem(),
		)

		items := []*models.Item{item.Item}

		itemsGetter.
			EXPECT().
			GetItems(ctx, currency).
			Return(items, nil)

		res, err := service.GetItems(ctx, currency)

		assert.Assert(t, is.Contains(res.Items, item))
		assert.NilError(t, err)
	})
}
