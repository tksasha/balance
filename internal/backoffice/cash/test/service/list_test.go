package service_test

import (
	"errors"
	"slices"
	"testing"

	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/backoffice/cash/service"
	"github.com/tksasha/balance/internal/backoffice/cash/test/mocks"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestList(t *testing.T) {
	controller := gomock.NewController(t)

	cashRepository := mocks.NewMockRepository(controller)

	service := service.New(cashRepository)

	ctx := t.Context()

	t.Run("returns error when list cashes returns error", func(t *testing.T) {
		cashRepository.EXPECT().List(ctx).Return(nil, errors.New("list cashes error"))

		_, err := service.List(ctx)

		assert.Error(t, err, "list cashes error")
	})

	t.Run("returns cashes when list cashes doesn't return error", func(t *testing.T) {
		expected := cash.Cashes{}

		cashRepository.EXPECT().List(ctx).Return(expected, nil)

		actual, err := service.List(ctx)

		assert.NilError(t, err)
		assert.Assert(t, slices.Equal(actual, expected))
	})
}
