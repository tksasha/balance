package component_test

import (
	"testing"

	"github.com/tksasha/balance/internal/common/component"
	"gotest.tools/v3/assert"
)

func TestEditCash(t *testing.T) {
	component := component.New()

	actual := component.EditCash(1127)

	expected := "/cashes/1127/edit"

	assert.Equal(t, actual, expected)
}
