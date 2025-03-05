package service_test

import (
	"errors"
	"testing"

	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/app/category/service"
	"github.com/tksasha/balance/internal/app/category/test/mocks"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestReport(t *testing.T) { //nolint:funlen
	controller := gomock.NewController(t)

	repository := mocks.NewMockRepository(controller)

	service := service.New(repository)

	ctx := t.Context()

	request := category.Request{
		Year:  "2024",
		Month: "2",
	}

	filters := category.Filters{
		From: "2024-02-01",
		To:   "2024-02-29",
	}

	t.Run("when repository returns error", func(t *testing.T) {
		repository.EXPECT().FindAllByFilters(ctx, filters).Return(nil, errors.New("grouping error"))

		_, err := service.GroupedList(ctx, request)

		assert.Error(t, err, "grouping error")
	})

	t.Run("when repository returns entities", func(t *testing.T) {
		categories := category.Categories{
			{Name: "Food", Slug: "food", Sum: 11.11, Supercategory: 1},
			{Name: "Beverages", Slug: "beverages", Sum: 22.22, Supercategory: 1},
			{Name: "Rent", Slug: "rent", Sum: 33.33, Supercategory: 2},
			{Name: "Salary", Slug: "salary", Sum: 44.44, Supercategory: 0},
			{Name: "Stocks", Slug: "stocks", Sum: 55.55, Supercategory: 0},
			{Name: "Bonds", Slug: "bonds", Sum: 66.66, Supercategory: 0},
		}

		repository.EXPECT().FindAllByFilters(ctx, filters).Return(categories, nil)

		actual, err := service.GroupedList(ctx, request)

		expected := category.GroupedCategories{
			0: {
				{Name: "Salary", Slug: "salary", Sum: 44.44, Supercategory: 0},
				{Name: "Stocks", Slug: "stocks", Sum: 55.55, Supercategory: 0},
				{Name: "Bonds", Slug: "bonds", Sum: 66.66, Supercategory: 0},
			},
			1: {
				{Name: "Food", Slug: "food", Sum: 11.11, Supercategory: 1},
				{Name: "Beverages", Slug: "beverages", Sum: 22.22, Supercategory: 1},
			},
			2: {
				{Name: "Rent", Slug: "rent", Sum: 33.33, Supercategory: 2},
			},
		}

		assert.NilError(t, err)

		for supercategory := range expected {
			for idx := range expected[supercategory] {
				assert.Equal(t, *actual[supercategory][idx], *expected[supercategory][idx])
			}
		}
	})
}
