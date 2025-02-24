//go:build wireinject

package wire

import (
	"context"

	"github.com/google/wire"
	"github.com/tksasha/balance/internal/app/cash"
	cashcomponents "github.com/tksasha/balance/internal/app/cash/components"
	cashhandlers "github.com/tksasha/balance/internal/app/cash/handlers"
	cashrepository "github.com/tksasha/balance/internal/app/cash/repository"
	cashservice "github.com/tksasha/balance/internal/app/cash/service"
	"github.com/tksasha/balance/internal/app/category"
	categorycomponents "github.com/tksasha/balance/internal/app/category/components"
	categoryhandlers "github.com/tksasha/balance/internal/app/category/handlers"
	categoryrepository "github.com/tksasha/balance/internal/app/category/repository"
	categoryservice "github.com/tksasha/balance/internal/app/category/service"
	"github.com/tksasha/balance/internal/app/index"
	indexcomponents "github.com/tksasha/balance/internal/app/index/components"
	indexhandler "github.com/tksasha/balance/internal/app/index/handler"
	indexrepository "github.com/tksasha/balance/internal/app/index/repository"
	indexservice "github.com/tksasha/balance/internal/app/index/service"
	"github.com/tksasha/balance/internal/app/item"
	itemcomponents "github.com/tksasha/balance/internal/app/item/components"
	itemhandlers "github.com/tksasha/balance/internal/app/item/handlers"
	itemrepository "github.com/tksasha/balance/internal/app/item/repository"
	itemservice "github.com/tksasha/balance/internal/app/item/service"
	backofficeCategory "github.com/tksasha/balance/internal/backoffice/category"
	backofficeCategoryComponent "github.com/tksasha/balance/internal/backoffice/category/component"
	backofficeCategoryHandlers "github.com/tksasha/balance/internal/backoffice/category/handlers"
	backofficeCategoryRepository "github.com/tksasha/balance/internal/backoffice/category/repository"
	backofficeCategoryService "github.com/tksasha/balance/internal/backoffice/category/service"
	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/balance/internal/common/component"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/db/nameprovider"
	"github.com/tksasha/balance/internal/server"
	"github.com/tksasha/balance/internal/server/config"
	"github.com/tksasha/balance/internal/server/middlewares"
	"github.com/tksasha/balance/internal/server/routes"
)

func InitializeServer() *server.Server {
	wire.Build(
		backofficeCategoryComponent.New,
		backofficeCategoryHandlers.NewCreateHandler,
		backofficeCategoryHandlers.NewListHandler,
		backofficeCategoryRepository.New,
		backofficeCategoryService.New,
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
		wire.Bind(new(backofficeCategory.Repository), new(*backofficeCategoryRepository.Repository)),
		wire.Bind(new(backofficeCategory.Service), new(*backofficeCategoryService.Service)),
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
