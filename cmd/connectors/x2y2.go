package connectors

import (
	"log"
	"sync"

	"github.com/fremantle-industries/workshop/pkg/connectors/x2y2"
	"github.com/spf13/cobra"
)

var X2Y2Cmd = &cobra.Command{
	Use:   "x2y2",
	Short: "Commands that manage X2Y2 inbound & outbound connectors",
}

var x2y2StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start X2Y2 inbound & outbound connector",
	Long: `Run X2Y2 connectors configured for the following streams:

- [ ] poll events
- [ ] poll listings
- [ ] poll offers
- [ ] poll collections
- [ ] poll collection stats
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s := x2y2.NewProducer([]string{})

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
	X2Y2Cmd.AddCommand(x2y2StartCmd)
}
