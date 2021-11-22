package connectors

import (
	"log"
	"sync"

	"github.com/fremantle-industries/workshop/pkg/connectors/bitmex"
	"github.com/spf13/cobra"
)

var BitmexCmd = &cobra.Command{
	Use:   "bitmex",
	Short: "Commands that manage Bitmex inbound & outbound connectors",
}

var bitmexStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Bitmex inbound & outbound connector",
	Long: `Run Bitmex connectors configured for the following streams:

- [ ] websocket order book
- [ ] websocket trades
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s := bitmex.NewProducer([]string{})

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
	BitmexCmd.AddCommand(bitmexStartCmd)
}
