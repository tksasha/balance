// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"context"
	handler2 "github.com/tksasha/balance/internal/app/balance/handler"
	repository3 "github.com/tksasha/balance/internal/app/balance/repository"
	service3 "github.com/tksasha/balance/internal/app/balance/service"
	handlers3 "github.com/tksasha/balance/internal/app/cash/handlers"
	repository4 "github.com/tksasha/balance/internal/app/cash/repository"
	service4 "github.com/tksasha/balance/internal/app/cash/service"
	handlers4 "github.com/tksasha/balance/internal/app/category/handlers"
	repository5 "github.com/tksasha/balance/internal/app/category/repository"
	service5 "github.com/tksasha/balance/internal/app/category/service"
	handler3 "github.com/tksasha/balance/internal/app/index/handler"
	handlers5 "github.com/tksasha/balance/internal/app/item/handlers"
	repository6 "github.com/tksasha/balance/internal/app/item/repository"
	service6 "github.com/tksasha/balance/internal/app/item/service"
	"github.com/tksasha/balance/internal/backoffice/cash/handlers"
	"github.com/tksasha/balance/internal/backoffice/cash/repository"
	"github.com/tksasha/balance/internal/backoffice/cash/service"
	handlers2 "github.com/tksasha/balance/internal/backoffice/category/handlers"
	repository2 "github.com/tksasha/balance/internal/backoffice/category/repository"
	service2 "github.com/tksasha/balance/internal/backoffice/category/service"
	"github.com/tksasha/balance/internal/backoffice/index/handler"
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
	nameProvider := nameprovider.New()
	sqlDB := db.Open(contextContext, nameProvider)
	repositoryRepository := repository.New(sqlDB)
	serviceService := service.New(repositoryRepository)
	createHandler := handlers.NewCreateHandler(serviceService)
	deleteHandler := handlers.NewDeleteHandler(serviceService)
	editHandler := handlers.NewEditHandler(serviceService)
	indexHandler := handlers.NewIndexHandler(serviceService)
	newHandler := handlers.NewNewHandler()
	updateHandler := handlers.NewUpdateHandler(serviceService)
	repository7 := repository2.New(sqlDB)
	service7 := service2.New(repository7)
	handlersCreateHandler := handlers2.NewCreateHandler(service7)
	handlersEditHandler := handlers2.NewEditHandler(service7)
	handlersIndexHandler := handlers2.NewIndexHandler(service7)
	handlersNewHandler := handlers2.NewNewHandler()
	handlersUpdateHandler := handlers2.NewUpdateHandler(service7)
	handlerIndexHandler := handler.NewIndexHandler()
	repository8 := repository3.New(sqlDB)
	service8 := service3.New(repository8)
	showHandler := handler2.NewShowHandler(service8)
	repository9 := repository4.New(sqlDB)
	service9 := service4.New(repository9)
	editHandler2 := handlers3.NewEditHandler(service9)
	indexHandler2 := handlers3.NewIndexHandler(service9)
	updateHandler2 := handlers3.NewUpdateHandler(service9)
	repository10 := repository5.New(sqlDB)
	service10 := service5.New(repository10)
	indexHandler3 := handlers4.NewIndexHandler(service10)
	handlerHandler := handler3.New()
	repository11 := repository6.New(sqlDB)
	service11 := service6.New(repository11, repository10)
	createHandler2 := handlers5.NewCreateHandler(service11, service10)
	handlersDeleteHandler := handlers5.NewDeleteHandler(service11)
	editHandler3 := handlers5.NewEditHandler(service11, service10)
	indexHandler4 := handlers5.NewIndexHandler(service11)
	newHandler2 := handlers5.NewNewHandler(service10)
	updateHandler3 := handlers5.NewUpdateHandler(service11, service10)
	routesRoutes := routes.New(createHandler, deleteHandler, editHandler, indexHandler, newHandler, updateHandler, handlersCreateHandler, handlersEditHandler, handlersIndexHandler, handlersNewHandler, handlersUpdateHandler, handlerIndexHandler, showHandler, editHandler2, indexHandler2, updateHandler2, indexHandler3, handlerHandler, createHandler2, handlersDeleteHandler, editHandler3, indexHandler4, newHandler2, updateHandler3)
	v := middlewares.New()
	serverServer := server.New(configConfig, routesRoutes, v)
	return serverServer
}
