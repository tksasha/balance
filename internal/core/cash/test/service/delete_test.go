package service_test

import (
	"errors"
	"testing"

	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/balance/internal/core/cash/service"
	"github.com/tksasha/balance/internal/core/cash/test/mocks"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestCashService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)

	repository := mocks.NewMockRepository(ctrl)

	service := service.New(repository)

	ctx := t.Context()

	t.Run("returns error when id is invalid", func(t *testing.T) {
		err := service.Delete(ctx, "abc")

		assert.Error(t, err, "resource not found")
	})

	t.Run("returns error when delete cash failed", func(t *testing.T) {
		repository.
			EXPECT().
			Delete(ctx, 1059).
			Return(errors.New("failed to delete cash"))

		err := service.Delete(ctx, "1059")

		assert.Error(t, err, "failed to delete cash")
	})

	t.Run("returns error when cash not found", func(t *testing.T) {
		repository.
			EXPECT().
			Delete(ctx, 1103).
			Return(common.ErrRecordNotFound)

		err := service.Delete(ctx, "1103")

		assert.Error(t, err, "resource not found")
	})

	t.Run("returns nil when delete cash succeeded", func(t *testing.T) {
		repository.
			EXPECT().
			Delete(ctx, 1106).
			Return(nil)

		err := service.Delete(ctx, "1106")

		assert.NilError(t, err)
	})
}
