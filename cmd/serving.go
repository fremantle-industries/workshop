package cmd

import (
	"github.com/fremantle-industries/workshop/cmd/serving"
	"github.com/spf13/cobra"
)

var servingCmd = &cobra.Command{
	Use:   "serving",
	Short: "Commands that manage serving layer endpoints",
}

func init() {
	servingCmd.AddCommand(serving.VgraphCmd)
	servingCmd.AddCommand(serving.RegistryCmd)
	servingCmd.AddCommand(serving.RollupCmd)
	servingCmd.AddCommand(serving.SearchCmd)
	servingCmd.AddCommand(serving.NftactivityCmd)
	servingCmd.AddCommand(serving.OrderbookCmd)
	servingCmd.AddCommand(serving.OptionchainCmd)
}
