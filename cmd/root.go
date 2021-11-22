package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var VERSION = "dev"

var (
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use: "workshop",
		Long: `CLI interface to manage the workshop platform.

It allows users to execute actions on insights extracted from data that has been streamed,
recorded & analyzed via a single binary categorized by subcommands.
    `,
		Version: VERSION,
	}
)

func Execute(ctx context.Context) {
	cobra.CheckErr(rootCmd.ExecuteContext(ctx))
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.workshop.yaml)")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))

	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(topicsCmd)
	rootCmd.AddCommand(connectorsCmd)
	rootCmd.AddCommand(processorsCmd)
	rootCmd.AddCommand(servingCmd)
	rootCmd.AddCommand(controlCmd)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".workshop" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".workshop")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
