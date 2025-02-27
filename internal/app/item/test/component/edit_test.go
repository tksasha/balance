package component_test

import (
	"bytes"
	"testing"

	"github.com/tksasha/balance/internal/app/item/component"
	"github.com/tksasha/validation"
	"gotest.tools/v3/golden"
)

func TestEdit(t *testing.T) {
	component := component.New()

	errors := validation.Errors{}

	w := bytes.NewBuffer([]byte{})

	if err := component.Edit(item(t), categories(t), errors).Render(w); err != nil {
		t.Fatal(err)
	}

	golden.Assert(t, w.String(), "edit.html")
}
