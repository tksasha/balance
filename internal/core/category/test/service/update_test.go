package service_test

import (
	"errors"
	"testing"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/category/service"
	"github.com/tksasha/balance/internal/core/category/test/mocks"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestUpdate(t *testing.T) { //nolint:funlen
	ctrl := gomock.NewController(t)

	repository := mocks.NewMockRepository(ctrl)

	service := service.New(repository)

	ctx := t.Context()

	t.Run("returns error when id is invalid", func(t *testing.T) {
		request := category.UpdateRequest{
			ID: "abc",
		}

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "resource not found")
	})

	t.Run("returns error when category not found", func(t *testing.T) {
		request := category.UpdateRequest{
			ID: "1513",
		}

		repository.
			EXPECT().
			FindByID(ctx, 1513).
			Return(nil, apperrors.ErrRecordNotFound)

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "resource not found")
	})

	t.Run("returns error when find category failed", func(t *testing.T) {
		request := category.UpdateRequest{
			ID: "1515",
		}

		repository.
			EXPECT().
			FindByID(ctx, 1515).
			Return(nil, errors.New("find category error"))

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "find category error")
	})

	t.Run("returns error when name is blank", func(t *testing.T) {
		request := category.UpdateRequest{
			ID:   "1517",
			Name: "",
		}

		category := &category.Category{
			ID: 1517,
		}

		repository.
			EXPECT().
			FindByID(ctx, 1517).
			Return(category, nil)

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "name: is required")
	})

	t.Run("returns error when find by name fails", func(t *testing.T) {
		request := category.UpdateRequest{
			ID:   "1529",
			Name: "Entrepreneurship",
		}

		category := &category.Category{
			ID: 1529,
		}

		repository.
			EXPECT().
			FindByID(ctx, 1529).
			Return(category, nil)

		repository.
			EXPECT().
			FindByName(ctx, "Entrepreneurship").
			Return(nil, errors.New("find category by name error"))

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "find category by name error")
	})

	t.Run("returns error when name already exists", func(t *testing.T) {
		request := category.UpdateRequest{
			ID:   "1030",
			Name: "Beverages",
		}

		categoryToUpdate := &category.Category{
			ID:   1030,
			Name: "Food",
		}

		repository.
			EXPECT().
			FindByID(ctx, 1030).
			Return(categoryToUpdate, nil)

		categoryFound := &category.Category{
			ID:   1029,
			Name: "Beverages",
		}

		repository.
			EXPECT().
			FindByName(ctx, "Beverages").
			Return(categoryFound, nil)

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "name: already exists")
	})

	t.Run("returns error when update category failed", func(t *testing.T) {
		request := category.UpdateRequest{
			ID:   "1524",
			Name: "Miscellaneous",
		}

		categoryFound := &category.Category{
			ID: 1524,
		}

		repository.
			EXPECT().
			FindByID(ctx, 1524).
			Return(categoryFound, nil)

		repository.
			EXPECT().
			FindByName(ctx, "Miscellaneous").
			Return(nil, apperrors.ErrRecordNotFound)

		categoryToUpdate := &category.Category{
			ID:   1524,
			Name: "Miscellaneous",
		}

		repository.
			EXPECT().
			Update(ctx, categoryToUpdate).
			Return(errors.New("update category error"))

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "update category error")
	})

	t.Run("returns updated category when update succeeds", func(t *testing.T) {
		request := category.UpdateRequest{
			ID:            "1529",
			Name:          "Stationery",
			Income:        "true",
			Visible:       "true",
			Supercategory: "3",
		}

		categoryFound := &category.Category{
			ID:            1529,
			Name:          "Food",
			Income:        false,
			Visible:       false,
			Supercategory: 2,
		}

		repository.
			EXPECT().
			FindByID(ctx, 1529).
			Return(categoryFound, nil)

		repository.
			EXPECT().
			FindByName(ctx, "Stationery").
			Return(nil, apperrors.ErrRecordNotFound)

		categoryToUpdate := &category.Category{
			ID:            1529,
			Name:          "Stationery",
			Income:        true,
			Visible:       true,
			Supercategory: 3,
		}

		repository.
			EXPECT().
			Update(ctx, categoryToUpdate).
			Return(nil)

		category, err := service.Update(ctx, request)

		assert.NilError(t, err)
		assert.Equal(t, category.ID, 1529)
		assert.Equal(t, category.Name, "Stationery")
		assert.Equal(t, category.Income, true)
		assert.Equal(t, category.Visible, true)
		assert.Equal(t, category.Supercategory, 3)
	})
}
