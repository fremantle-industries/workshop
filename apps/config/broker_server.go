package config

import (
	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
)

func NewBrokerServer() gen.ServerBehavior {
	return &brokerServer{}
}

type brokerServer struct {
	gen.Server
}

func (s *brokerServer) Init(process *gen.ServerProcess, args ...etf.Term) error {
	return nil
}
