package connectors

import (
	"log"
	"sync"

	"github.com/fremantle-industries/workshop/pkg/connectors/okx"
	"github.com/spf13/cobra"
)

var OkxCmd = &cobra.Command{
	Use:   "okx",
	Short: "Commands that manage Okx inbound & outbound connectors",
}

var okxStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Okx inbound & outbound connector",
	Long: `Run Okx connectors configured for the following streams:

- [ ] websocket order book
- [ ] websocket trades
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s := okx.NewProducer([]string{})

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
	OkxCmd.AddCommand(okxStartCmd)
}
