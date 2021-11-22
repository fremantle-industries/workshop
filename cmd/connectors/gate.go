package connectors

import (
	"log"
	"sync"

	"github.com/fremantle-industries/workshop/pkg/connectors/gate"
	"github.com/spf13/cobra"
)

var GateCmd = &cobra.Command{
	Use:   "gate",
	Short: "Commands that manage Gate inbound & outbound connectors",
}

var gateStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Gate inbound & outbound connector",
	Long: `Run Gate connectors configured for the following streams:

- [ ] websocket order book
- [ ] websocket trades
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s := gate.NewProducer([]string{})

		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Fatal(s.Start())
		}()

		wg.Wait()
		return nil
	},
}

func init() {
	GateCmd.AddCommand(gateStartCmd)
}
