package services_test

import (
	"errors"
	"testing"

	"github.com/tksasha/balance/internal/apperrors"
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

	ctx := t.Context()

	t.Run("returns error when name is blank", func(t *testing.T) {
		request := requests.CategoryCreateRequest{
			Name:          "",
			Supercategory: "59",
		}

		err := service.Create(ctx, request)

		assert.Error(t, err, "name: is required")
	})

	t.Run("returns error when find category by name fails", func(t *testing.T) {
		request := requests.CategoryCreateRequest{
			Name: "Miscellaneous",
		}

		categoryRepository.
			EXPECT().
			FindByName(ctx, "Miscellaneous").
			Return(nil, errors.New("find category by name error"))

		err := service.Create(ctx, request)

		assert.Error(t, err, "find category by name error")
	})

	t.Run("returns error when name already exists", func(t *testing.T) {
		request := requests.CategoryCreateRequest{
			Name:          "Pharmaceutical",
			Supercategory: "59",
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

	t.Run("returns error when supercategory is invalid", func(t *testing.T) {
		request := requests.CategoryCreateRequest{
			Name:          "Miscellaneous",
			Supercategory: "abc",
		}

		categoryRepository.
			EXPECT().
			FindByName(ctx, "Miscellaneous").
			Return(nil, apperrors.ErrRecordNotFound)

		err := service.Create(ctx, request)

		assert.Error(t, err, "supercategory: is invalid")
	})

	t.Run("returns error when create fails", func(t *testing.T) {
		request := requests.CategoryCreateRequest{
			Name:          "Confectionery",
			Supercategory: "58",
		}

		category := &models.Category{
			Name:          "Confectionery",
			Supercategory: 58,
		}

		categoryRepository.
			EXPECT().
			FindByName(ctx, "Confectionery").
			Return(nil, apperrors.ErrRecordNotFound)

		categoryRepository.
			EXPECT().
			Create(ctx, category).
			Return(errors.New("create category error"))

		err := service.Create(ctx, request)

		assert.Error(t, err, "create category error")
	})

	t.Run("creates successfully", func(t *testing.T) {
		request := requests.CategoryCreateRequest{
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
			Return(nil, apperrors.ErrRecordNotFound)

		categoryRepository.
			EXPECT().
			Create(ctx, category).
			Return(nil)

		err := service.Create(ctx, request)

		assert.NilError(t, err)
	})
}

func TestCategoryService_FindByID(t *testing.T) {
	controller := gomock.NewController(t)

	categoryRepository := mocksforservices.NewMockCategoryRepository(controller)

	service := services.NewCategoryService(categoryRepository)

	ctx := t.Context()

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

	ctx := t.Context()

	t.Run("returns error when name is blank", func(t *testing.T) {
		category := &models.Category{
			Name: "",
		}

		err := service.Update(ctx, category)

		assert.Error(t, err, "name: is required")
	})

	t.Run("returns error when find by name fails", func(t *testing.T) {
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

	t.Run("returns error when name already exists", func(t *testing.T) {
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

	t.Run("returns error when update fails", func(t *testing.T) {
		category := &models.Category{
			Name: "Miscellaneous",
		}

		categoryRepository.
			EXPECT().
			FindByName(ctx, "Miscellaneous").
			Return(nil, apperrors.ErrRecordNotFound)

		categoryRepository.
			EXPECT().
			Update(ctx, category).
			Return(errors.New("update category error"))

		err := service.Update(ctx, category)

		assert.Error(t, err, "update category error")
	})

	t.Run("updates successfully", func(t *testing.T) {
		category := &models.Category{
			Name: "Stationery",
		}

		categoryRepository.
			EXPECT().
			FindByName(ctx, "Stationery").
			Return(nil, apperrors.ErrRecordNotFound)

		categoryRepository.
			EXPECT().
			Update(ctx, category).
			Return(nil)

		err := service.Update(ctx, category)

		assert.NilError(t, err)
	})
}
