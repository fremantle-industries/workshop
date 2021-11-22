package subscriptions

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List the current subscriptions in the cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("TODO> marketdata subscriptions list")

			return nil
		},
	}

	addCmd = &cobra.Command{
		Use:   "add",
		Short: "Subscribe the cluster to the given market data",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("TODO> marketdata subscriptions add")

			return nil
		},
	}

	removeCmd = &cobra.Command{
		Use:   "remove",
		Short: "Unsubscribe the cluster from the given market data",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("TODO> marketdata subscriptions remove")

			return nil
		},
	}
)

func AddCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		listCmd,
		addCmd,
		removeCmd,
	)
}
