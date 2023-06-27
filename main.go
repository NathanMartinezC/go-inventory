package main

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/nathanmartinezc/go-inventory/database"
	"github.com/nathanmartinezc/go-inventory/internal/api"
	"github.com/nathanmartinezc/go-inventory/internal/repository"
	"github.com/nathanmartinezc/go-inventory/internal/service"
	"github.com/nathanmartinezc/go-inventory/settings"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
			repository.New,
			service.New,
			api.New,
			echo.New,
		),
		fx.Invoke(
			setLifeCycle,
		),
	)
	app.Run()
}

func setLifeCycle(lc fx.Lifecycle, a *api.API, s *settings.Settings, e *echo.Echo) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			address := fmt.Sprintf(":%s", s.Port)
			go a.Start(e, address)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
