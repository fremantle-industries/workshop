package cmd

import (
	"github.com/fremantle-industries/workshop/cmd/generate"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "TODO... generate description",
}

func init() {
	generateCmd.AddCommand(generate.ComposeCmd)
}
