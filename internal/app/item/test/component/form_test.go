package component_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/app/item/component"
	"github.com/tksasha/validation"
	"gotest.tools/v3/golden"
)

func TestForm(t *testing.T) {
	component := component.New()

	errors := validation.Errors{}

	categories := category.Categories{
		{ID: 1045, Name: "Beverages", Income: false},
		{ID: 1046, Name: "Food", Income: false},
		{ID: 1048, Name: "Rent", Income: false},
		{ID: 1049, Name: "Salary", Income: true},
		{ID: 1050, Name: "Cashback", Income: true},
	}

	t.Run("when create item", func(t *testing.T) {
		item := &item.Item{
			ID:      0,
			Date:    time.Date(2025, 2, 27, 0, 0, 0, 0, time.UTC),
			Formula: "3+4",
		}

		w := bytes.NewBuffer([]byte{})

		if err := component.Form(item, categories, errors).Render(w); err != nil {
			t.Fatal(err)
		}

		golden.Assert(t, w.String(), "form-create.html")
	})

	t.Run("when update item", func(t *testing.T) {
		item := &item.Item{
			ID:      1027,
			Date:    time.Date(2025, 2, 27, 0, 0, 0, 0, time.UTC),
			Formula: "3+4",
		}

		w := bytes.NewBuffer([]byte{})

		if err := component.Form(item, categories, errors).Render(w); err != nil {
			t.Fatal(err)
		}

		golden.Assert(t, w.String(), "form-update.html")
	})
}
