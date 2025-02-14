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

func TestFindByID(t *testing.T) {
	ctrl := gomock.NewController(t)

	repository := mocks.NewMockRepository(ctrl)

	service := service.New(repository)

	ctx := t.Context()

	t.Run("returns error when id is invalid", func(t *testing.T) {
		_, err := service.FindByID(ctx, "abc")

		assert.Error(t, err, "resource not found")
	})

	t.Run("returns error when category not found", func(t *testing.T) {
		repository.
			EXPECT().
			FindByID(ctx, 1452).
			Return(nil, apperrors.ErrRecordNotFound)

		_, err := service.FindByID(ctx, "1452")

		assert.Error(t, err, "resource not found")
	})

	t.Run("returns error when find category failed", func(t *testing.T) {
		repository.
			EXPECT().
			FindByID(ctx, 1531).
			Return(nil, errors.New("find category by id error"))

		_, err := service.FindByID(ctx, "1531")

		assert.Error(t, err, "find category by id error")
	})

	t.Run("returns category when it is found", func(t *testing.T) {
		expected := &category.Category{}

		repository.
			EXPECT().
			FindByID(ctx, 1536).
			Return(expected, nil)

		result, err := service.FindByID(ctx, "1536")

		assert.NilError(t, err)
		assert.Equal(t, result, expected)
	})
}
