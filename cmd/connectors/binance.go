package connectors

import (
	"log"
	"sync"

	"github.com/fremantle-industries/workshop/pkg/connectors/binance"
	"github.com/spf13/cobra"
)

var BinanceCmd = &cobra.Command{
	Use:   "binance",
	Short: "Commands that manage Binance inbound & outbound connectors",
}

var binanceStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Binance inbound & outbound connector",
	Long: `Run Binance connectors configured for the following streams:

- [ ] websocket order book
- [ ] websocket trades
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s := binance.NewProducer([]string{})

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
	BinanceCmd.AddCommand(binanceStartCmd)
}
