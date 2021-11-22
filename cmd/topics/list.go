package topics

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Display topics under the 'workshop.*' namespace",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("TODO> topics list")
		return nil
	},
}
