package connectors

import (
	"log"
	"sync"

	"github.com/fremantle-industries/workshop/pkg/connectors/deribit"
	"github.com/spf13/cobra"
)

var DeribitCmd = &cobra.Command{
	Use:   "deribit",
	Short: "Commands that manage Deribit inbound & outbound connectors",
}

var deribitStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Deribit inbound & outbound connector",
	Long: `Run Deribit connectors configured for the following streams:

- [ ] websocket order book
- [ ] websocket trades
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s := deribit.NewProducer([]string{})

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
	DeribitCmd.AddCommand(deribitStartCmd)
}
