package cmd

import (
	"github.com/fremantle-industries/workshop/cmd/topics"
	"github.com/spf13/cobra"
)

var topicsCmd = &cobra.Command{
	Use:   "topics",
	Short: "Commands that manage topics required to operate workshop",
}

func init() {
	topicsCmd.AddCommand(topics.ListCmd)
}
