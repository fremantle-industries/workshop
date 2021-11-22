package connectors

import (
	"log"
	"sync"

	"github.com/fremantle-industries/workshop/pkg/connectors/kucoin"
	"github.com/spf13/cobra"
)

var KucoinCmd = &cobra.Command{
	Use:   "kucoin",
	Short: "Commands that manage Kucoin inbound & outbound connectors",
}

var kucoinStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Kucoin inbound & outbound connector",
	Long: `Run Kucoin connectors configured for the following streams:

- [ ] websocket order book
- [ ] websocket trades
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s := kucoin.NewProducer([]string{})

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
	KucoinCmd.AddCommand(kucoinStartCmd)
}
