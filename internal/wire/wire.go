//go:build wireinject

package wire

import (
	"context"

	"github.com/google/wire"
	"github.com/tksasha/balance/internal/core/cash"
	cashhandlers "github.com/tksasha/balance/internal/core/cash/handlers"
	cashrepository "github.com/tksasha/balance/internal/core/cash/repository"
	cashservice "github.com/tksasha/balance/internal/core/cash/service"
	"github.com/tksasha/balance/internal/core/category"
	categoryhandlers "github.com/tksasha/balance/internal/core/category/handlers"
	categoryrepository "github.com/tksasha/balance/internal/core/category/repository"
	categoryservice "github.com/tksasha/balance/internal/core/category/service"
	commoncomponents "github.com/tksasha/balance/internal/core/common/components"
	"github.com/tksasha/balance/internal/core/common/helpers"
	"github.com/tksasha/balance/internal/core/common/providers"
	"github.com/tksasha/balance/internal/core/common/valueobjects"
	"github.com/tksasha/balance/internal/core/index"
	indexcomponents "github.com/tksasha/balance/internal/core/index/components"
	indexhandlers "github.com/tksasha/balance/internal/core/index/handlers"
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
		helpers.New,
		providers.NewTimeProvider,
		commoncomponents.NewBaseComponent,
		cashhandlers.NewCreateHandler,
		cashhandlers.NewDeleteHandler,
		cashhandlers.NewEditHandler,
		cashhandlers.NewIndexHandler,
		cashhandlers.NewNewHandler,
		cashhandlers.NewUpdateHandler,
		cashrepository.New,
		cashservice.New,
		categoryhandlers.NewCreateHandler,
		categoryhandlers.NewDeleteHandler,
		categoryhandlers.NewEditHandler,
		categoryhandlers.NewIndexHandler,
		categoryhandlers.NewUpdateHandler,
		categoryrepository.New,
		categoryservice.New,
		indexcomponents.NewIndexPageComponent,
		indexcomponents.NewMonthsComponent,
		itemcomponents.NewItemsComponent,
		config.New,
		context.Background,
		db.Open,
		indexhandlers.NewIndexHandler,
		indexrepository.New,
		indexservice.New,
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
		wire.Bind(new(cash.Repository), new(*cashrepository.Repository)),
		wire.Bind(new(cash.Service), new(*cashservice.Service)),
		wire.Bind(new(category.Repository), new(*categoryrepository.Repository)),
		wire.Bind(new(category.Service), new(*categoryservice.Service)),
		wire.Bind(new(db.NameProvider), new(*nameprovider.Provider)),
		wire.Bind(new(index.Repository), new(*indexrepository.Repository)),
		wire.Bind(new(index.Service), new(*indexservice.Service)),
		wire.Bind(new(item.Repository), new(*itemrepository.Repository)),
		wire.Bind(new(item.Service), new(*itemservice.Service)),
		wire.Bind(new(valueobjects.CurrentDateProvider), new(*providers.TimeProvider)),
	)

	return nil
}
