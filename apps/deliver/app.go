package deliver

import (
	"log"

	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
)

var (
	appName = "deliver"
)

func NewApp() gen.ApplicationBehavior {
	return &app{}
}

type app struct {
	gen.Application
}

func (a *app) Load(args ...etf.Term) (gen.ApplicationSpec, error) {
	return gen.ApplicationSpec{
		Name:        appName,
		Description: "Deliver application",
		Version:     "v0.0.1",
		Children:    []gen.ApplicationChildSpec{},
	}, nil
}

func (a *app) Start(process gen.Process, args ...etf.Term) {
	log.Printf(`Application started with pid=%s, app_name="%s", node_name="%s"`,
		process.Self(),
		appName,
		process.NodeName(),
	)
}
