package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/services"
	mocksforservices "github.com/tksasha/balance/mocks/services"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestCategoryService_Create(t *testing.T) { //nolint:funlen
	controller := gomock.NewController(t)

	categoryRepository := mocksforservices.NewMockCategoryRepository(controller)

	service := services.NewCategoryService(categoryRepository)

	ctx := context.Background()

	t.Run("when category name is empty, it should return error", func(t *testing.T) {
		category := &models.Category{
			Name: "",
		}

		categoryRepository.
			EXPECT().
			FindByName(ctx, "").
			Return(nil, errors.New("record not found"))

		err := service.Create(ctx, category)

		assert.Error(t, err, "name: is required")
	})

	t.Run("when category name is not unique, it should return error", func(t *testing.T) {
		category := &models.Category{
			Name: "Food",
		}

		categoryRepository.
			EXPECT().
			FindByName(ctx, "Food").
			Return(category, nil)

		err := service.Create(ctx, category)

		assert.Error(t, err, "name: already exists")
	})

	t.Run("when create category returns error, it should return error", func(t *testing.T) {
		category := &models.Category{
			Name: "Drinks",
		}

		categoryRepository.
			EXPECT().
			FindByName(ctx, "Drinks").
			Return(nil, errors.New("record not found"))

		categoryRepository.
			EXPECT().
			Create(ctx, category).
			Return(errors.New("create category error"))

		err := service.Create(ctx, category)

		assert.Error(t, err, "create category error")
	})

	t.Run("when create category does not return error, it should return nil", func(t *testing.T) {
		category := &models.Category{
			Name: "Products",
		}

		categoryRepository.
			EXPECT().
			FindByName(ctx, "Products").
			Return(nil, errors.New("record not found"))

		categoryRepository.
			EXPECT().
			Create(ctx, category).
			Return(nil)

		err := service.Create(ctx, category)

		assert.NilError(t, err)
	})
}
