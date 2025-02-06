package cash_test

import (
	"context"
	"testing"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/services"
	mocksforservices "github.com/tksasha/balance/mocks/services"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestFindByID(t *testing.T) {
	controller := gomock.NewController(t)

	cashRepository := mocksforservices.NewMockCashRepository(controller)

	service := services.NewCashService(cashRepository)

	ctx := context.Background()

	t.Run("returns error on invalid id", func(t *testing.T) {
		_, err := service.FindByID(ctx, "abc")

		assert.Error(t, err, "resource not found")
	})

	t.Run("returns error on not found", func(t *testing.T) {
		cashRepository.
			EXPECT().
			FindByID(ctx, 1230).
			Return(nil, apperrors.ErrRecordNotFound)

		_, err := service.FindByID(ctx, "1230")

		assert.Error(t, err, "resource not found")
	})

	t.Run("returns cash on success", func(t *testing.T) {
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
