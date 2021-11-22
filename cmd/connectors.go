package cmd

import (
	"github.com/fremantle-industries/workshop/cmd/connectors"
	"github.com/spf13/cobra"
)

var connectorsCmd = &cobra.Command{
	Use:   "connectors",
	Short: "Commands that manage inbound & outbound connectors",
}

func init() {
	connectorsCmd.AddCommand(connectors.OpenSeaCmd)
	connectorsCmd.AddCommand(connectors.X2Y2Cmd)
	connectorsCmd.AddCommand(connectors.LooksRareCmd)
	connectorsCmd.AddCommand(connectors.MagicEdenCmd)
	connectorsCmd.AddCommand(connectors.HyperspaceCmd)
	connectorsCmd.AddCommand(connectors.BinanceCmd)
	connectorsCmd.AddCommand(connectors.CoinbaseCmd)
	connectorsCmd.AddCommand(connectors.KrakenCmd)
	connectorsCmd.AddCommand(connectors.BitfinexCmd)
	connectorsCmd.AddCommand(connectors.KucoinCmd)
	connectorsCmd.AddCommand(connectors.BitstampCmd)
	connectorsCmd.AddCommand(connectors.GeminiCmd)
	connectorsCmd.AddCommand(connectors.OkxCmd)
	connectorsCmd.AddCommand(connectors.HuobiCmd)
	connectorsCmd.AddCommand(connectors.BybitCmd)
	connectorsCmd.AddCommand(connectors.BitmexCmd)
	connectorsCmd.AddCommand(connectors.GateCmd)
	connectorsCmd.AddCommand(connectors.DeribitCmd)
	connectorsCmd.AddCommand(connectors.BitCmd)
	connectorsCmd.AddCommand(connectors.DeltaExchangeCmd)
}
