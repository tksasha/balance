package repository_test

import (
	"testing"

	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/app/category/repository"
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

	repository := repository.New(db)

	t.Run("return filtered categories", func(t *testing.T) {
		cleanup(t, db)

		createCategory(t, db, &category.Category{ID: 1, Name: "Food", Slug: "food", Supercategory: 1})
		createCategory(t, db, &category.Category{ID: 2, Name: "Beverages", Slug: "beverages", Supercategory: 1})
		createCategory(t, db, &category.Category{ID: 3, Name: "Salary", Slug: "salary", Income: true})
		createCategory(t, db, &category.Category{ID: 4, Name: "AFU", Slug: "afu", Supercategory: 2})

		createItem(t, db, &item.Item{Date: date(t, "2024-02-01"), CategoryID: 1, Sum: 11.11})
		createItem(t, db, &item.Item{Date: date(t, "2024-02-02"), CategoryID: 2, Sum: 22.22})
		createItem(t, db, &item.Item{Date: date(t, "2024-02-03"), CategoryID: 2, Sum: 33.33})
		createItem(t, db, &item.Item{Date: date(t, "2024-01-31"), CategoryID: 1, Sum: 11.11})
		createItem(t, db, &item.Item{Date: date(t, "2024-02-04"), CategoryID: 3, Sum: 44.44})
		createItem(t, db, &item.Item{Date: date(t, "2024-02-05"), CategoryID: 4, Sum: 55.55})

		filters := category.Filters{
			From: "2024-02-01",
			To:   "2024-02-29",
		}

		actual, err := repository.FindAllByFilters(ctx, filters)

		expected := category.Categories{
			{Name: "Food", Slug: "food", Sum: 11.11, Supercategory: 1},
			{Name: "Beverages", Slug: "beverages", Sum: 55.55, Supercategory: 1},
			{Name: "Salary", Slug: "salary", Sum: 44.44, Supercategory: 0},
			{Name: "AFU", Slug: "afu", Sum: 55.55, Supercategory: 2},
		}

		assert.NilError(t, err)

		for idx := range actual {
			assert.Equal(t, *actual[idx], *expected[idx])
		}
	})
}
