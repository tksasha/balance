package component_test

import (
	"bytes"
	"testing"

	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/balance/internal/app/cash/component"
	"github.com/tksasha/validation"
	"gotest.tools/v3/golden"
)

func TestEdit(t *testing.T) {
	component := component.New()

	t.Run("when cash is invalid", func(t *testing.T) {
		cash := &cash.Cash{ID: 1354}

		errors := validation.Errors{}

		errors.Set("name", "is required")
		errors.Set("sum", "can not be zero")

		w := bytes.NewBuffer([]byte{})

		if err := component.Edit(cash, errors).Render(w); err != nil {
			t.Fatal(err)
		}

		golden.Assert(t, w.String(), "edit-with-errors.html")
	})

	t.Run("when cash is valid", func(t *testing.T) {
		cash := &cash.Cash{
			ID:   1531,
			Name: "Stocks",
			Sum:  15.32,
		}

		w := bytes.NewBuffer([]byte{})

		if err := component.Edit(cash, nil).Render(w); err != nil {
			t.Fatal(err)
		}

		golden.Assert(t, w.String(), "edit.html")
	})
}
