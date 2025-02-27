package component_test

import (
	"bytes"
	"testing"

	"github.com/tksasha/balance/internal/app/item/component"
	"github.com/tksasha/validation"
	"gotest.tools/v3/golden"
)

func TestNew(t *testing.T) {
	component := component.New()

	errors := validation.Errors{}

	w := bytes.NewBuffer([]byte{})

	if err := component.New(item(t), categories(t), errors).Render(w); err != nil {
		t.Fatal(err)
	}

	golden.Assert(t, w.String(), "new.html")
}
