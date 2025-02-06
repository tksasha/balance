package cash_test

import (
	"context"
	"errors"
	"testing"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/requests"
	"github.com/tksasha/balance/internal/services"
	mocksforservices "github.com/tksasha/balance/mocks/services"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestCreate(t *testing.T) { //nolint:funlen
	controller := gomock.NewController(t)

	cashRepository := mocksforservices.NewMockCashRepository(controller)

	service := services.NewCashService(cashRepository)

	ctx := context.Background()

	t.Run("returns error when name is blank", func(t *testing.T) {
		request := requests.CreateCashRequest{
			Name:          "",
			Formula:       "2+3",
			Supercategory: "23",
			Favorite:      "false",
		}

		err := service.Create(ctx, request)

		assert.Error(t, err, "name: is required")
	})

	t.Run("returns error when find cash by name fails", func(t *testing.T) {
		request := requests.CreateCashRequest{
			Name: "Bonds",
		}

		cashRepository.
			EXPECT().
			FindByName(ctx, "Bonds").
			Return(nil, errors.New("find cash by name error"))

		err := service.Create(ctx, request)

		assert.Error(t, err, "find cash by name error")
	})

	t.Run("returns error when name already exists", func(t *testing.T) {
		request := requests.CreateCashRequest{
			Name:          "Bonds",
			Formula:       "2+3",
			Supercategory: "25",
			Favorite:      "true",
		}

		cash := &models.Cash{
			ID:   1151,
			Name: "Bonds",
		}

		cashRepository.
			EXPECT().
			FindByName(ctx, "Bonds").
			Return(cash, nil)

		err := service.Create(ctx, request)

		assert.Error(t, err, "name: already exists")
	})

	t.Run("returns error when formula is blank", func(t *testing.T) {
		request := requests.CreateCashRequest{
			Name:          "Bonds",
			Formula:       "",
			Supercategory: "25",
			Favorite:      "false",
		}

		cashRepository.
			EXPECT().
			FindByName(ctx, "Bonds").
			Return(nil, apperrors.ErrRecordNotFound)

		err := service.Create(ctx, request)

		assert.Error(t, err, "formula: is required")
	})

	t.Run("returns error when formula is invalid", func(t *testing.T) {
		request := requests.CreateCashRequest{
			Name:          "Bonds",
			Formula:       "abc",
			Supercategory: "26",
			Favorite:      "false",
		}

		cashRepository.
			EXPECT().
			FindByName(ctx, "Bonds").
			Return(nil, apperrors.ErrRecordNotFound)

		err := service.Create(ctx, request)

		assert.Error(t, err, "formula: is invalid")
	})

	t.Run("returns error when supercategory is invalid", func(t *testing.T) {
		request := requests.CreateCashRequest{
			Name:          "Bonds",
			Formula:       "2+3",
			Supercategory: "abc",
			Favorite:      "false",
		}

		cashRepository.
			EXPECT().
			FindByName(ctx, "Bonds").
			Return(nil, apperrors.ErrRecordNotFound)

		err := service.Create(ctx, request)

		assert.Error(t, err, "supercategory: is invalid")
	})

	t.Run("returns error when create cash fails", func(t *testing.T) {
		request := requests.CreateCashRequest{
			Name:          "Bonds",
			Formula:       "2+3",
			Supercategory: "1242",
			Favorite:      "true",
		}

		cashRepository.
			EXPECT().
			FindByName(ctx, "Bonds").
			Return(nil, apperrors.ErrRecordNotFound)

		cash := &models.Cash{
			Name:          "Bonds",
			Formula:       "2+3",
			Sum:           5.0,
			Supercategory: 1242,
			Favorite:      true,
		}

		cashRepository.
			EXPECT().
			Create(ctx, cash).
			Return(errors.New("create cash error"))

		err := service.Create(ctx, request)

		assert.Error(t, err, "create cash error")
	})

	t.Run("creates cash successfully", func(t *testing.T) {
		request := requests.CreateCashRequest{
			Name:          "Bonds",
			Formula:       "2+3",
			Supercategory: "1242",
			Favorite:      "true",
		}

		cashRepository.
			EXPECT().
			FindByName(ctx, "Bonds").
			Return(nil, apperrors.ErrRecordNotFound)

		cash := &models.Cash{
			Name:          "Bonds",
			Formula:       "2+3",
			Sum:           5.0,
			Supercategory: 1242,
			Favorite:      true,
		}

		cashRepository.
			EXPECT().
			Create(ctx, cash).
			Return(nil)

		err := service.Create(ctx, request)

		assert.NilError(t, err)
	})
}
