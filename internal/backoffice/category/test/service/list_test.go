package service_test

import (
	"errors"
	"slices"
	"testing"

	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/balance/internal/backoffice/category/service"
	"github.com/tksasha/balance/internal/backoffice/category/test/mocks"
	"github.com/tksasha/balance/internal/common"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestList(t *testing.T) {
	controller := gomock.NewController(t)

	categoryRepository := mocks.NewMockRepository(controller)

	service := service.New(common.NewBaseService(), categoryRepository)

	ctx := t.Context()

	t.Run("returns error when find all categories fails", func(t *testing.T) {
		categoryRepository.EXPECT().FindAll(ctx).Return(nil, errors.New("find all categories error"))

		_, err := service.List(ctx)

		assert.Error(t, err, "find all categories error")
	})

	t.Run("returns categories when find all categories succeeds", func(t *testing.T) {
		expected := category.Categories{}

		categoryRepository.EXPECT().FindAll(ctx).Return(expected, nil)

		actual, err := service.List(ctx)

		assert.NilError(t, err)
		assert.Assert(t, slices.Equal(actual, expected))
	})
}
