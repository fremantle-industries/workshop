package coinbase

import (
	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
)

func NewSupervisor() gen.SupervisorBehavior {
	return &supervisor{}
}

type supervisor struct {
	gen.Supervisor
}

func (s *supervisor) Init(args ...etf.Term) (gen.SupervisorSpec, error) {
	url := "wss://ws-feed.exchange.coinbase.com"
	seedBrokers := []string{"redpanda-broker-1:27095", "redpanda-broker-2:27095", "redpanda-broker-3:27095"}

	spec := gen.SupervisorSpec{
		Name: "coinbaseSupervisor",
		Children: []gen.SupervisorChildSpec{
			{
				Name:  "websocketProducer",
				Child: NewWebsocketProducer(url, seedBrokers),
			},
		},
		Strategy: gen.SupervisorStrategy{
			Type:      gen.SupervisorStrategyOneForAll,
			Intensity: 2,
			Period:    5,
			Restart:   gen.SupervisorStrategyRestartTemporary,
		},
	}
	return spec, nil
}
