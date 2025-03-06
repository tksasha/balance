package component_test

import (
	"bytes"
	"testing"

	"github.com/tksasha/balance/internal/app/item/component"
	"gotest.tools/v3/golden"
)

func TestNew(t *testing.T) {
	component := component.New()

	w := bytes.NewBuffer([]byte{})

	if err := component.New(categories(t)).Render(w); err != nil {
		t.Fatal(err)
	}

	golden.Assert(t, w.String(), "new.html")
}
