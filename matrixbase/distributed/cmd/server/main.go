package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/matrixorigin/talent-challenge/matrixbase/distributed/pkg/api"
	"github.com/matrixorigin/talent-challenge/matrixbase/distributed/pkg/cfg"
)

var (
	file = flag.String("cfg", "/etc/cfg.toml", "The configuration file")
)

var (
	stopping = false
)

func main() {
	flag.Parse()

	c := cfg.MustParseCfg(*file)
	svr, err := api.NewServer(c)
	if err != nil {
		log.Fatalf("create api server failed with %+v", err)
	}

	go func() {
		log.Fatalf("api server start failed with %+v", svr.Start())
	}()

	sc := make(chan os.Signal, 2)
	signal.Notify(sc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	for {
		sig := <-sc

		if !stopping {
			stopping = true
			go func() {
				svr.Stop()
				log.Printf("exit: signal=<%d>.", sig)

				switch sig {
				case syscall.SIGTERM:
					log.Printf("exit: bye :-).")
					os.Exit(0)
				default:
					log.Printf("exit: bye :-(.")
					os.Exit(1)
				}
			}()
			continue
		}

		log.Printf("exit: bye :-).")
		os.Exit(0)
	}
}
