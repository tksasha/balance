package service_test

import (
	"errors"
	"testing"

	"github.com/tksasha/balance/internal/core/indexpage/service"
	"github.com/tksasha/balance/internal/core/indexpage/test/mocks"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestResidual(t *testing.T) {
	controller := gomock.NewController(t)

	repository := mocks.NewMockRepository(controller)

	service := service.New(repository)

	ctx := t.Context()

	t.Run("returns error when income returns error", func(t *testing.T) {
		repository.
			EXPECT().
			Income(ctx).
			Return(0.0, errors.New("calculate income error"))

		_, err := service.Residual(ctx)

		assert.Error(t, err, "calculate income error")
	})

	t.Run("returns error when expense returns error", func(t *testing.T) {
		repository.
			EXPECT().
			Income(ctx).
			Return(99.99, nil)

		repository.
			EXPECT().
			Expense(ctx).
			Return(0.0, errors.New("calculate expense error"))

		_, err := service.Residual(ctx)

		assert.Error(t, err, "calculate expense error")
	})

	t.Run("returns calculated balance when no errors", func(t *testing.T) {
		repository.
			EXPECT().
			Income(ctx).
			Return(99.99, nil)

		repository.
			EXPECT().
			Expense(ctx).
			Return(44.44, nil)

		balance, err := service.Residual(ctx)

		assert.NilError(t, err)
		assert.Equal(t, balance, 55.55)
	})
}
