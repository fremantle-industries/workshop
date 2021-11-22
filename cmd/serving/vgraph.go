package serving

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VgraphCmd = &cobra.Command{
	Use:   "vgraph",
	Short: "Run the federated GraphQL API for the serving layer",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("TODO> serving vgraph")
		return nil
	},
}
