package service_test

import (
	"errors"
	"testing"

	"github.com/tksasha/balance/internal/core/category/service"
	"github.com/tksasha/balance/internal/core/category/test/mocks"
	"github.com/tksasha/balance/internal/core/common"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)

	repository := mocks.NewMockRepository(ctrl)

	service := service.New(repository)

	ctx := t.Context()

	t.Run("returns error when id is invalid", func(t *testing.T) {
		err := service.Delete(ctx, "abc")

		assert.Error(t, err, "resource not found")
	})

	t.Run("returns error when category not found", func(t *testing.T) {
		repository.
			EXPECT().
			Delete(ctx, 1230).
			Return(common.ErrRecordNotFound)

		err := service.Delete(ctx, "1230")

		assert.Error(t, err, "resource not found")
	})

	t.Run("returns error when delete category failed", func(t *testing.T) {
		repository.
			EXPECT().
			Delete(ctx, 1230).
			Return(errors.New("delete category error"))

		err := service.Delete(ctx, "1230")

		assert.Error(t, err, "delete category error")
	})

	t.Run("returns nil when category deleted", func(t *testing.T) {
		repository.
			EXPECT().
			Delete(ctx, 1230).
			Return(nil)

		err := service.Delete(ctx, "1230")

		assert.NilError(t, err)
	})
}
