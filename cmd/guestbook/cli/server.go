package cli

import (
	"context"
	"log"
	"time"

	"github.com/spf13/cobra"
	"go.uber.org/fx"

	"github.com/hyeonjae/guestbook/internal/bootstrap"
	"github.com/hyeonjae/guestbook/internal/config"
	"github.com/hyeonjae/guestbook/internal/gateway"
)

func NewServerCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:     "server",
		Example: "./guestbook server --config config/local.yaml",
		Run: func(cmd *cobra.Command, args []string) {
			app := fx.New(
				fx.Provide(
					config.ParseFlags(cmd),
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
		},
	}

	flags := &config.Flags{}
	command.Flags().StringVarP(&flags.ConfigPath, "config", "c", "config/local.yaml", "Source directory to read from")

	return command
}

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
