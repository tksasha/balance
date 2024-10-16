package models_test

import (
	"testing"
	"time"

	"github.com/tksasha/balance/internal/models"
	"gotest.tools/v3/assert"
)

func TestItem_GetIDAsString(t *testing.T) {
	item := models.Item{
		ID: 1155,
	}

	assert.Equal(t, item.GetIDAsString(), "1155")
}

func TestItem_GetDateAsString(t *testing.T) {
	t.Run("when date is zero it should return an empty string", func(t *testing.T) {
		item := models.Item{}

		assert.Equal(t, item.GetDateAsString(), "")
	})

	t.Run("when date is not zero it should return formatted date", func(t *testing.T) {
		item := models.Item{
			Date: time.Date(2024, 10, 9, 10, 8, 53, 0, time.UTC),
		}

		assert.Equal(t, item.GetDateAsString(), "2024-10-09")
	})
}

func TestItem_GetSumAsString(t *testing.T) {
	item := models.Item{
		Sum: 123456.78,
	}

	assert.Equal(t, item.GetSumAsString(), "123\u00A0456,78")
}
