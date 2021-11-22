package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/fremantle-industries/tabletop/cmd/start"
	"github.com/fremantle-industries/workshop/cmd/apps"
	"github.com/fremantle-industries/workshop/cmd/lake"
	"github.com/fremantle-industries/workshop/cmd/marketdata"
)

var (
	appsCmd = &cobra.Command{
		Use:   "apps",
		Short: "Commands to manage OTP applications",
	}

	brokersCmd = &cobra.Command{
		Use:   "brokers",
		Short: "List cluster brokers",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("TODO> brokers\n")

			return nil
		},
	}

	lakesCmd = &cobra.Command{
		Use:   "lakes",
		Short: "Commands to manage data lakes available to the cluster",
	}

	marketDataCmd = &cobra.Command{
		Use:   "marketdata",
		Short: "Commands to manage market data",
	}

	newCmd = &cobra.Command{
		Use:   "new",
		Short: "Create a new workshop project",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("TODO> new\n")

			return nil
		},
	}
)

func AddCommands() {
	apps.AddCommands(appsCmd)
	lake.AddCommands(lakesCmd)
	marketdata.AddCommands(marketDataCmd)

	rootCmd.AddCommand(
		start.StartCmd,
		appsCmd,
		brokersCmd,
		lakesCmd,
		marketDataCmd,
		newCmd,
	)
}
