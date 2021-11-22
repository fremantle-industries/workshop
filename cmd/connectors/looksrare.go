package connectors

import (
	"log"
	"sync"

	"github.com/fremantle-industries/workshop/pkg/connectors/looksrare"
	"github.com/spf13/cobra"
)

var LooksRareCmd = &cobra.Command{
	Use:   "looksrare",
	Short: "Commands that manage LooksRare inbound & outbound connectors",
}

var looksRareStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start LooksRare inbound & outbound connector",
	Long: `Run LooksRare connectors configured for the following streams:

- [ ] poll events
- [ ] poll listings
- [ ] poll offers
- [ ] poll collections
- [ ] poll collection stats
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s := looksrare.NewProducer([]string{})

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
	LooksRareCmd.AddCommand(looksRareStartCmd)
}
