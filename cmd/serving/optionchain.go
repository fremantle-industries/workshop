package serving

import (
	"fmt"

	"github.com/spf13/cobra"
)

var OptionchainCmd = &cobra.Command{
	Use:   "optionchain",
	Short: "Run the option chain GraphQL API for the serving layer",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("TODO> serving optionchain")
		return nil
	},
}
