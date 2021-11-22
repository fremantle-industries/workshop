package generate

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ComposeCmd = &cobra.Command{
	Use:   "compose",
	Short: "Create a docker compose file for local development",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("TODO> generate compose")
		return nil
	},
}
