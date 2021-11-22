package marketdata

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"

	"github.com/fremantle-industries/workshop/apps/marketdata"
	"github.com/fremantle-industries/workshop/cmd/marketdata/products"
	"github.com/fremantle-industries/workshop/cmd/marketdata/subscriptions"
)

var (
	sourcesCmd = &cobra.Command{
		Use:   "sources",
		Short: "List all market data sources available to the node",
		RunE: func(cmd *cobra.Command, args []string) error {
			sources := marketdata.Sources()

			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			t.AppendHeader(table.Row{"Name", "Status"})
			if len(sources) == 0 {
				t.AppendRow(table.Row{"No sources available"})
			}
			for k := range sources {
				t.AppendRow(table.Row{k, "-"})
			}
			t.SortBy([]table.SortBy{
				{
					Name: "Name",
					Mode: table.Asc,
				},
			})
			t.Render()

			return nil
		},
	}

	productsCmd = &cobra.Command{
		Use:   "products",
		Short: "Commands to manage product market data",
	}

	subscriptionsCmd = &cobra.Command{
		Use:   "subscriptions",
		Short: "Commands to manage subscription market data",
	}
)

func AddCommands(cmd *cobra.Command) {
	products.AddCommands(productsCmd)
	subscriptions.AddCommands(subscriptionsCmd)

	cmd.AddCommand(
		sourcesCmd,
		productsCmd,
		subscriptionsCmd,
	)
}
