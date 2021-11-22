package cmd

import (
	"fmt"

	"github.com/fremantle-industries/tabletop/cmd"
	"github.com/spf13/cobra"

	"github.com/fremantle-industries/workshop/cmd/marketdata"
)

var (
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
	cmd.AddDataCommands(rootCmd)
	cmd.AddAppCommands(rootCmd)

	marketdata.AddCommands(marketDataCmd)

	rootCmd.AddCommand(
		marketDataCmd,
		newCmd,
	)
}
