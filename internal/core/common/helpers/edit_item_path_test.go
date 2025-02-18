package helpers_test

import (
	"net/http"
	"testing"

	"github.com/tksasha/balance/internal/core/common/helpers"
	"github.com/tksasha/balance/internal/core/common/valueobjects/mocks"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestEditItemPath(t *testing.T) {
	controller := gomock.NewController(t)

	currentDateProvider := mocks.NewMockCurrentDateProvider(controller)

	helpers := helpers.New(currentDateProvider)

	testmap := map[*http.Request]string{
		nil: "/items/1418/edit",
	}

	for _, exp := range testmap {
		path := helpers.EditItemPath(1418)

		assert.Equal(t, path, exp)
	}
}
