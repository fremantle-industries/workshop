package connectors

import (
	"log"
	"sync"

	"github.com/fremantle-industries/workshop/pkg/connectors/bybit"
	"github.com/spf13/cobra"
)

var BybitCmd = &cobra.Command{
	Use:   "bybit",
	Short: "Commands that manage Bybit inbound & outbound connectors",
}

var bybitStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Bybit inbound & outbound connector",
	Long: `Run Bybit connectors configured for the following streams:

- [ ] websocket order book
- [ ] websocket trades
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s := bybit.NewProducer([]string{})

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
	BybitCmd.AddCommand(bybitStartCmd)
}
