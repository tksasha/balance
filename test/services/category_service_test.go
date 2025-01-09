package services_test

import (
	"context"
	"errors"
	"slices"
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

func TestCategoryService_GetAll(t *testing.T) {
	controller := gomock.NewController(t)

	categoryRepository := mocksforservices.NewMockCategoryRepository(controller)

	service := services.NewCategoryService(categoryRepository)

	ctx := context.Background()

	t.Run("when get all categories returns error, it should return error", func(t *testing.T) {
		categoryRepository.
			EXPECT().
			GetAll(ctx).
			Return(nil, errors.New("get all categories error"))

		_, err := service.GetAll(ctx)

		assert.Error(t, err, "get all categories error")
	})

	t.Run("when get all categories does not return error, it should not return error", func(t *testing.T) {
		expected := models.Categories{}

		categoryRepository.
			EXPECT().
			GetAll(ctx).
			Return(expected, nil)

		result, err := service.GetAll(ctx)

		assert.Assert(t, slices.Equal(result, expected))
		assert.NilError(t, err)
	})
}

func TestCategoryService_FindByID(t *testing.T) {
	controller := gomock.NewController(t)

	categoryRepository := mocksforservices.NewMockCategoryRepository(controller)

	service := services.NewCategoryService(categoryRepository)

	ctx := context.Background()

	t.Run("when find category by id returns error, it should return error", func(t *testing.T) {
		categoryRepository.
			EXPECT().
			FindByID(ctx, 1531).
			Return(nil, errors.New("find category by id error"))

		_, err := service.FindByID(ctx, 1531)

		assert.Error(t, err, "find category by id error")
	})

	t.Run("when find category by id returns category, it should return category", func(t *testing.T) {
		expected := &models.Category{}

		categoryRepository.
			EXPECT().
			FindByID(ctx, 1536).
			Return(expected, nil)

		result, err := service.FindByID(ctx, 1536)

		assert.Equal(t, result, expected)
		assert.NilError(t, err)
	})
}
