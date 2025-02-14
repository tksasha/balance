// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"context"
	"github.com/tksasha/balance/internal/core/cash/handlers"
	"github.com/tksasha/balance/internal/core/cash/repository"
	"github.com/tksasha/balance/internal/core/cash/service"
	handlers2 "github.com/tksasha/balance/internal/core/category/handlers"
	repository2 "github.com/tksasha/balance/internal/core/category/repository"
	service2 "github.com/tksasha/balance/internal/core/category/service"
	handlers3 "github.com/tksasha/balance/internal/core/index/handler"
	handlers4 "github.com/tksasha/balance/internal/core/item/handlers"
	repository3 "github.com/tksasha/balance/internal/core/item/repository"
	service3 "github.com/tksasha/balance/internal/core/item/service"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/db/nameprovider"
	"github.com/tksasha/balance/internal/server"
	"github.com/tksasha/balance/internal/server/config"
	"github.com/tksasha/balance/internal/server/middlewares"
	"github.com/tksasha/balance/internal/server/routes"
)

// Injectors from wire.go:

func InitializeServer() *server.Server {
	configConfig := config.New()
	contextContext := context.Background()
	provider := nameprovider.New()
	sqlDB := db.Open(contextContext, provider)
	repositoryRepository := repository.New(sqlDB)
	serviceService := service.New(repositoryRepository)
	createHandler := handlers.NewCreateHandler(serviceService)
	deleteHandler := handlers.NewDeleteHandler(serviceService)
	editHandler := handlers.NewEditHandler(serviceService)
	indexHandler := handlers.NewIndexHandler(serviceService)
	newHandler := handlers.NewNewHandler()
	updateHandler := handlers.NewUpdateHandler(serviceService)
	repository4 := repository2.New(sqlDB)
	service4 := service2.New(repository4)
	handlersCreateHandler := handlers2.NewCreateHandler(service4)
	handlersDeleteHandler := handlers2.NewDeleteHandler(service4)
	handlersEditHandler := handlers2.NewEditHandler(service4)
	handlersIndexHandler := handlers2.NewIndexHandler(service4)
	handlersUpdateHandler := handlers2.NewUpdateHandler(service4)
	handler := handlers3.NewHandler(service4)
	repository5 := repository3.New(sqlDB)
	service5 := service3.New(repository5, repository4)
	createHandler2 := handlers4.NewCreateHandler(service5)
	editHandler2 := handlers4.NewEditHandler(service5)
	indexHandler2 := handlers4.NewIndexHandler(service5)
	updateHandler2 := handlers4.NewUpdateHandler(service5)
	routesRoutes := routes.New(createHandler, deleteHandler, editHandler, indexHandler, newHandler, updateHandler, handlersCreateHandler, handlersDeleteHandler, handlersEditHandler, handlersIndexHandler, handlersUpdateHandler, handler, createHandler2, editHandler2, indexHandler2, updateHandler2)
	v := middlewares.New()
	serverServer := server.New(configConfig, routesRoutes, v)
	return serverServer
}
