package service_test

import (
	"errors"
	"testing"

	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/category/service"
	"github.com/tksasha/balance/internal/core/category/test/mocks"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestCreate(t *testing.T) { //nolint:funlen
	ctrl := gomock.NewController(t)

	repository := mocks.NewMockRepository(ctrl)

	service := service.New(common.NewBaseService(), repository)

	ctx := t.Context()

	t.Run("returns error when name is blank", func(t *testing.T) {
		request := category.CreateRequest{
			Name:          "",
			Supercategory: "59",
		}

		_, err := service.Create(ctx, request)

		assert.Error(t, err, "name: is required")
	})

	t.Run("returns error when find category by name fails", func(t *testing.T) {
		request := category.CreateRequest{
			Name: "Miscellaneous",
		}

		repository.
			EXPECT().
			FindByName(ctx, "Miscellaneous").
			Return(nil, errors.New("find category by name error"))

		_, err := service.Create(ctx, request)

		assert.Error(t, err, "find category by name error")
	})

	t.Run("returns error when name already exists", func(t *testing.T) {
		request := category.CreateRequest{
			Name:          "Pharmaceutical",
			Supercategory: "59",
		}

		category := &category.Category{
			ID:   1100,
			Name: "Pharmaceutical",
		}

		repository.
			EXPECT().
			FindByName(ctx, "Pharmaceutical").
			Return(category, nil)

		_, err := service.Create(ctx, request)

		assert.Error(t, err, "name: already exists")
	})

	t.Run("returns error when supercategory is invalid", func(t *testing.T) {
		request := category.CreateRequest{
			Name:          "Miscellaneous",
			Supercategory: "abc",
		}

		repository.
			EXPECT().
			FindByName(ctx, "Miscellaneous").
			Return(nil, common.ErrRecordNotFound)

		_, err := service.Create(ctx, request)

		assert.Error(t, err, "supercategory: is invalid")
	})

	t.Run("returns error when create fails", func(t *testing.T) {
		request := category.CreateRequest{
			Name:          "Confectionery",
			Supercategory: "58",
		}

		category := &category.Category{
			Name:          "Confectionery",
			Supercategory: 58,
		}

		repository.
			EXPECT().
			FindByName(ctx, "Confectionery").
			Return(nil, common.ErrRecordNotFound)

		repository.
			EXPECT().
			Create(ctx, category).
			Return(errors.New("create category error"))

		_, err := service.Create(ctx, request)

		assert.Error(t, err, "create category error")
	})

	t.Run("creates successfully", func(t *testing.T) {
		request := category.CreateRequest{
			Name:          "Haberdashery",
			Income:        "true",
			Visible:       "true",
			Supercategory: "23",
		}

		category := &category.Category{
			Name:          "Haberdashery",
			Income:        true,
			Visible:       true,
			Supercategory: 23,
		}

		repository.
			EXPECT().
			FindByName(ctx, "Haberdashery").
			Return(nil, common.ErrRecordNotFound)

		repository.
			EXPECT().
			Create(ctx, category).
			Return(nil)

		_, err := service.Create(ctx, request)

		assert.NilError(t, err)
	})
}
