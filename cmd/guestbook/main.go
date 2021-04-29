package main

import (
	"context"
	"log"
	"time"

	"go.uber.org/fx"

	"github.com/daangn/guestbook/cmd/guestbook/cli"
	"github.com/daangn/guestbook/internal/bootstrap"
	"github.com/daangn/guestbook/internal/config"
	"github.com/daangn/guestbook/internal/gateway"
)

func serve(lc fx.Lifecycle, guestbook *bootstrap.Guestbook) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go guestbook.Serve()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}

func main() {
	app := fx.New(
		fx.Provide(
			cli.ParseFlags,
			config.NewConfig,
			gateway.NewGuestbook,
			bootstrap.NewGuestbook,
		),
		fx.Invoke(serve),
	)

	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}

	<-app.Done()
}
