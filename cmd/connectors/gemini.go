package connectors

import (
	"log"
	"sync"

	"github.com/fremantle-industries/workshop/pkg/connectors/gemini"
	"github.com/spf13/cobra"
)

var GeminiCmd = &cobra.Command{
	Use:   "gemini",
	Short: "Commands that manage Gemini inbound & outbound connectors",
}

var geminiStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Gemini inbound & outbound connector",
	Long: `Run Gemini connectors configured for the following streams:

- [ ] websocket order book
- [ ] websocket trades
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s := gemini.NewProducer([]string{})

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
	GeminiCmd.AddCommand(geminiStartCmd)
}
