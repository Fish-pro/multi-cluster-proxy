package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/urfave/cli/v2"

	core "github.com/Fish-pro/multi-cluster-proxy/cmd/core/app"
)

var version = "1.0.0"

func main() {
	app := cli.NewApp()
	app.Name = "multi-cluster-proxy"
	app.Usage = "KubCube is foundation of the world upon"
	app.Version = version
	app.Compiled = time.Now()
	app.Copyright = "(c) " + strconv.Itoa(time.Now().Year()) + " multi-cluster-proxy"

	app.Action = core.Start

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
