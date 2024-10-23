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
	categoryRepository := mockedrepositories.NewMockCategoryRepository(
		gomock.NewController(t),
	)

	service := services.NewCategoryService(categoryRepository)

	ctx := context.Background()

	t.Run("when GetCategories returns an error it should return this error", func(t *testing.T) {
		categoryRepository.EXPECT().GetCategories(ctx).Return(nil, errors.New("get categories error"))

		categories, err := service.GetCategories(ctx)

		assert.Assert(t, categories == nil)
		assert.Error(t, err, "get categories error")
	})

	t.Run("when GetCategories doesn't return any error it should return categories", func(t *testing.T) {
		categoryRepository.EXPECT().GetCategories(ctx).Return(models.Categories{}, nil)

		categories, err := service.GetCategories(ctx)

		assert.Assert(t, categories != nil)
		assert.NilError(t, err)
	})
}
