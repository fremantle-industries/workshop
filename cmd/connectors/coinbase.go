package connectors

import (
	"log"
	"sync"

	"github.com/fremantle-industries/workshop/pkg/connectors/coinbase"
	"github.com/spf13/cobra"
)

var CoinbaseCmd = &cobra.Command{
	Use:   "coinbase",
	Short: "Commands that manage Coinbase inbound & outbound connectors",
}

var coinbaseStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Coinbase inbound & outbound connector",
	Long: `Run Coinbase connectors configured for the following streams:

- [ ] websocket order book
- [ ] websocket trades
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s := coinbase.NewProducer([]string{})

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
	CoinbaseCmd.AddCommand(coinbaseStartCmd)
}
