package polygonio

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
	spec := gen.SupervisorSpec{
		Name:     "polygonioSupervisor",
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
