package services_test

import (
	"errors"
	"testing"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/services"
	mocksforservices "github.com/tksasha/balance/mocks/services"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestCashService_FindByID(t *testing.T) {
	controller := gomock.NewController(t)

	cashRepository := mocksforservices.NewMockCashRepository(controller)

	service := services.NewCashService(cashRepository)

	ctx := t.Context()

	t.Run("returns error when id is invalid", func(t *testing.T) {
		_, err := service.FindByID(ctx, "abc")

		assert.Error(t, err, "resource not found")
	})

	t.Run("returns error when cash was not found", func(t *testing.T) {
		cashRepository.
			EXPECT().
			FindByID(ctx, 1230).
			Return(nil, apperrors.ErrRecordNotFound)

		_, err := service.FindByID(ctx, "1230")

		assert.Error(t, err, "resource not found")
	})

	t.Run("returns cash when it was found", func(t *testing.T) {
		cash := &models.Cash{}

		cashRepository.
			EXPECT().
			FindByID(ctx, 1230).
			Return(cash, nil)

		res, err := service.FindByID(ctx, "1230")

		assert.Equal(t, res, cash)
		assert.NilError(t, err)
	})
}

func TestCashService_Delete(t *testing.T) {
	controller := gomock.NewController(t)

	cashRepository := mocksforservices.NewMockCashRepository(controller)

	service := services.NewCashService(cashRepository)

	ctx := t.Context()

	t.Run("returns error when id is invalid", func(t *testing.T) {
		err := service.Delete(ctx, "abc")

		assert.Error(t, err, "resource not found")
	})

	t.Run("returns error when delete cash failed", func(t *testing.T) {
		cashRepository.
			EXPECT().
			Delete(ctx, 1059).
			Return(errors.New("failed to delete cash"))

		err := service.Delete(ctx, "1059")

		assert.Error(t, err, "failed to delete cash")
	})

	t.Run("returns error when cash not found", func(t *testing.T) {
		cashRepository.
			EXPECT().
			Delete(ctx, 1103).
			Return(apperrors.ErrRecordNotFound)

		err := service.Delete(ctx, "1103")

		assert.Error(t, err, "resource not found")
	})

	t.Run("returns nil when delete cash succeeded", func(t *testing.T) {
		cashRepository.
			EXPECT().
			Delete(ctx, 1106).
			Return(nil)

		err := service.Delete(ctx, "1106")

		assert.NilError(t, err)
	})
}
