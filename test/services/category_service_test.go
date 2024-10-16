package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/services"
	mockedrepositories "github.com/tksasha/balance/mocks/repositories"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestGetCategoriesService_GetCategories(t *testing.T) {
	controller := gomock.NewController(t)

	categoriesGetter := mockedrepositories.NewMockCategoriesGetter(controller)

	service := services.NewCategoryService(categoriesGetter)

	ctx := context.Background()

	currency, _ := models.GetDefaultCurrency()

	t.Run("when categoriesGetter returns an error it should return this error", func(t *testing.T) {
		categoriesGetter.
			EXPECT().
			GetCategories(ctx, currency).
			Return(nil, errors.New("get categories error"))

		categories, err := service.GetCategories(ctx, currency)

		assert.Assert(t, categories == nil)
		assert.Error(t, err, "get categories error")
	})

	t.Run("when categoriesGetter doesn't return any error it should return categories", func(t *testing.T) {
		categoriesGetter.
			EXPECT().
			GetCategories(ctx, currency).
			Return(models.Categories{}, nil)

		categories, err := service.GetCategories(ctx, currency)

		assert.Assert(t, categories != nil)
		assert.NilError(t, err)
	})
}
