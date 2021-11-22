package cmd

import (
	"context"

	"github.com/fremantle-industries/tabletop/pkg/apps"
	"github.com/spf13/cobra"

	"github.com/fremantle-industries/workshop/apps/config"
	"github.com/fremantle-industries/workshop/apps/deliver"
	"github.com/fremantle-industries/workshop/apps/marketdata"
	"github.com/fremantle-industries/workshop/apps/orderbook"
	"github.com/fremantle-industries/workshop/apps/process"
	"github.com/fremantle-industries/workshop/pkg"
)

var (
	rootCmd = &cobra.Command{
		Use: "workshop",
		Long: `CLI interface for workshop. Create, manage & improve your automated trading strategies
with rich and diverse data sets, a first class local development experience and a progression
story for deployment across clouds.

workshop is a tabletop application using the kappa architecture. Coordinate your microservices to
process and deliver a large scale data driven application.`,
		Version: pkg.Version,
	}
)

func Execute(ctx context.Context) {
	apps.Add(config.NewApp())
	apps.Add(marketdata.NewApp())
	apps.Add(orderbook.NewApp())
	apps.Add(process.NewApp())
	apps.Add(deliver.NewApp())

	cobra.CheckErr(rootCmd.ExecuteContext(ctx))
}
