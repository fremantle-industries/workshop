package cmd

import (
	"github.com/fremantle-industries/workshop/cmd/compose"
	"github.com/spf13/cobra"
)

var composeCmd = &cobra.Command{
	Use:   "compose",
	Short: "Commands that manage local development through docker compose",
}

func init() {
	composeCmd.AddCommand(compose.GenerateCmd)
}
