package cli

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"

	httpserver "github.com/xozrc/benchmark/redis/server/http"
)

var (
	serverType int    = 0
	host       string = "127.0.0.1"
	port       int    = 3000
)

func Start() {
	app := cli.NewApp()
	app.Name = "benchmark redis server"
	app.Usage = "benchmark redis server"

	app.Author = ""
	app.Email = ""

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "type,t",
			Value:       serverType,
			Usage:       "server type",
			Destination: &serverType,
		},
		cli.StringFlag{
			Name:        "host",
			Value:       host,
			Usage:       "host for listen",
			Destination: &host,
		},
		cli.IntFlag{
			Name:        "port,p",
			Value:       port,
			Usage:       "port for listen",
			Destination: &port,
		},
	}

	app.Action = action
	app.Run(os.Args)
}

func action(ctx *cli.Context) {
	addr := fmt.Sprintf("%s:%d", host, port)
	httpserver.Listen(addr)
}
