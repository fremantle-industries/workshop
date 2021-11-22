package connectors

import (
	"log"
	"sync"

	"github.com/fremantle-industries/workshop/pkg/connectors/hyperspace"
	"github.com/spf13/cobra"
)

var HyperspaceCmd = &cobra.Command{
	Use:   "hyperspace",
	Short: "Commands that manage Hyperspace inbound & outbound connectors",
}

var hyperspaceStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Hyperspace inbound & outbound connector",
	Long: `Run Hyperspace connectors configured for the following streams:

- [ ] poll events
- [ ] poll listings
- [ ] poll offers
- [ ] poll collections
- [ ] poll collection stats
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s := hyperspace.NewProducer([]string{})

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
	HyperspaceCmd.AddCommand(hyperspaceStartCmd)
}
