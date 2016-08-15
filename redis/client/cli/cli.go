package cli

import (
	"fmt"
	"os"
	"sync"

	"github.com/codegangsta/cli"
	httpclient "github.com/xozrc/benchmark/redis/client/client/http"
)

var (
	clientType  int    = 0
	host        string = "127.0.0.1"
	port        int    = 3000
	number      int    = 1
	concurrency int    = 1
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
			Value:       clientType,
			Usage:       "client type",
			Destination: &clientType,
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
		cli.IntFlag{
			Name:        "number,n",
			Value:       number,
			Usage:       "number",
			Destination: &number,
		},
		cli.IntFlag{
			Name:        "concurrency,c",
			Value:       concurrency,
			Usage:       "concurrency",
			Destination: &concurrency,
		},
	}

	app.Action = action
	app.Run(os.Args)
}

var (
	wg sync.WaitGroup
)

func action(ctx *cli.Context) {
	addr := fmt.Sprintf("http://%s:%d/api/set", host, port)
	tasks := make(chan struct{}, concurrency)
	go func(n int) {
		for i := 0; i < n; i++ {
			tasks <- struct{}{}
		}
		close(tasks)
	}(number)

	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			c := httpclient.NewHttpClient(addr)
			for {
				select {
				case _, ok := <-tasks:
					{
						if !ok {

							goto breakL
						}
						err := c.Set("test", "123")
						if err != nil {
							fmt.Print(err.Error())
							os.Exit(1)
						}
					}
				}
			}
		breakL:
			wg.Done()
		}()
	}

	wg.Wait()
}
