package service_test

import (
	"errors"
	"testing"

	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/backoffice/cash/service"
	"github.com/tksasha/balance/internal/backoffice/cash/test/mocks"
	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/balance/internal/common/currency"
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

		cashRepository.EXPECT().FindByID(ctx, 1530).Return(nil, common.ErrRecordNotFound)

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "resource not found")
	})

	t.Run("returns error when find failed", func(t *testing.T) {
		request := cash.UpdateRequest{
			ID: "1530",
		}

		cashRepository.EXPECT().FindByID(ctx, 1530).Return(nil, errors.New("failed to find cash by id"))

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "failed to find cash by id")
	})

	t.Run("returns error when formula is blank", func(t *testing.T) {
		request := cash.UpdateRequest{
			ID:       "1540",
			Formula:  "",
			Name:     "Bonds",
			Currency: "eur",
		}

		cashToFind := &cash.Cash{
			ID: 1540,
		}

		cashRepository.EXPECT().FindByID(ctx, 1540).Return(cashToFind, nil)

		cashToCheck := &cash.Cash{
			ID:       1540,
			Name:     "Bonds",
			Currency: currency.EUR,
		}

		cashRepository.EXPECT().NameExists(ctx, cashToCheck).Return(false, nil)

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "sum: is required")
	})

	t.Run("returns error when formula is invalid", func(t *testing.T) {
		request := cash.UpdateRequest{
			ID:       "1540",
			Formula:  "abc",
			Name:     "Bonds",
			Currency: "eur",
		}

		cashToFind := &cash.Cash{
			ID: 1540,
		}

		cashRepository.EXPECT().FindByID(ctx, 1540).Return(cashToFind, nil)

		cashToCheck := &cash.Cash{
			ID:       1540,
			Name:     "Bonds",
			Currency: currency.EUR,
		}

		cashRepository.EXPECT().NameExists(ctx, cashToCheck).Return(false, nil)

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "sum: is invalid")
	})

	t.Run("returns error when name is blank", func(t *testing.T) {
		request := cash.UpdateRequest{
			ID:      "1540",
			Formula: "2+3",
			Name:    "",
		}

		cashToFind := &cash.Cash{
			ID: 1540,
		}

		cashRepository.EXPECT().FindByID(ctx, 1540).Return(cashToFind, nil)

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "name: is required")
	})

	t.Run("returns error when name already exists", func(t *testing.T) {
		request := cash.UpdateRequest{
			ID:       "1540",
			Formula:  "2+3",
			Name:     "Bonds",
			Currency: "eur",
		}

		cashToFind := &cash.Cash{
			ID: 1540,
		}

		cashRepository.EXPECT().FindByID(ctx, 1540).Return(cashToFind, nil)

		cashToCheck := &cash.Cash{
			ID:       1540,
			Name:     "Bonds",
			Currency: currency.EUR,
		}

		cashRepository.EXPECT().NameExists(ctx, cashToCheck).Return(true, nil)

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "name: already exists")
	})

	t.Run("returns error when supercategory is invalid", func(t *testing.T) {
		request := cash.UpdateRequest{
			ID:            "1540",
			Formula:       "2+3",
			Name:          "Bonds",
			Supercategory: "abc",
			Currency:      "eur",
		}

		cashToFind := &cash.Cash{
			ID: 1540,
		}

		cashRepository.EXPECT().FindByID(ctx, 1540).Return(cashToFind, nil)

		cashToCheck := &cash.Cash{
			ID:       1540,
			Name:     "Bonds",
			Currency: currency.EUR,
		}

		cashRepository.EXPECT().NameExists(ctx, cashToCheck).Return(false, nil)

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "supercategory: is invalid")
	})

	t.Run("returns error when update failed", func(t *testing.T) {
		request := cash.UpdateRequest{
			ID:            "1540",
			Formula:       "2+3",
			Name:          "Bonds",
			Supercategory: "1548",
			Currency:      "eur",
		}

		cashToFind := &cash.Cash{
			ID: 1540,
		}

		cashRepository.EXPECT().FindByID(ctx, 1540).Return(cashToFind, nil)

		cashToCheck := &cash.Cash{
			ID:       1540,
			Name:     "Bonds",
			Currency: currency.EUR,
		}

		cashRepository.EXPECT().NameExists(ctx, cashToCheck).Return(false, nil)

		cashRepository.EXPECT().Update(ctx, gomock.Any()).Return(errors.New("failed to update cash"))

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "failed to update cash")
	})

	t.Run("returns updated cash when update succeeded", func(t *testing.T) {
		ds := []struct {
			currencyCode string
			currency     currency.Currency
		}{
			{"", currency.UAH},
			{"abc", currency.UAH},
			{"uah", currency.UAH},
			{"usd", currency.USD},
			{"eur", currency.EUR},
		}

		for _, d := range ds {
			request := cash.UpdateRequest{
				ID:            "1540",
				Formula:       "2+3",
				Name:          "Bonds",
				Supercategory: "1548",
				Currency:      d.currencyCode,
			}

			cashToFind := &cash.Cash{
				ID:       1540,
				Currency: d.currency,
			}

			cashRepository.EXPECT().FindByID(ctx, 1540).Return(cashToFind, nil)

			cashToCheck := &cash.Cash{
				ID:       1540,
				Name:     "Bonds",
				Currency: d.currency,
			}

			cashRepository.EXPECT().NameExists(ctx, cashToCheck).Return(false, nil)

			cashRepository.EXPECT().Update(ctx, gomock.Any()).Return(nil)

			cashUpdated, err := service.Update(ctx, request)

			assert.NilError(t, err)

			assert.Equal(t, cashUpdated.ID, 1540)
			assert.Equal(t, cashUpdated.Currency, d.currency)
			assert.Equal(t, cashUpdated.Formula, "2+3")
			assert.Equal(t, cashUpdated.Sum, 5.0)
			assert.Equal(t, cashUpdated.Name, "Bonds")
			assert.Equal(t, cashUpdated.Supercategory, 1548)
		}
	})
}
