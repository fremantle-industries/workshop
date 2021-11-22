package serving

import (
	"fmt"

	"github.com/spf13/cobra"
)

var SearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Run the search GraphQL API for the serving layer",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("TODO> serving search")
		return nil
	},
}
