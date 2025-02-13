package service_test

import (
	"errors"
	"slices"
	"testing"

	"github.com/tksasha/balance/internal/category"
	"github.com/tksasha/balance/internal/category/service"
	"github.com/tksasha/balance/internal/category/test/mocks"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestList(t *testing.T) {
	ctrl := gomock.NewController(t)

	repository := mocks.NewMockRepository(ctrl)

	service := service.New(repository)

	ctx := t.Context()

	t.Run("returns error when list categories failed", func(t *testing.T) {
		repository.
			EXPECT().
			List(ctx).
			Return(nil, errors.New("list categories error"))

		_, err := service.List(ctx)

		assert.Error(t, err, "list categories error")
	})

	t.Run("returns categories when list succeeded", func(t *testing.T) {
		expected := category.Categories{}

		repository.
			EXPECT().
			List(ctx).
			Return(expected, nil)

		result, err := service.List(ctx)

		assert.Assert(t, slices.Equal(result, expected))
		assert.NilError(t, err)
	})
}
