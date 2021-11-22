package connectors

import (
	"log"
	"sync"

	"github.com/fremantle-industries/workshop/pkg/connectors/kraken"
	"github.com/spf13/cobra"
)

var KrakenCmd = &cobra.Command{
	Use:   "kraken",
	Short: "Commands that manage Kraken inbound & outbound connectors",
}

var krakenStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Kraken inbound & outbound connector",
	Long: `Run Kraken connectors configured for the following streams:

- [ ] websocket order book
- [ ] websocket trades
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s := kraken.NewProducer([]string{})

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
	KrakenCmd.AddCommand(krakenStartCmd)
}
