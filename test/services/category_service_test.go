package services_test

import (
	"context"
	"errors"
	"slices"
	"testing"

	internalerrors "github.com/tksasha/balance/internal/errors"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/requests"
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
		request := requests.CreateCategoryRequest{
			Name: "",
		}

		err := service.Create(ctx, request)

		assert.Error(t, err, "name: is required")
	})

	t.Run("when find category by name returns error, it should return error", func(t *testing.T) {
		request := requests.CreateCategoryRequest{
			Name: "Miscellaneous",
		}

		categoryRepository.
			EXPECT().
			FindByName(ctx, "Miscellaneous").
			Return(nil, errors.New("find category by name error"))

		err := service.Create(ctx, request)

		assert.Error(t, err, "find category by name error")
	})

	t.Run("when category name is not unique, it should return error", func(t *testing.T) {
		request := requests.CreateCategoryRequest{
			Name: "Pharmaceutical",
		}

		category := &models.Category{
			ID:   1100,
			Name: "Pharmaceutical",
		}

		categoryRepository.
			EXPECT().
			FindByName(ctx, "Pharmaceutical").
			Return(category, nil)

		err := service.Create(ctx, request)

		assert.Error(t, err, "name: already exists")
	})

	t.Run("when supercategory is invalid, it should return error", func(t *testing.T) {
		request := requests.CreateCategoryRequest{
			Name:          "Miscellaneous",
			Supercategory: "abc",
		}

		categoryRepository.
			EXPECT().
			FindByName(ctx, "Miscellaneous").
			Return(nil, nil)

		err := service.Create(ctx, request)

		assert.Error(t, err, "supercategory: is invalid")
	})

	t.Run("when create category returns error, it should return error", func(t *testing.T) {
		request := requests.CreateCategoryRequest{
			Name: "Confectionery",
		}

		category := &models.Category{
			Name: "Confectionery",
		}

		categoryRepository.
			EXPECT().
			FindByName(ctx, "Confectionery").
			Return(nil, internalerrors.ErrRecordNotFound)

		categoryRepository.
			EXPECT().
			Create(ctx, category).
			Return(errors.New("create category error"))

		err := service.Create(ctx, request)

		assert.Error(t, err, "create category error")
	})

	t.Run("when create category does not return error, it should return nil", func(t *testing.T) {
		request := requests.CreateCategoryRequest{
			Name:          "Haberdashery",
			Income:        "true",
			Visible:       "true",
			Supercategory: "23",
		}

		category := &models.Category{
			Name:          "Haberdashery",
			Income:        true,
			Visible:       true,
			Supercategory: 23,
		}

		categoryRepository.
			EXPECT().
			FindByName(ctx, "Haberdashery").
			Return(nil, internalerrors.ErrRecordNotFound)

		categoryRepository.
			EXPECT().
			Create(ctx, category).
			Return(nil)

		err := service.Create(ctx, request)

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

func TestCategoryService_Update(t *testing.T) { //nolint:funlen
	controller := gomock.NewController(t)

	categoryRepository := mocksforservices.NewMockCategoryRepository(controller)

	service := services.NewCategoryService(categoryRepository)

	ctx := context.Background()

	t.Run("when category name is blank, it should return error", func(t *testing.T) {
		category := &models.Category{
			Name: "",
		}

		err := service.Update(ctx, category)

		assert.Error(t, err, "name: is required")
	})

	t.Run("when find category by name returns error, it should return error", func(t *testing.T) {
		category := &models.Category{
			Name: "Entrepreneurship",
		}

		categoryRepository.
			EXPECT().
			FindByName(ctx, "Entrepreneurship").
			Return(nil, errors.New("find category by name error"))

		err := service.Update(ctx, category)

		assert.Error(t, err, "find category by name error")
	})

	t.Run("when category already exists, it should return error", func(t *testing.T) {
		categoryToUpdate := &models.Category{
			ID:   1030,
			Name: "Beverages",
		}

		categoryFound := &models.Category{
			ID:   1029,
			Name: "Beverages",
		}

		categoryRepository.
			EXPECT().
			FindByName(ctx, "Beverages").
			Return(categoryFound, nil)

		err := service.Update(ctx, categoryToUpdate)

		assert.Error(t, err, "name: already exists")
	})

	t.Run("when update category returns error, it should return error", func(t *testing.T) {
		category := &models.Category{
			Name: "Miscellaneous",
		}

		categoryRepository.
			EXPECT().
			FindByName(ctx, "Miscellaneous").
			Return(nil, internalerrors.ErrRecordNotFound)

		categoryRepository.
			EXPECT().
			Update(ctx, category).
			Return(errors.New("update category error"))

		err := service.Update(ctx, category)

		assert.Error(t, err, "update category error")
	})

	t.Run("when update category does not return error, it should return nil", func(t *testing.T) {
		category := &models.Category{
			Name: "Stationery",
		}

		categoryRepository.
			EXPECT().
			FindByName(ctx, "Stationery").
			Return(nil, internalerrors.ErrRecordNotFound)

		categoryRepository.
			EXPECT().
			Update(ctx, category).
			Return(nil)

		err := service.Update(ctx, category)

		assert.NilError(t, err)
	})
}

func TestCategoryService_Delete(t *testing.T) {
	controller := gomock.NewController(t)

	categoryRepository := mocksforservices.NewMockCategoryRepository(controller)

	service := services.NewCategoryService(categoryRepository)

	ctx := context.Background()

	category := &models.Category{}

	t.Run("when delete category returns error, it should return error", func(t *testing.T) {
		categoryRepository.
			EXPECT().
			Delete(ctx, category).
			Return(errors.New("delete category error"))

		err := service.Delete(ctx, category)

		assert.Error(t, err, "delete category error")
	})

	t.Run("when delete category does not return error, it should return nil", func(t *testing.T) {
		categoryRepository.
			EXPECT().
			Delete(ctx, category).
			Return(nil)

		err := service.Delete(ctx, category)

		assert.NilError(t, err)
	})
}
