package repository_test

import (
	"testing"

	"github.com/shopspring/decimal"
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

	createCategory(t, db, &category.Category{ID: 1, Name: "Food", Slug: "food", Supercategory: 1})
	createCategory(t, db, &category.Category{ID: 2, Name: "Beverages", Slug: "beverages", Supercategory: 1})
	createCategory(t, db, &category.Category{ID: 3, Name: "Salary", Slug: "salary", Income: true})
	createCategory(t, db, &category.Category{ID: 4, Name: "AFU", Slug: "afu", Supercategory: 2})

	createItem(t, db, &item.Item{Date: date(t, "2024-02-01"), CategoryID: 1, Sum: dec(t, 11.11)})
	createItem(t, db, &item.Item{Date: date(t, "2024-02-02"), CategoryID: 2, Sum: dec(t, 22.22)})
	createItem(t, db, &item.Item{Date: date(t, "2024-02-03"), CategoryID: 2, Sum: dec(t, 33.33)})
	createItem(t, db, &item.Item{Date: date(t, "2024-01-31"), CategoryID: 1, Sum: dec(t, 11.11)})
	createItem(t, db, &item.Item{Date: date(t, "2024-02-04"), CategoryID: 3, Sum: dec(t, 44.44)})
	createItem(t, db, &item.Item{Date: date(t, "2024-02-05"), CategoryID: 4, Sum: dec(t, 55.55)})

	filters := categoryreport.Filters{
		From: "2024-02-01",
		To:   "2024-02-29",
	}

	actual, err := repository.Group(ctx, filters)

	expected := categoryreport.Entities{
		{CategoryName: "Food", CategorySlug: "food", Sum: dec(t, 11.11), Supercategory: 1},
		{CategoryName: "Beverages", CategorySlug: "beverages", Sum: dec(t, 55.55), Supercategory: 1},
		{CategoryName: "Salary", CategorySlug: "salary", Sum: dec(t, 44.44), Supercategory: 0},
		{CategoryName: "AFU", CategorySlug: "afu", Sum: dec(t, 55.55), Supercategory: 2},
	}

	assert.NilError(t, err)

	for idx := range actual {
		assert.Equal(t, actual[idx].CategoryName, expected[idx].CategoryName)
		assert.Equal(t, actual[idx].CategorySlug, expected[idx].CategorySlug)
		assert.Assert(t, actual[idx].Sum.Equal(expected[idx].Sum))
		assert.Equal(t, actual[idx].Supercategory, expected[idx].Supercategory)
	}
}

func dec(t *testing.T, f float64) decimal.Decimal {
	t.Helper()

	return decimal.NewFromFloat(f)
}
