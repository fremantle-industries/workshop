package connectors

import (
	"github.com/fremantle-industries/workshop/pkg/connectors/opensea"
	"github.com/spf13/cobra"
)

var OpenSeaCmd = &cobra.Command{
	Use:   "opensea",
	Short: "Commands that manage OpenSea inbound & outbound connectors",
}

var openSeaStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start OpenSea inbound & outbound connector",
	Long: `Run OpenSea connectors configured for the following streams:

- [.] websocket events
- [.] poll events
- [.] poll listings
- [.] poll offers
- [.] poll collections
- [.] poll collection stats
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: read broker seed config from viper
		seeds := []string{"redpanda-broker-1:29092", "redpanda-broker-2:29092", "redpanda-broker-3:29092"}
		// TODO: read OpenSea API key from viper
		apiKey := "83225c250c4f43ed85ec070f525a6d69"
		s, err := opensea.NewProducer(
			seeds,
			apiKey,
			// TODO: read stream enabled options from viper
			&opensea.WebsocketEventsOpt{Enabled: true, Collections: []string{"*"}},
			&opensea.PollEventsOpt{Enabled: false},
			&opensea.PollListingsOpt{Enabled: false},
			&opensea.PollOffersOpt{Enabled: false},
			&opensea.PollCollectionsOpt{Enabled: false},
			&opensea.PollCollectionStatsOpt{Enabled: false},
		)
		if err != nil {
			return err
		}

		return s.Start()
	},
}

func init() {
	OpenSeaCmd.AddCommand(openSeaStartCmd)
}
