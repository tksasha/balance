package service_test

import (
	"errors"
	"testing"

	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/balance/internal/app/cash/service"
	"github.com/tksasha/balance/internal/app/cash/test/mocks"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestCreate(t *testing.T) { //nolint:funlen
	ctrl := gomock.NewController(t)

	cashRepository := mocks.NewMockRepository(ctrl)

	service := service.New(cashRepository)

	ctx := t.Context()

	t.Run("returns error when name is blank", func(t *testing.T) {
		request := cash.CreateRequest{
			Name:          "",
			Formula:       "2+3",
			Supercategory: "23",
			Favorite:      "false",
		}

		_, err := service.Create(ctx, request)

		assert.Error(t, err, "name: is required")
	})

	t.Run("returns error when check name existence failed", func(t *testing.T) {
		request := cash.CreateRequest{
			Name: "Bonds",
		}

		cashRepository.
			EXPECT().
			NameExists(ctx, "Bonds", 0).
			Return(false, errors.New("check name existence error"))

		_, err := service.Create(ctx, request)

		assert.Error(t, err, "check name existence error")
	})

	t.Run("returns error when name already exists", func(t *testing.T) {
		request := cash.CreateRequest{
			Name:    "Bonds",
			Formula: "2+3",
		}

		cashRepository.
			EXPECT().
			NameExists(ctx, "Bonds", 0).
			Return(true, nil)

		_, err := service.Create(ctx, request)

		assert.Error(t, err, "name: already exists")
	})

	t.Run("returns error when formula is blank", func(t *testing.T) {
		request := cash.CreateRequest{
			Name:    "Bonds",
			Formula: "",
		}

		cashRepository.
			EXPECT().
			NameExists(ctx, "Bonds", 0).
			Return(false, nil)

		_, err := service.Create(ctx, request)

		assert.Error(t, err, "formula: is required")
	})

	t.Run("returns error when formula is invalid", func(t *testing.T) {
		request := cash.CreateRequest{
			Name:    "Bonds",
			Formula: "abc",
		}

		cashRepository.
			EXPECT().
			NameExists(ctx, "Bonds", 0).
			Return(false, nil)

		_, err := service.Create(ctx, request)

		assert.Error(t, err, "formula: is invalid")
	})

	t.Run("returns error when supercategory is invalid", func(t *testing.T) {
		request := cash.CreateRequest{
			Name:          "Bonds",
			Formula:       "2+3",
			Supercategory: "abc",
		}

		cashRepository.
			EXPECT().
			NameExists(ctx, "Bonds", 0).
			Return(false, nil)

		_, err := service.Create(ctx, request)

		assert.Error(t, err, "supercategory: is invalid")
	})

	t.Run("returns error when create failed", func(t *testing.T) {
		request := cash.CreateRequest{
			Name:          "Bonds",
			Formula:       "2+3",
			Supercategory: "1242",
			Favorite:      "true",
		}

		cashRepository.
			EXPECT().
			NameExists(ctx, "Bonds", 0).
			Return(false, nil)

		cash := &cash.Cash{
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

		_, err := service.Create(ctx, request)

		assert.Error(t, err, "create cash error")
	})

	t.Run("returns nil when create succeeded", func(t *testing.T) {
		request := cash.CreateRequest{
			Name:          "Bonds",
			Formula:       "2+3",
			Supercategory: "1242",
			Favorite:      "true",
		}

		cashRepository.
			EXPECT().
			NameExists(ctx, "Bonds", 0).
			Return(false, nil)

		cash := &cash.Cash{
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

		_, err := service.Create(ctx, request)

		assert.NilError(t, err)
	})
}
