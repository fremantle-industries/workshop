package serving

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RollupCmd = &cobra.Command{
	Use:   "rollup",
	Short: "Run the entity rollup GraphQL API for the serving layer",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("TODO> serving rollup")
		return nil
	},
}
