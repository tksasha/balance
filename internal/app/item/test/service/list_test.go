package service_test

import (
	"errors"
	"slices"
	"testing"

	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/app/item/service"
	"github.com/tksasha/balance/internal/app/item/test/mocks"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestList(t *testing.T) {
	controller := gomock.NewController(t)

	itemRepository := mocks.NewMockRepository(controller)

	categoryRepository := mocks.NewMockCategoryRepository(controller)

	service := service.New(itemRepository, categoryRepository)

	ctx := t.Context()

	request := item.ListRequest{
		Year:  "2024",
		Month: "02",
	}

	filters := item.Filters{
		From: "2024-02-01",
		To:   "2024-02-29",
	}

	t.Run("returns error when find all by month fails", func(t *testing.T) {
		itemRepository.EXPECT().FindAll(ctx, filters).Return(nil, errors.New("find all by month error"))

		_, err := service.List(ctx, request)

		assert.Error(t, err, "find all by month error")
	})

	t.Run("returns items when find all by month is successful", func(t *testing.T) {
		foundItems := item.Items{}

		itemRepository.EXPECT().FindAll(ctx, filters).Return(foundItems, nil)

		items, err := service.List(ctx, request)

		assert.NilError(t, err)
		assert.Assert(t, slices.Equal(items, foundItems))
	})
}
