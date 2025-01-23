package services_test

import (
	"context"
	"errors"
	"testing"

	internalerrors "github.com/tksasha/balance/internal/errors"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/requests"
	"github.com/tksasha/balance/internal/services"
	mocksforservices "github.com/tksasha/balance/mocks/services"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestCreateCategoryService(t *testing.T) { //nolint:funlen
	controller := gomock.NewController(t)

	categoryRepository := mocksforservices.NewMockCategoryRepository(controller)

	service := services.NewCreateCategoryService(categoryRepository)

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
