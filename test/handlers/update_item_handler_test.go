package handlers_test

import (
	"testing"

	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/handlers"
	providers "github.com/tksasha/balance/internal/providers/test"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/services"
	"gotest.tools/v3/assert"
)

func TestUpdateItemHandler(t *testing.T) {
	dbNameProvider := providers.NewDBNameProvider()

	dbConnection := db.Open(dbNameProvider)

	itemRepository := repositories.NewItemRepository(dbConnection)

	itemService := services.NewItemService(itemRepository)

	updateItemHandler := handlers.NewUpdateItemHandler(itemService)

	_ = updateItemHandler

	assert.Assert(t, true)
}
