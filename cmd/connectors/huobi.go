package connectors

import (
	"log"
	"sync"

	"github.com/fremantle-industries/workshop/pkg/connectors/huobi"
	"github.com/spf13/cobra"
)

var HuobiCmd = &cobra.Command{
	Use:   "huobi",
	Short: "Commands that manage Huobi inbound & outbound connectors",
}

var huobiStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Huobi inbound & outbound connector",
	Long: `Run Huobi connectors configured for the following streams:

- [ ] websocket order book
- [ ] websocket trades
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s := huobi.NewProducer([]string{})

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
	HuobiCmd.AddCommand(huobiStartCmd)
}
