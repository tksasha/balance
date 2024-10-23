package server

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"net/http"

	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/middlewares"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/server/config"
	"github.com/tksasha/balance/internal/server/db"
	"github.com/tksasha/balance/internal/services"
	"go.uber.org/fx"
)

func New(
	lifecycle fx.Lifecycle,
	config *config.Config,
	mux *http.ServeMux,
	middlewares []middlewares.Middleware,
) *http.Server {
	handler := http.Handler(mux)

	for _, middleware := range middlewares {
		handler = middleware.Wrap(handler)
	}

	server := &http.Server{
		Addr:              config.Address,
		ReadHeaderTimeout: config.ReadHeaderTimeout,
		Handler:           handler,
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			listener, err := net.Listen("tcp", server.Addr)
			if err != nil {
				slog.Error("failed to create listener", "error", err)

				return err
			}

			go func() {
				err := server.Serve(listener)
				if err != nil && !errors.Is(err, http.ErrServerClosed) {
					slog.Error("server error", "error", err)
				}
			}()

			slog.Info("server started")

			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := server.Shutdown(ctx); err != nil {
				slog.Error("failed to shutdown server", "error", err)

				if err := server.Close(); err != nil {
					slog.Error("failed to close server", "error", err)

					return err
				}

				slog.Info("server closed")
			}

			slog.Info("server shutdown")

			return nil
		},
	})

	return server
}

func Run() {
	fx.New(
		fx.Provide(
			config.New,
			db.Open,
			fx.Annotate(
				New,
				fx.ParamTags("", "", "", `group:"middlewares"`),
			),
			fx.Annotate(
				NewServeMux,
				fx.ParamTags(`group:"routes"`),
			),
			fx.Annotate(
				handlers.NewIndexPageHandler,
				fx.As(new(handlers.Route)),
				fx.ResultTags(`group:"routes"`),
			),
			fx.Annotate(
				handlers.NewCreateItemHandler,
				fx.As(new(handlers.Route)),
				fx.ResultTags(`group:"routes"`),
			),
			fx.Annotate(
				handlers.NewGetCategoriesHandler,
				fx.As(new(handlers.Route)),
				fx.ResultTags(`group:"routes"`),
			),
			fx.Annotate(
				middlewares.NewCurrencyMiddleware,
				fx.As(new(middlewares.Middleware)),
				fx.ResultTags(`group:"middlewares"`),
			),
			fx.Annotate(
				repositories.NewItemRepository,
				fx.As(new(repositories.ItemRepository)),
			),
			fx.Annotate(
				services.NewCategoryService,
				fx.As(new(services.CategoryService)),
			),
			fx.Annotate(
				repositories.NewCategoryRepository,
				fx.As(new(repositories.CategoryRepository)),
			),
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
