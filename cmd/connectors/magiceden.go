package connectors

import (
	"log"
	"sync"

	"github.com/fremantle-industries/workshop/pkg/connectors/magiceden"
	"github.com/spf13/cobra"
)

var MagicEdenCmd = &cobra.Command{
	Use:   "magiceden",
	Short: "Commands that manage MagicEden inbound & outbound connectors",
}

var magicEdenStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start MagicEden inbound & outbound connector",
	Long: `Run MagicEden connectors configured for the following streams:

- [ ] poll events
- [ ] poll listings
- [ ] poll offers
- [ ] poll collections
- [ ] poll collection stats
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s := magiceden.NewProducer([]string{})

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
	MagicEdenCmd.AddCommand(magicEdenStartCmd)
}
