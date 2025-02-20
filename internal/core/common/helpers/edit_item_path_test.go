package helpers_test

import (
	"net/http"
	"testing"

	"github.com/tksasha/balance/internal/core/common/helpers"
	"gotest.tools/v3/assert"
)

func TestEditItemPath(t *testing.T) {
	testmap := map[*http.Request]string{
		nil: "/items/1418/edit",
	}

	for _, exp := range testmap {
		path := helpers.EditItemPath(1418)

		assert.Equal(t, path, exp)
	}
}
