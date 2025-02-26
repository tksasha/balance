package component_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/tksasha/balance/internal/app/item/component"
	"gotest.tools/v3/golden"
)

func TestDate(t *testing.T) {
	component := component.New()

	date := time.Date(2025, 2, 26, 0, 0, 0, 0, time.UTC)

	t.Run("when date is invalid", func(t *testing.T) {
		w := bytes.NewBuffer([]byte{})

		message := "date: is required"

		if err := component.Date(date, &message).Render(w); err != nil {
			t.Fatal(err)
		}

		golden.Assert(t, w.String(), "date-with-errors.html")
	})

	t.Run("when date is valid", func(t *testing.T) {
		w := bytes.NewBuffer([]byte{})

		if err := component.Date(date, nil).Render(w); err != nil {
			t.Fatal(err)
		}

		golden.Assert(t, w.String(), "date.html")
	})
}
