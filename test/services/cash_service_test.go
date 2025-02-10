package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/requests"
	"github.com/tksasha/balance/internal/services"
	mocksforservices "github.com/tksasha/balance/mocks/services"
	"github.com/tksasha/balance/pkg/currencies"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestCashService_Create(t *testing.T) { //nolint:funlen
	controller := gomock.NewController(t)

	cashRepository := mocksforservices.NewMockCashRepository(controller)

	service := services.NewCashService(cashRepository)

	ctx := context.Background()

	t.Run("returns error when name is blank", func(t *testing.T) {
		request := requests.CashCreateRequest{
			Name:          "",
			Formula:       "2+3",
			Supercategory: "23",
			Favorite:      "false",
		}

		err := service.Create(ctx, request)

		assert.Error(t, err, "name: is required")
	})

	t.Run("returns error when check name existence failed", func(t *testing.T) {
		request := requests.CashCreateRequest{
			Name: "Bonds",
		}

		cashRepository.
			EXPECT().
			NameExists(ctx, "Bonds", 0).
			Return(false, errors.New("check name existence error"))

		err := service.Create(ctx, request)

		assert.Error(t, err, "check name existence error")
	})

	t.Run("returns error when name already exists", func(t *testing.T) {
		request := requests.CashCreateRequest{
			Name:    "Bonds",
			Formula: "2+3",
		}

		cashRepository.
			EXPECT().
			NameExists(ctx, "Bonds", 0).
			Return(true, nil)

		err := service.Create(ctx, request)

		assert.Error(t, err, "name: already exists")
	})

	t.Run("returns error when formula is blank", func(t *testing.T) {
		request := requests.CashCreateRequest{
			Name:    "Bonds",
			Formula: "",
		}

		cashRepository.
			EXPECT().
			NameExists(ctx, "Bonds", 0).
			Return(false, nil)

		err := service.Create(ctx, request)

		assert.Error(t, err, "formula: is required")
	})

	t.Run("returns error when formula is invalid", func(t *testing.T) {
		request := requests.CashCreateRequest{
			Name:    "Bonds",
			Formula: "abc",
		}

		cashRepository.
			EXPECT().
			NameExists(ctx, "Bonds", 0).
			Return(false, nil)

		err := service.Create(ctx, request)

		assert.Error(t, err, "formula: is invalid")
	})

	t.Run("returns error when supercategory is invalid", func(t *testing.T) {
		request := requests.CashCreateRequest{
			Name:          "Bonds",
			Formula:       "2+3",
			Supercategory: "abc",
		}

		cashRepository.
			EXPECT().
			NameExists(ctx, "Bonds", 0).
			Return(false, nil)

		err := service.Create(ctx, request)

		assert.Error(t, err, "supercategory: is invalid")
	})

	t.Run("returns error when create failed", func(t *testing.T) {
		request := requests.CashCreateRequest{
			Name:          "Bonds",
			Formula:       "2+3",
			Supercategory: "1242",
			Favorite:      "true",
		}

		cashRepository.
			EXPECT().
			NameExists(ctx, "Bonds", 0).
			Return(false, nil)

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

	t.Run("returns nil when create succeeded", func(t *testing.T) {
		request := requests.CashCreateRequest{
			Name:          "Bonds",
			Formula:       "2+3",
			Supercategory: "1242",
			Favorite:      "true",
		}

		cashRepository.
			EXPECT().
			NameExists(ctx, "Bonds", 0).
			Return(false, nil)

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

func TestCashService_FindByID(t *testing.T) {
	controller := gomock.NewController(t)

	cashRepository := mocksforservices.NewMockCashRepository(controller)

	service := services.NewCashService(cashRepository)

	ctx := context.Background()

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

func TestCashService_Update(t *testing.T) { //nolint:funlen
	controller := gomock.NewController(t)

	cashRepository := mocksforservices.NewMockCashRepository(controller)

	service := services.NewCashService(cashRepository)

	ctx := context.Background()

	t.Run("returns error when id is invalid", func(t *testing.T) {
		request := requests.CashUpdateRequest{
			ID: "abc",
		}

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "resource not found")
	})

	t.Run("returns error when cash was not found", func(t *testing.T) {
		request := requests.CashUpdateRequest{
			ID: "1530",
		}

		cashRepository.
			EXPECT().
			FindByID(ctx, 1530).
			Return(nil, apperrors.ErrRecordNotFound)

		_, err := service.Update(ctx, request)

		assert.Error(t, err, "resource not found")
	})

	t.Run("returns error when find failed", func(t *testing.T) {
		request := requests.CashUpdateRequest{
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
		request := requests.CashUpdateRequest{
			ID:      "1540",
			Formula: "",
			Name:    "Bonds",
		}

		cash := &models.Cash{
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
		request := requests.CashUpdateRequest{
			ID:      "1540",
			Formula: "abc",
			Name:    "Bonds",
		}

		cash := &models.Cash{
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
		request := requests.CashUpdateRequest{
			ID:      "1540",
			Formula: "2+3",
			Name:    "",
		}

		cash := &models.Cash{
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
		request := requests.CashUpdateRequest{
			ID:      "1540",
			Formula: "2+3",
			Name:    "Bonds",
		}

		cash := &models.Cash{
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
		request := requests.CashUpdateRequest{
			ID:            "1540",
			Formula:       "2+3",
			Name:          "Bonds",
			Supercategory: "abc",
		}

		cash := &models.Cash{
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
		request := requests.CashUpdateRequest{
			ID:            "1540",
			Formula:       "2+3",
			Name:          "Bonds",
			Supercategory: "1548",
			Favorite:      "abc",
		}

		cash := &models.Cash{
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
		request := requests.CashUpdateRequest{
			ID:            "1540",
			Formula:       "2+3",
			Name:          "Bonds",
			Supercategory: "1548",
			Favorite:      "true",
		}

		cash := &models.Cash{
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
		request := requests.CashUpdateRequest{
			ID:            "1540",
			Formula:       "2+3",
			Name:          "Bonds",
			Supercategory: "1548",
			Favorite:      "true",
		}

		cash := &models.Cash{
			ID:       1540,
			Currency: currencies.USD,
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
		assert.Equal(t, res.Currency, currencies.USD)
		assert.Equal(t, res.Formula, "2+3")
		assert.Equal(t, res.Sum, 5.0)
		assert.Equal(t, res.Name, "Bonds")
		assert.Equal(t, res.Supercategory, 1548)
		assert.Equal(t, res.Favorite, true)
	})
}

func TestCashService_Delete(t *testing.T) {
	controller := gomock.NewController(t)

	cashRepository := mocksforservices.NewMockCashRepository(controller)

	service := services.NewCashService(cashRepository)

	ctx := context.Background()

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
