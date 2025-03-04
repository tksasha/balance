package service_test

import (
	"errors"
	"slices"
	"testing"

	"github.com/tksasha/balance/internal/app/categoryreport"
	"github.com/tksasha/balance/internal/app/categoryreport/service"
	"github.com/tksasha/balance/internal/app/categoryreport/test/mocks"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestReport(t *testing.T) {
	controller := gomock.NewController(t)

	repository := mocks.NewMockRepository(controller)

	service := service.New(repository)

	ctx := t.Context()

	request := categoryreport.Request{
		Year:  "2024",
		Month: "2",
	}

	filters := categoryreport.Filters{
		From: "2024-02-01",
		To:   "2024-02-29",
	}

	t.Run("when repository returns error", func(t *testing.T) {
		repository.EXPECT().Group(ctx, filters).Return(nil, errors.New("grouping error"))

		_, err := service.Report(ctx, request)

		assert.Error(t, err, "grouping error")
	})

	t.Run("when repository returns entities", func(t *testing.T) {
		expected := categoryreport.Entities{}

		repository.EXPECT().Group(ctx, filters).Return(expected, nil)

		actual, err := service.Report(ctx, request)

		assert.NilError(t, err)
		assert.Assert(t, slices.Equal(actual, expected))
	})
}
