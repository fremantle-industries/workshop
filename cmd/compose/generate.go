package compose

import (
	"fmt"

	"github.com/spf13/cobra"
)

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Create a docker compose file for local development",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("TODO> compose generate")
		return nil
	},
}
