// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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
	cashRepository := repositories.NewCashRepository(dbDB)
	cashService := services.NewCashService(cashRepository)
	cashCreateHandler := handlers.NewCashCreateHandler(cashService)
	cashEditHandler := handlers.NewCashEditHandler(cashService)
	categoryRepository := repositories.NewCategoryRepository(dbDB)
	categoryService := services.NewCategoryService(categoryRepository)
	createCategoryHandler := handlers.NewCreateCategoryHandler(categoryService)
	itemRepository := repositories.NewItemRepository(dbDB)
	itemService := services.NewItemService(itemRepository, categoryRepository)
	createItemHandler := handlers.NewCreateItemHandler(itemService)
	editCategoryHandler := handlers.NewEditCategoryHandler(categoryService)
	getCategoriesHandler := handlers.NewGetCategoriesHandler(categoryService)
	getItemHandler := handlers.NewGetItemHandler(itemService)
	getItemsHandler := handlers.NewGetItemsHandler(itemService)
	indexPageHandler := handlers.NewIndexPageHandler(categoryService)
	updateCategoryHandler := handlers.NewUpdateCategoryHandler(categoryService)
	routesRoutes := routes.New(cashCreateHandler, cashEditHandler, createCategoryHandler, createItemHandler, editCategoryHandler, getCategoriesHandler, getItemHandler, getItemsHandler, indexPageHandler, updateCategoryHandler)
	serverServer := server.New(configConfig, routesRoutes)
	return serverServer
}
