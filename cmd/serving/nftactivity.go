package serving

import (
	"fmt"

	"github.com/spf13/cobra"
)

var NftactivityCmd = &cobra.Command{
	Use:   "nftactivity",
	Short: "Run the nft activity stream GraphQL API for the serving layer",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("TODO> serving nft activity")

		return nil
	},
}
