package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var VERSION = "dev"

var rootCmd = &cobra.Command{
	Use:     "workshop",
	Version: VERSION,
}

func Execute(ctx context.Context) {
	cobra.CheckErr(rootCmd.ExecuteContext(ctx))
}
