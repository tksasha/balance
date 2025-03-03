package service_test

import (
	"errors"
	"testing"

	"github.com/tksasha/balance/internal/app/balance/service"
	"github.com/tksasha/balance/internal/app/balance/test/mocks"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestBalance(t *testing.T) {
	controller := gomock.NewController(t)

	repository := mocks.NewMockRepository(controller)

	service := service.New(repository)

	ctx := t.Context()

	t.Run("returns error when calculate income failed", func(t *testing.T) {
		repository.EXPECT().Income(ctx).Return(0.0, errors.New("calculate income error"))

		_, _, err := service.Balance(ctx)

		assert.Error(t, err, "calculate income error")
	})

	t.Run("returns error when calculate expense failed", func(t *testing.T) {
		repository.EXPECT().Income(ctx).Return(99.99, nil)

		repository.EXPECT().Expense(ctx).Return(0.0, errors.New("calculate expense error"))

		_, _, err := service.Balance(ctx)

		assert.Error(t, err, "calculate expense error")
	})

	t.Run("returns error when calculate cashes failed", func(t *testing.T) {
		repository.EXPECT().Income(ctx).Return(99.99, nil)

		repository.EXPECT().Expense(ctx).Return(44.44, nil)

		repository.EXPECT().Cashes(ctx).Return(0.0, errors.New("calculate cashes error"))

		_, _, err := service.Balance(ctx)

		assert.Error(t, err, "calculate cashes error")
	})

	t.Run("returns balance when no errors", func(t *testing.T) {
		repository.EXPECT().Income(ctx).Return(99.99, nil)

		repository.EXPECT().Expense(ctx).Return(44.44, nil)

		repository.EXPECT().Cashes(ctx).Return(33.33, nil)

		residual, balance, err := service.Balance(ctx)

		assert.NilError(t, err)
		assert.Equal(t, residual, 55.55)
		assert.Equal(t, balance, -22.22)
	})
}
