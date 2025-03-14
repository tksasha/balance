package service_test

import (
	"errors"
	"testing"

	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/backoffice/cash/service"
	"github.com/tksasha/balance/internal/backoffice/cash/test/mocks"
	"github.com/tksasha/balance/internal/common/currency"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestCreate(t *testing.T) { //nolint:funlen
	controller := gomock.NewController(t)

	cashRepository := mocks.NewMockRepository(controller)

	service := service.New(cashRepository)

	ctx := t.Context()

	t.Run("returns error when name is blank", func(t *testing.T) {
		request := cash.CreateRequest{
			Name:          "",
			Formula:       "2+3",
			Supercategory: "23",
		}

		_, err := service.Create(ctx, request)

		assert.Error(t, err, "name: is required")
	})

	t.Run("when check name exists for currency returns error", func(t *testing.T) {
		request := cash.CreateRequest{
			Name:     "Bonds",
			Currency: "eur",
		}

		cashToCheck := &cash.Cash{
			ID:       0,
			Name:     "Bonds",
			Currency: currency.EUR,
		}

		cashRepository.EXPECT().NameExists(ctx, cashToCheck).Return(false, errors.New("check name existence error"))

		_, err := service.Create(ctx, request)

		assert.Error(t, err, "check name existence error")
	})

	t.Run("when name already exists for this currency", func(t *testing.T) {
		request := cash.CreateRequest{
			Name:     "Bonds",
			Formula:  "2+3",
			Currency: "eur",
		}

		cashToCheck := &cash.Cash{
			ID:       0,
			Name:     "Bonds",
			Currency: currency.EUR,
		}

		cashRepository.EXPECT().NameExists(ctx, cashToCheck).Return(true, nil)

		_, err := service.Create(ctx, request)

		assert.Error(t, err, "name: already exists")
	})

	t.Run("returns error when formula is blank", func(t *testing.T) {
		request := cash.CreateRequest{
			Name:     "Bonds",
			Formula:  "",
			Currency: "eur",
		}

		cashToCheck := &cash.Cash{
			ID:       0,
			Name:     "Bonds",
			Currency: currency.EUR,
		}

		cashRepository.EXPECT().NameExists(ctx, cashToCheck).Return(false, nil)

		_, err := service.Create(ctx, request)

		assert.Error(t, err, "sum: is required")
	})

	t.Run("returns error when formula is invalid", func(t *testing.T) {
		request := cash.CreateRequest{
			Name:     "Bonds",
			Formula:  "abc",
			Currency: "eur",
		}

		cashToCheck := &cash.Cash{
			ID:       0,
			Name:     "Bonds",
			Currency: currency.EUR,
		}

		cashRepository.EXPECT().NameExists(ctx, cashToCheck).Return(false, nil)

		_, err := service.Create(ctx, request)

		assert.Error(t, err, "sum: is invalid")
	})

	t.Run("returns error when supercategory is invalid", func(t *testing.T) {
		request := cash.CreateRequest{
			Name:          "Bonds",
			Formula:       "2+3",
			Supercategory: "abc",
			Currency:      "eur",
		}

		cashToCheck := &cash.Cash{
			ID:       0,
			Name:     "Bonds",
			Currency: currency.EUR,
		}

		cashRepository.EXPECT().NameExists(ctx, cashToCheck).Return(false, nil)

		_, err := service.Create(ctx, request)

		assert.Error(t, err, "supercategory: is invalid")
	})

	t.Run("when create returns error", func(t *testing.T) {
		request := cash.CreateRequest{
			Name:          "Bonds",
			Formula:       "2+3",
			Supercategory: "1242",
			Currency:      "eur",
		}

		cashToCheck := &cash.Cash{
			ID:       0,
			Name:     "Bonds",
			Currency: currency.EUR,
		}

		cashRepository.EXPECT().NameExists(ctx, cashToCheck).Return(false, nil)

		cashToCreate := &cash.Cash{
			Name:          "Bonds",
			Formula:       "2+3",
			Sum:           5,
			Supercategory: 1242,
			Currency:      currency.EUR,
		}

		cashRepository.EXPECT().Create(ctx, cashToCreate).Return(errors.New("create cash error"))

		_, err := service.Create(ctx, request)

		assert.Error(t, err, "create cash error")
	})

	t.Run("when create cash", func(t *testing.T) {
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
			request := cash.CreateRequest{
				Name:          "Bonds",
				Formula:       "2+3",
				Supercategory: "1242",
				Currency:      d.currencyCode,
			}

			cashToCheck := &cash.Cash{
				ID:       0,
				Name:     "Bonds",
				Currency: d.currency,
			}

			cashRepository.EXPECT().NameExists(ctx, cashToCheck).Return(false, nil)

			cashToCreate := &cash.Cash{
				Name:          "Bonds",
				Formula:       "2+3",
				Sum:           5,
				Supercategory: 1242,
				Currency:      d.currency,
			}

			cashRepository.EXPECT().Create(ctx, cashToCreate).Return(nil)

			_, err := service.Create(ctx, request)

			assert.NilError(t, err)
		}
	})
}
