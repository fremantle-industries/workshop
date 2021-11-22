package marketdata

import (
	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"

	"github.com/fremantle-industries/workshop/apps/marketdata/source/alpaca"
	"github.com/fremantle-industries/workshop/apps/marketdata/source/binance"
	"github.com/fremantle-industries/workshop/apps/marketdata/source/bitfinex"
	"github.com/fremantle-industries/workshop/apps/marketdata/source/bitmex"
	"github.com/fremantle-industries/workshop/apps/marketdata/source/bybit"
	"github.com/fremantle-industries/workshop/apps/marketdata/source/coinbase"
	"github.com/fremantle-industries/workshop/apps/marketdata/source/deltaexchange"
	"github.com/fremantle-industries/workshop/apps/marketdata/source/deribit"
	"github.com/fremantle-industries/workshop/apps/marketdata/source/gateio"
	"github.com/fremantle-industries/workshop/apps/marketdata/source/kraken"
	"github.com/fremantle-industries/workshop/apps/marketdata/source/kucoin"
	"github.com/fremantle-industries/workshop/apps/marketdata/source/ledgerx"
	"github.com/fremantle-industries/workshop/apps/marketdata/source/okx"
	"github.com/fremantle-industries/workshop/apps/marketdata/source/polygonio"
)

var (
	sources = map[string]gen.SupervisorChildSpec{
		"alpaca": {
			Child: alpaca.NewSupervisor(),
			Name:  "alpacaSupervisor",
		},
		"binance": {
			Child: binance.NewSupervisor(),
			Name:  "binanceSupervisor",
		},
		"bitfinex": {
			Child: bitfinex.NewSupervisor(),
			Name:  "bitfinexSupervisor",
		},
		"bitmex": {
			Child: bitmex.NewSupervisor(),
			Name:  "bitmexSupervisor",
		},
		"bybit": {
			Child: bybit.NewSupervisor(),
			Name:  "bybitSupervisor",
		},
		"coinbase": {
			Child: coinbase.NewSupervisor(),
			Name:  "coinbaseSupervisor",
		},
		"deltaexchange": {
			Child: deltaexchange.NewSupervisor(),
			Name:  "deltaexchangeSupervisor",
		},
		"deribit": {
			Child: deribit.NewSupervisor(),
			Name:  "deribitSupervisor",
		},
		"gateio": {
			Child: gateio.NewSupervisor(),
			Name:  "gateioSupervisor",
		},
		"kraken": {
			Child: kraken.NewSupervisor(),
			Name:  "krakenSupervisor",
		},
		"kucoin": {
			Child: kucoin.NewSupervisor(),
			Name:  "kucoinSupervisor",
		},
		"ledgerx": {
			Child: ledgerx.NewSupervisor(),
			Name:  "ledgerxSupervisor",
		},
		"okx": {
			Child: okx.NewSupervisor(),
			Name:  "okxSupervisor",
		},
		"polygonio": {
			Child: polygonio.NewSupervisor(),
			Name:  "polygonioSupervisor",
		},
	}
)

func NewSourceSupervisor() gen.SupervisorBehavior {
	return &sourceSupervisor{}
}

type sourceSupervisor struct {
	gen.Supervisor
}

func (s *sourceSupervisor) Init(args ...etf.Term) (gen.SupervisorSpec, error) {
	spec := gen.SupervisorSpec{
		Name: "sourceSupervisor",
		// TODO: Dynamically add source supervisors in response to commands
		Children: []gen.SupervisorChildSpec{},
		Strategy: gen.SupervisorStrategy{
			Type:      gen.SupervisorStrategyOneForAll,
			Intensity: 2,
			Period:    5,
			Restart:   gen.SupervisorStrategyRestartTemporary,
		},
	}
	return spec, nil
}

func Sources() map[string]gen.SupervisorChildSpec {
	return sources
}
