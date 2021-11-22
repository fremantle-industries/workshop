package connectors

import (
	"log"
	"sync"

	"github.com/fremantle-industries/workshop/pkg/connectors/bit"
	"github.com/spf13/cobra"
)

var BitCmd = &cobra.Command{
	Use:   "bit",
	Short: "Commands that manage Bit inbound & outbound connectors",
}

var bitStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Bit inbound & outbound connector",
	Long: `Run Bit connectors configured for the following streams:

- [ ] websocket order book
- [ ] websocket trades
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s := bit.NewProducer([]string{})

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
	BitCmd.AddCommand(bitStartCmd)
}
