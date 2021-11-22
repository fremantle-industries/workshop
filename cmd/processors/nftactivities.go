package processors

import (
	"github.com/fremantle-industries/workshop/cmd/processors/nftactivities"
	"github.com/spf13/cobra"
)

var NftActivitiesCmd = &cobra.Command{
	Use:   "nftactivities",
	Short: "Commands that manage nft activity processors",
}

func init() {
	NftActivitiesCmd.AddCommand(nftactivities.AggregateCmd)
}
