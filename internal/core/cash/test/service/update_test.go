package service_test

import (
	"errors"
	"testing"

	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/core/cash"
	"github.com/tksasha/balance/internal/core/cash/service"
	"github.com/tksasha/balance/internal/core/cash/test/mocks"
	"github.com/tksasha/balance/internal/core/common"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestUpdate(t *testing.T) { //nolint:funlen
	ctrl := gomock.NewController(t)

	cashRepository := mocks.NewMockRepository(ctrl)

	service := service.New(cashRepository)

	ctx := t.Context()

	t.Run("returns error when id is invalid", func(t *testing.T) {
		request := cash.UpdateRequest{
			ID: "abc",
		}

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "resource not found")
	})

	t.Run("returns error when cash was not found", func(t *testing.T) {
		request := cash.UpdateRequest{
			ID: "1530",
		}

		cashRepository.
			EXPECT().
			FindByID(ctx, 1530).
			Return(nil, common.ErrRecordNotFound)

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "resource not found")
	})

	t.Run("returns error when find failed", func(t *testing.T) {
		request := cash.UpdateRequest{
			ID: "1530",
		}

		cashRepository.
			EXPECT().
			FindByID(ctx, 1530).
			Return(nil, errors.New("failed to find cash by id"))

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "failed to find cash by id")
	})

	t.Run("returns error when formula is blank", func(t *testing.T) {
		request := cash.UpdateRequest{
			ID:      "1540",
			Formula: "",
			Name:    "Bonds",
		}

		cash := &cash.Cash{
			ID: 1540,
		}

		cashRepository.
			EXPECT().
			FindByID(ctx, 1540).
			Return(cash, nil)

		cashRepository.
			EXPECT().
			NameExists(ctx, "Bonds", 1540).
			Return(false, nil)

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "sum: is required")
	})

	t.Run("returns error when formula is invalid", func(t *testing.T) {
		request := cash.UpdateRequest{
			ID:      "1540",
			Formula: "abc",
			Name:    "Bonds",
		}

		cash := &cash.Cash{
			ID: 1540,
		}

		cashRepository.
			EXPECT().
			FindByID(ctx, 1540).
			Return(cash, nil)

		cashRepository.
			EXPECT().
			NameExists(ctx, "Bonds", 1540).
			Return(false, nil)

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "sum: is invalid")
	})

	t.Run("returns error when name is blank", func(t *testing.T) {
		request := cash.UpdateRequest{
			ID:      "1540",
			Formula: "2+3",
			Name:    "",
		}

		cash := &cash.Cash{
			ID: 1540,
		}

		cashRepository.
			EXPECT().
			FindByID(ctx, 1540).
			Return(cash, nil)

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "name: is required")
	})

	t.Run("returns error when name already exists", func(t *testing.T) {
		request := cash.UpdateRequest{
			ID:      "1540",
			Formula: "2+3",
			Name:    "Bonds",
		}

		cash := &cash.Cash{
			ID: 1540,
		}

		cashRepository.
			EXPECT().
			FindByID(ctx, 1540).
			Return(cash, nil)

		cashRepository.
			EXPECT().
			NameExists(ctx, "Bonds", 1540).
			Return(true, nil)

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "name: already exists")
	})

	t.Run("returns error when supercategory is invalid", func(t *testing.T) {
		request := cash.UpdateRequest{
			ID:            "1540",
			Formula:       "2+3",
			Name:          "Bonds",
			Supercategory: "abc",
		}

		cash := &cash.Cash{
			ID: 1540,
		}

		cashRepository.
			EXPECT().
			FindByID(ctx, 1540).
			Return(cash, nil)

		cashRepository.
			EXPECT().
			NameExists(ctx, "Bonds", 1540).
			Return(false, nil)

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "supercategory: is invalid")
	})

	t.Run("returns error when favorite is invalid", func(t *testing.T) {
		request := cash.UpdateRequest{
			ID:            "1540",
			Formula:       "2+3",
			Name:          "Bonds",
			Supercategory: "1548",
			Favorite:      "abc",
		}

		cash := &cash.Cash{
			ID: 1540,
		}

		cashRepository.
			EXPECT().
			FindByID(ctx, 1540).
			Return(cash, nil)

		cashRepository.
			EXPECT().
			NameExists(ctx, "Bonds", 1540).
			Return(false, nil)

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "favorite: is invalid")
	})

	t.Run("returns error when update failed", func(t *testing.T) {
		request := cash.UpdateRequest{
			ID:            "1540",
			Formula:       "2+3",
			Name:          "Bonds",
			Supercategory: "1548",
			Favorite:      "true",
		}

		cash := &cash.Cash{
			ID: 1540,
		}

		cashRepository.
			EXPECT().
			FindByID(ctx, 1540).
			Return(cash, nil)

		cashRepository.
			EXPECT().
			Update(ctx, cash).
			Return(errors.New("failed to update cash"))

		cashRepository.
			EXPECT().
			NameExists(ctx, "Bonds", 1540).
			Return(false, nil)

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "failed to update cash")
	})

	t.Run("returns updated cash when update succeeded", func(t *testing.T) {
		request := cash.UpdateRequest{
			ID:            "1540",
			Formula:       "2+3",
			Name:          "Bonds",
			Supercategory: "1548",
			Favorite:      "true",
		}

		cash := &cash.Cash{
			ID:       1540,
			Currency: currency.USD,
		}

		cashRepository.
			EXPECT().
			FindByID(ctx, 1540).
			Return(cash, nil)

		cashRepository.
			EXPECT().
			Update(ctx, cash).
			Return(nil)

		cashRepository.
			EXPECT().
			NameExists(ctx, "Bonds", 1540).
			Return(false, nil)

		res, err := service.Update(ctx, request)

		assert.NilError(t, err)

		assert.Equal(t, res.ID, 1540)
		assert.Equal(t, res.Currency, currency.USD)
		assert.Equal(t, res.Formula, "2+3")
		assert.Equal(t, res.Sum, 5.0)
		assert.Equal(t, res.Name, "Bonds")
		assert.Equal(t, res.Supercategory, 1548)
		assert.Equal(t, res.Favorite, true)
	})
}
