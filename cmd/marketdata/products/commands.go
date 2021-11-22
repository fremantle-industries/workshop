package products

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all products currently in the data lake",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("TODO> marketdata products list")

			return nil
		},
	}

	pullCmd = &cobra.Command{
		Use:   "pull",
		Short: "Pull the latest products into the data lake for the given sources",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("TODO> marketdata products pull")

			return nil
		},
	}

	removeCmd = &cobra.Command{
		Use:   "remove",
		Short: "Remove the given products from the data lake for the given source",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("TODO> marketdata products remove")

			return nil
		},
	}

	clearCmd = &cobra.Command{
		Use:   "clear",
		Short: "Remove all products from the data lake for the given source",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("TODO> marketdata products remove")

			return nil
		},
	}
)

func AddCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		listCmd,
		pullCmd,
		removeCmd,
		clearCmd,
	)
}
