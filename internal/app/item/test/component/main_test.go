package component_test

import (
	"testing"
	"time"

	"github.com/tksasha/balance/internal/app/category"
	model "github.com/tksasha/balance/internal/app/item"
)

func categories(t *testing.T) category.Categories {
	t.Helper()

	return category.Categories{
		{ID: 1045, Name: "Beverages", Income: false},
		{ID: 1046, Name: "Food", Income: false},
		{ID: 1048, Name: "Rent", Income: false},
		{ID: 1049, Name: "Salary", Income: true},
		{ID: 1050, Name: "Cashback", Income: true},
	}
}

func item(t *testing.T) *model.Item {
	t.Helper()

	return &model.Item{
		ID:      1027,
		Date:    time.Date(2025, 2, 27, 0, 0, 0, 0, time.UTC),
		Formula: "3+4",
	}
}
