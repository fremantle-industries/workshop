package cmd

import (
	"fmt"

	// "github.com/fremantle-industries/workshop/pkg/web"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "TODO...",
	Long:  "TODO...",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("TODO> web...")

    // port := "8080"
    //
    // s := server.New(port)
    // s.ListenAndServe()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
