// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject

package wire

import (
	"github.com/tksasha/balance/internal/config"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/providers"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/routes"
	"github.com/tksasha/balance/internal/server"
	"github.com/tksasha/balance/internal/services"
)

// Injectors from wire.go:

func InitializeServer() *server.Server {
	configConfig := config.New()
	dbNameProvider := providers.NewDBNameProvider()
	dbDB := db.Open(dbNameProvider)
	categoryRepository := repositories.NewCategoryRepository(dbDB)
	categoryService := services.NewCategoryService(categoryRepository)
	indexPageHandler := handlers.NewIndexPageHandler(categoryService)
	itemRepository := repositories.NewItemRepository(dbDB)
	createItemService := services.NewCreateItemService(itemRepository, categoryRepository)
	createItemHandler := handlers.NewCreateItemHandler(createItemService)
	itemService := services.NewItemService(itemRepository, categoryRepository)
	getItemsHandler := handlers.NewGetItemsHandler(itemService)
	getItemHandler := handlers.NewGetItemHandler(itemService)
	createCategoryHandler := handlers.NewCreateCategoryHandler(categoryService)
	getCategoriesHandler := handlers.NewGetCategoriesHandler(categoryService)
	editCategoryHandler := handlers.NewEditCategoryHandler(categoryService)
	updateCategoryHandler := handlers.NewUpdateCategoryHandler(categoryService)
	routesRoutes := routes.New(indexPageHandler, createItemHandler, getItemsHandler, getItemHandler, createCategoryHandler, getCategoriesHandler, editCategoryHandler, updateCategoryHandler)
	serverServer := server.New(configConfig, routesRoutes)
	return serverServer
}
