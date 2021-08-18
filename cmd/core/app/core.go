package app

import (
	"github.com/urfave/cli/v2"
	"k8s.io/sample-controller/pkg/signals"

	"github.com/Fish-pro/multi-cluster-proxy/pkg/apiserver"
	"github.com/Fish-pro/multi-cluster-proxy/pkg/config"
	"github.com/Fish-pro/multi-cluster-proxy/pkg/core"
	"github.com/Fish-pro/multi-cluster-proxy/pkg/mlog"
)

var Start = func(c *cli.Context) error {
	conf := config.NewConfigFromEnv()

	run(conf, signals.SetupSignalHandler())

	return nil
}

func run(c *config.Config, stop <-chan struct{}) {

	// init logger
	mlog.NewLoggerWithOptions(c.LogLevel)

	core := core.New(c)

	core.IntegrateWith("multi-cluster-proxy", apiserver.NewAPIServerWithOpts(c))

	err := core.Initialize()
	if err != nil {
		mlog.Fatal("cube initialized failed: %v", err)
	}

	core.Run(stop)
}
