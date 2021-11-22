package cmd

import (
	"log"
	"sync"

	"github.com/fremantle-industries/workshop/pkg/control"
	"github.com/spf13/cobra"
)

var controlCmd = &cobra.Command{
	Use:   "control",
	Short: "Commands that manage control plane UI & API",
	Long:  "Workshop control plane UI & API for system administration",
}

var controlStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Run the control plane server",
	RunE: func(cmd *cobra.Command, args []string) error {
		wg := sync.WaitGroup{}

		uiPort := "9090"
		ui := control.NewUIServer(uiPort)
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Fatal(ui.ListenAndServe())
		}()

		apiPort := "9091"
		api := control.NewAPIServer(apiPort)
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Fatal(api.ListenAndServe())
		}()

		wg.Wait()
		return nil
	},
}

func init() {
	controlCmd.AddCommand(controlStartCmd)
}
