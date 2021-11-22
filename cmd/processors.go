package cmd

import (
	"github.com/fremantle-industries/workshop/cmd/processors"
	"github.com/spf13/cobra"
)

var processorsCmd = &cobra.Command{
	Use:   "processors",
	Short: "Commands that manage processors from inbound and intermediate topics",
}

func init() {
	processorsCmd.AddCommand(processors.NftActivitiesCmd)
}
