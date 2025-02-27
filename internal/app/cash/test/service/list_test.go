package service_test

import (
	"errors"
	"slices"
	"testing"

	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/balance/internal/app/cash/service"
	"github.com/tksasha/balance/internal/app/cash/test/mocks"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestList(t *testing.T) {
	controller := gomock.NewController(t)

	cashRepository := mocks.NewMockRepository(controller)

	service := service.New(cashRepository)

	ctx := t.Context()

	t.Run("when repository returns error", func(t *testing.T) {
		cashRepository.EXPECT().FindAll(ctx).Return(nil, errors.New("find all error"))

		_, err := service.List(ctx)

		assert.Error(t, err, "find all error")
	})

	t.Run("when repository returns cashes", func(t *testing.T) {
		expected := cash.Cashes{}

		cashRepository.EXPECT().FindAll(ctx).Return(expected, nil)

		actual, err := service.List(ctx)

		assert.NilError(t, err)
		assert.Assert(t, slices.Equal(actual, expected))
	})
}
