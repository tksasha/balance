package service_test

import (
	"errors"
	"slices"
	"testing"

	"github.com/tksasha/balance/internal/common"
	categorymocks "github.com/tksasha/balance/internal/core/category/test/mocks"
	"github.com/tksasha/balance/internal/core/item"
	"github.com/tksasha/balance/internal/core/item/service"
	"github.com/tksasha/balance/internal/core/item/test/mocks"
	"github.com/tksasha/month"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestList(t *testing.T) {
	controller := gomock.NewController(t)

	itemRepository := mocks.NewMockRepository(controller)

	categoryRepository := categorymocks.NewMockRepository(controller)

	service := service.New(common.NewBaseService(), itemRepository, categoryRepository)

	ctx := t.Context()

	month := month.New("2024", "02")

	request := item.ListRequest{
		Year:  "2024",
		Month: "02",
	}

	t.Run("returns error when find all by month fails", func(t *testing.T) {
		itemRepository.EXPECT().FindAllByMonth(ctx, month).Return(nil, errors.New("find all by month error"))

		_, err := service.List(ctx, request)

		assert.Error(t, err, "find all by month error")
	})

	t.Run("returns items when find all by month is successful", func(t *testing.T) {
		foundItems := item.Items{}

		itemRepository.EXPECT().FindAllByMonth(ctx, month).Return(foundItems, nil)

		items, err := service.List(ctx, request)

		assert.NilError(t, err)
		assert.Assert(t, slices.Equal(items, foundItems))
	})
}
