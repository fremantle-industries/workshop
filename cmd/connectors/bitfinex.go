package connectors

import (
	"log"
	"sync"

	"github.com/fremantle-industries/workshop/pkg/connectors/bitfinex"
	"github.com/spf13/cobra"
)

var BitfinexCmd = &cobra.Command{
	Use:   "bitfinex",
	Short: "Commands that manage Bitfinex inbound & outbound connectors",
}

var bitfinexStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Bitfinex inbound & outbound connector",
	Long: `Run Bitfinex connectors configured for the following streams:

- [ ] websocket order book
- [ ] websocket trades
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s := bitfinex.NewProducer([]string{})

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
	BitfinexCmd.AddCommand(bitfinexStartCmd)
}
