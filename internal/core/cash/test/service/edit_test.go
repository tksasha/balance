package service_test

import (
	"testing"

	"github.com/tksasha/balance/internal/core/cash"
	"github.com/tksasha/balance/internal/core/cash/service"
	"github.com/tksasha/balance/internal/core/cash/test/mocks"
	"github.com/tksasha/balance/internal/core/common"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestEdit(t *testing.T) {
	ctrl := gomock.NewController(t)

	repository := mocks.NewMockRepository(ctrl)

	service := service.New(repository)

	ctx := t.Context()

	t.Run("returns error when id is invalid", func(t *testing.T) {
		_, err := service.Edit(ctx, "abc")

		assert.Error(t, err, "resource not found")
	})

	t.Run("returns error when cash was not found", func(t *testing.T) {
		repository.
			EXPECT().
			FindByID(ctx, 1230).
			Return(nil, common.ErrRecordNotFound)

		_, err := service.Edit(ctx, "1230")

		assert.Error(t, err, "resource not found")
	})

	t.Run("returns cash when it was found", func(t *testing.T) {
		cash := &cash.Cash{}

		repository.
			EXPECT().
			FindByID(ctx, 1230).
			Return(cash, nil)

		res, err := service.Edit(ctx, "1230")

		assert.Equal(t, res, cash)
		assert.NilError(t, err)
	})
}
