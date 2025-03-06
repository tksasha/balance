//go:build wireinject

package wire

import (
	"context"

	"github.com/google/wire"
	"github.com/tksasha/balance/internal/app/balance"
	balancehandler "github.com/tksasha/balance/internal/app/balance/handler"
	balancerepository "github.com/tksasha/balance/internal/app/balance/repository"
	balanceservice "github.com/tksasha/balance/internal/app/balance/service"
	"github.com/tksasha/balance/internal/app/cash"
	cashhandlers "github.com/tksasha/balance/internal/app/cash/handlers"
	cashrepository "github.com/tksasha/balance/internal/app/cash/repository"
	cashservice "github.com/tksasha/balance/internal/app/cash/service"
	"github.com/tksasha/balance/internal/app/category"
	categoryhandlers "github.com/tksasha/balance/internal/app/category/handlers"
	categoryrepository "github.com/tksasha/balance/internal/app/category/repository"
	categoryservice "github.com/tksasha/balance/internal/app/category/service"
	indexhandler "github.com/tksasha/balance/internal/app/index/handler"
	"github.com/tksasha/balance/internal/app/item"
	itemhandlers "github.com/tksasha/balance/internal/app/item/handlers"
	itemrepository "github.com/tksasha/balance/internal/app/item/repository"
	itemservice "github.com/tksasha/balance/internal/app/item/service"
	backofficecash "github.com/tksasha/balance/internal/backoffice/cash"
	backofficecashhandlers "github.com/tksasha/balance/internal/backoffice/cash/handlers"
	backofficecashrepository "github.com/tksasha/balance/internal/backoffice/cash/repository"
	backofficecashservice "github.com/tksasha/balance/internal/backoffice/cash/service"
	backofficecategory "github.com/tksasha/balance/internal/backoffice/category"
	backofficecategoryhandlers "github.com/tksasha/balance/internal/backoffice/category/handlers"
	backofficecategoryrepository "github.com/tksasha/balance/internal/backoffice/category/repository"
	backofficecategoryservice "github.com/tksasha/balance/internal/backoffice/category/service"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/db/nameprovider"
	"github.com/tksasha/balance/internal/server"
	"github.com/tksasha/balance/internal/server/config"
	"github.com/tksasha/balance/internal/server/middlewares"
	"github.com/tksasha/balance/internal/server/routes"
)

func InitializeServer() *server.Server {
	wire.Build(
		backofficecashhandlers.NewCreateHandler,
		backofficecashhandlers.NewDeleteHandler,
		backofficecashhandlers.NewEditHandler,
		backofficecashhandlers.NewListHandler,
		backofficecashhandlers.NewNewHandler,
		backofficecashhandlers.NewUpdateHandler,
		backofficecashrepository.New,
		backofficecashservice.New,
		backofficecategoryhandlers.NewCreateHandler,
		backofficecategoryhandlers.NewDeleteHandler,
		backofficecategoryhandlers.NewEditHandler,
		backofficecategoryhandlers.NewListHandler,
		backofficecategoryhandlers.NewUpdateHandler,
		backofficecategoryrepository.New,
		backofficecategoryservice.New,
		balancehandler.NewShowHandler,
		balancerepository.New,
		balanceservice.New,
		cashhandlers.NewEditHandler,
		cashhandlers.NewIndexHandler,
		cashhandlers.NewUpdateHandler,
		cashrepository.New,
		cashservice.New,
		categoryhandlers.NewIndexHandler,
		categoryrepository.New,
		categoryservice.New,
		config.New,
		context.Background,
		db.Open,
		indexhandler.New,
		itemhandlers.NewCreateHandler,
		itemhandlers.NewEditHandler,
		itemhandlers.NewIndexHandler,
		itemhandlers.NewUpdateHandler,
		itemrepository.New,
		itemservice.New,
		middlewares.New,
		nameprovider.New,
		routes.New,
		server.New,
		wire.Bind(new(backofficecash.Repository), new(*backofficecashrepository.Repository)),
		wire.Bind(new(backofficecash.Service), new(*backofficecashservice.Service)),
		wire.Bind(new(backofficecategory.Repository), new(*backofficecategoryrepository.Repository)),
		wire.Bind(new(backofficecategory.Service), new(*backofficecategoryservice.Service)),
		wire.Bind(new(balance.Repository), new(*balancerepository.Repository)),
		wire.Bind(new(balance.Service), new(*balanceservice.Service)),
		wire.Bind(new(cash.Repository), new(*cashrepository.Repository)),
		wire.Bind(new(cash.Service), new(*cashservice.Service)),
		wire.Bind(new(category.Repository), new(*categoryrepository.Repository)),
		wire.Bind(new(category.Service), new(*categoryservice.Service)),
		wire.Bind(new(db.NameProvider), new(*nameprovider.NameProvider)),
		wire.Bind(new(item.CategoryRepository), new(*categoryrepository.Repository)),
		wire.Bind(new(item.CategoryService), new(*categoryservice.Service)),
		wire.Bind(new(item.Repository), new(*itemrepository.Repository)),
		wire.Bind(new(item.Service), new(*itemservice.Service)),
	)

	return nil
}
