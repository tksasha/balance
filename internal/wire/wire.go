//go:build wireinject

package wire

import (
	"context"

	"github.com/google/wire"
	"github.com/tksasha/balance/internal/core/cash"
	cashcomponents "github.com/tksasha/balance/internal/core/cash/components"
	cashhandlers "github.com/tksasha/balance/internal/core/cash/handlers"
	cashrepository "github.com/tksasha/balance/internal/core/cash/repository"
	cashservice "github.com/tksasha/balance/internal/core/cash/service"
	"github.com/tksasha/balance/internal/core/category"
	categorycomponents "github.com/tksasha/balance/internal/core/category/components"
	categoryhandlers "github.com/tksasha/balance/internal/core/category/handlers"
	categoryrepository "github.com/tksasha/balance/internal/core/category/repository"
	categoryservice "github.com/tksasha/balance/internal/core/category/service"
	"github.com/tksasha/balance/internal/core/common"
	"github.com/tksasha/balance/internal/core/common/component"
	"github.com/tksasha/balance/internal/core/index"
	indexcomponents "github.com/tksasha/balance/internal/core/index/components"
	indexhandler "github.com/tksasha/balance/internal/core/index/handler"
	indexrepository "github.com/tksasha/balance/internal/core/index/repository"
	indexservice "github.com/tksasha/balance/internal/core/index/service"
	"github.com/tksasha/balance/internal/core/item"
	itemcomponents "github.com/tksasha/balance/internal/core/item/components"
	itemhandlers "github.com/tksasha/balance/internal/core/item/handlers"
	itemrepository "github.com/tksasha/balance/internal/core/item/repository"
	itemservice "github.com/tksasha/balance/internal/core/item/service"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/db/nameprovider"
	"github.com/tksasha/balance/internal/server"
	"github.com/tksasha/balance/internal/server/config"
	"github.com/tksasha/balance/internal/server/middlewares"
	"github.com/tksasha/balance/internal/server/routes"
)

func InitializeServer() *server.Server {
	wire.Build(
		cashcomponents.NewCashComponent,
		cashhandlers.NewCreateHandler,
		cashhandlers.NewDeleteHandler,
		cashhandlers.NewEditHandler,
		cashhandlers.NewListHandler,
		cashhandlers.NewNewHandler,
		cashhandlers.NewUpdateHandler,
		cashrepository.New,
		cashservice.New,
		categorycomponents.NewCategoryComponent,
		categoryhandlers.NewCreateHandler,
		categoryhandlers.NewDeleteHandler,
		categoryhandlers.NewEditHandler,
		categoryhandlers.NewListHandler,
		categoryhandlers.NewUpdateHandler,
		categoryrepository.New,
		categoryservice.New,
		common.NewBaseHandler,
		common.NewBaseRepository,
		common.NewBaseService,
		component.New,
		config.New,
		context.Background,
		db.Open,
		indexcomponents.NewIndexComponent,
		indexcomponents.NewMonthsComponent,
		indexcomponents.NewYearsComponent,
		indexhandler.New,
		indexrepository.New,
		indexservice.New,
		itemcomponents.NewItemsComponent,
		itemhandlers.NewCreateHandler,
		itemhandlers.NewEditHandler,
		itemhandlers.NewListHandler,
		itemhandlers.NewUpdateHandler,
		itemrepository.New,
		itemservice.New,
		middlewares.New,
		nameprovider.New,
		routes.New,
		server.New,
		wire.Bind(new(cash.Repository), new(*cashrepository.Repository)),
		wire.Bind(new(cash.Service), new(*cashservice.Service)),
		wire.Bind(new(category.Repository), new(*categoryrepository.Repository)),
		wire.Bind(new(category.Service), new(*categoryservice.Service)),
		wire.Bind(new(db.NameProvider), new(*nameprovider.Provider)),
		wire.Bind(new(index.Repository), new(*indexrepository.Repository)),
		wire.Bind(new(index.Service), new(*indexservice.Service)),
		wire.Bind(new(item.Repository), new(*itemrepository.Repository)),
		wire.Bind(new(item.Service), new(*itemservice.Service)),
	)

	return nil
}
