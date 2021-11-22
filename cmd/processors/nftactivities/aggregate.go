package nftactivities

import (
	"fmt"

	"github.com/spf13/cobra"
)

var AggregateCmd = &cobra.Command{
	Use:   "aggregate",
	Short: "Aggregate normalized nft activity events from all venues",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("TODO> processors/nftactivities/aggregate")
		return nil
	},
}
