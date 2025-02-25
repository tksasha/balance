package service_test

import (
	"errors"
	"slices"
	"testing"

	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/app/category/service"
	"github.com/tksasha/balance/internal/app/category/test/mocks"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestList(t *testing.T) {
	ctrl := gomock.NewController(t)

	repository := mocks.NewMockRepository(ctrl)

	service := service.New(repository)

	ctx := t.Context()

	t.Run("returns error when find all categories failed", func(t *testing.T) {
		repository.
			EXPECT().
			FindAll(ctx).
			Return(nil, errors.New("find all categories error"))

		_, err := service.List(ctx)

		assert.Error(t, err, "find all categories error")
	})

	t.Run("returns categories when list succeeded", func(t *testing.T) {
		expected := category.Categories{}

		repository.
			EXPECT().
			FindAll(ctx).
			Return(expected, nil)

		result, err := service.List(ctx)

		assert.Assert(t, slices.Equal(result, expected))
		assert.NilError(t, err)
	})
}
