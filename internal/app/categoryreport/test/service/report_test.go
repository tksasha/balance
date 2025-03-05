package service_test

import (
	"errors"
	"testing"

	"github.com/tksasha/balance/internal/app/categoryreport"
	"github.com/tksasha/balance/internal/app/categoryreport/service"
	"github.com/tksasha/balance/internal/app/categoryreport/test/mocks"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestReport(t *testing.T) { //nolint:funlen
	controller := gomock.NewController(t)

	repository := mocks.NewMockRepository(controller)

	service := service.New(repository)

	ctx := t.Context()

	request := categoryreport.Request{
		Year:  "2024",
		Month: "2",
	}

	filters := categoryreport.Filters{
		From: "2024-02-01",
		To:   "2024-02-29",
	}

	t.Run("when repository returns error", func(t *testing.T) {
		repository.EXPECT().Group(ctx, filters).Return(nil, errors.New("grouping error"))

		_, err := service.Report(ctx, request)

		assert.Error(t, err, "grouping error")
	})

	t.Run("when repository returns entities", func(t *testing.T) {
		entities := categoryreport.Entities{
			{CategoryName: "Food", CategorySlug: "food", Sum: 11.11, Supercategory: 1},
			{CategoryName: "Beverages", CategorySlug: "beverages", Sum: 22.22, Supercategory: 1},
			{CategoryName: "Rent", CategorySlug: "rent", Sum: 33.33, Supercategory: 2},
			{CategoryName: "Salary", CategorySlug: "salary", Sum: 44.44, Supercategory: 0},
			{CategoryName: "Stocks", CategorySlug: "stocks", Sum: 55.55, Supercategory: 0},
			{CategoryName: "Bonds", CategorySlug: "bonds", Sum: 66.66, Supercategory: 0},
		}

		repository.EXPECT().Group(ctx, filters).Return(entities, nil)

		actual, err := service.Report(ctx, request)

		expected := map[int]categoryreport.Entities{
			0: {
				{CategoryName: "Salary", CategorySlug: "salary", Sum: 44.44, Supercategory: 0},
				{CategoryName: "Stocks", CategorySlug: "stocks", Sum: 55.55, Supercategory: 0},
				{CategoryName: "Bonds", CategorySlug: "bonds", Sum: 66.66, Supercategory: 0},
			},
			1: {
				{CategoryName: "Food", CategorySlug: "food", Sum: 11.11, Supercategory: 1},
				{CategoryName: "Beverages", CategorySlug: "beverages", Sum: 22.22, Supercategory: 1},
			},
			2: {
				{CategoryName: "Rent", CategorySlug: "rent", Sum: 33.33, Supercategory: 2},
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
