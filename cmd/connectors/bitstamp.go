package connectors

import (
	"log"
	"sync"

	"github.com/fremantle-industries/workshop/pkg/connectors/bitstamp"
	"github.com/spf13/cobra"
)

var BitstampCmd = &cobra.Command{
	Use:   "bitstamp",
	Short: "Commands that manage Bitstamp inbound & outbound connectors",
}

var bitstampStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Bitstamp inbound & outbound connector",
	Long: `Run Bitstamp connectors configured for the following streams:

- [ ] websocket order book
- [ ] websocket trades
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s := bitstamp.NewProducer([]string{})

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
	BitstampCmd.AddCommand(bitstampStartCmd)
}
