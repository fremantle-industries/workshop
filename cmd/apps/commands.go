package apps

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all applications",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("TODO> apps list")

			return nil
		},
	}

	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start applications",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("TODO> apps start")

			return nil
		},
	}

	stopCmd = &cobra.Command{
		Use:   "stop",
		Short: "Stop applications",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("TODO> apps stop")

			return nil
		},
	}
)

func AddCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		listCmd,
		startCmd,
		stopCmd,
	)
}
