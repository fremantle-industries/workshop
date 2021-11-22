package serving

import (
	"fmt"

	"github.com/spf13/cobra"
)

var OrderbookCmd = &cobra.Command{
	Use:   "orderbook",
	Short: "Run the orderbook GraphQL API for the serving layer",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("TODO> serving orderbook")
		return nil
	},
}
