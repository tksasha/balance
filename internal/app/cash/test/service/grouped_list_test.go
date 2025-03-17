package service_test

import (
	"errors"
	"slices"
	"testing"

	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/balance/internal/app/cash/service"
	"github.com/tksasha/balance/internal/app/cash/test/mocks"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestList(t *testing.T) {
	controller := gomock.NewController(t)

	cashRepository := mocks.NewMockRepository(controller)

	service := service.New(cashRepository)

	ctx := t.Context()

	t.Run("when repository returns error", func(t *testing.T) {
		cashRepository.EXPECT().FindAll(ctx).Return(nil, errors.New("find all error"))

		_, err := service.GroupedList(ctx)

		assert.Error(t, err, "find all error")
	})

	t.Run("when repository returns cashes", func(t *testing.T) {
		cashes := cash.Cashes{
			{ID: 1, Name: "First", Supercategory: 3},
			{ID: 2, Name: "Second", Supercategory: 2},
			{ID: 3, Name: "Third", Supercategory: 0},
			{ID: 4, Name: "Fourth", Supercategory: 2},
			{ID: 5, Name: "Fifth", Supercategory: 3},
		}

		expected := cash.GroupedCashes{
			0: {cashes[2]},
			2: {cashes[1], cashes[3]},
			3: {cashes[0], cashes[4]},
		}

		cashRepository.EXPECT().FindAll(ctx).Return(cashes, nil)

		actual, err := service.GroupedList(ctx)

		assert.NilError(t, err)
		assert.Assert(t, slices.Equal(actual[0], expected[0]))
		assert.Assert(t, slices.Equal(actual[2], expected[2]))
		assert.Assert(t, slices.Equal(actual[3], expected[3]))
	})
}
