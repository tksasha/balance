package server

import (
	"net/http"

	"github.com/tksasha/balance/internal/config"
	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/interfaces"
	"github.com/tksasha/balance/internal/middlewares"
	"github.com/tksasha/balance/internal/repositories"
	"go.uber.org/fx"
)

func Run() {
	fx.New(
		fx.Provide(
			fx.Annotate(
				NewHTTPServer,
				fx.ParamTags("", "", "", `group:"middlewares"`),
			),
			config.New,
			fx.Annotate(
				NewServeMux,
				fx.ParamTags(`group:"routes"`),
			),
			fx.Annotate(
				handlers.NewIndexPageHandler,
				fx.As(new(interfaces.Route)),
				fx.ResultTags(`group:"routes"`),
			),
			fx.Annotate(
				handlers.NewCreateItemHandler,
				fx.As(new(interfaces.Route)),
				fx.ResultTags(`group:"routes"`),
			),
			fx.Annotate(
				middlewares.NewCurrencyMiddleware,
				fx.As(new(interfaces.Middleware)),
				fx.ResultTags(`group:"middlewares"`),
			),
			fx.Annotate(
				repositories.NewItemRepository,
				fx.As(new(repositories.ItemCreator)),
			),
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
