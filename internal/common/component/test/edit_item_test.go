package component_test

import (
	"testing"

	"github.com/tksasha/balance/internal/core/common/component"
	"gotest.tools/v3/assert"
)

func TestEditItem(t *testing.T) {
	component := component.New()

	actual := component.EditItem(1730)

	expected := "/items/1730/edit"

	assert.Equal(t, actual, expected)
}
