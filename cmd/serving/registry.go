package serving

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RegistryCmd = &cobra.Command{
	Use:   "registry",
	Short: "Run the entity registry GraphQL API for the serving layer",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("TODO> serving registry")
		return nil
	},
}
