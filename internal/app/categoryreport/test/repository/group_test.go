package repository_test

import (
	"testing"

	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/app/categoryreport"
	"github.com/tksasha/balance/internal/app/categoryreport/repository"
	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/db/nameprovider"
	"gotest.tools/v3/assert"
)

func TestGroup(t *testing.T) {
	ctx := t.Context()

	db := db.Open(ctx, nameprovider.NewTestProvider())
	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	cleanup(t, db)

	repository := repository.New(db)

	createCategory(t, db, &category.Category{ID: 1, Name: "Food", Slug: "food"})
	createCategory(t, db, &category.Category{ID: 2, Name: "Beverages", Slug: "beverages"})

	createItem(t, db, &item.Item{Date: date(t, "2024-02-01"), CategoryID: 1, Sum: 11.11})
	createItem(t, db, &item.Item{Date: date(t, "2024-02-02"), CategoryID: 2, Sum: 22.22})
	createItem(t, db, &item.Item{Date: date(t, "2024-02-03"), CategoryID: 2, Sum: 33.33})
	createItem(t, db, &item.Item{Date: date(t, "2024-01-31"), CategoryID: 1, Sum: 11.11})

	filters := categoryreport.Filters{
		From: "2024-02-01",
		To:   "2024-02-29",
	}

	actual, err := repository.Group(ctx, filters)

	expected := categoryreport.Entities{
		{CategoryName: "Beverages", CategorySlug: "beverages", Sum: 55.55},
		{CategoryName: "Food", CategorySlug: "food", Sum: 11.11},
	}

	assert.NilError(t, err)

	for idx := range actual {
		assert.Equal(t, *actual[idx], *expected[idx])
	}
}
