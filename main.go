package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/LotteWong/giotto-gateway-admin/router"
	"github.com/e421083458/golang_common/lib"
)

var (
	config = flag.String("config", "./configs/dev/", "config file path")
)

func main() {
	flag.Parse()
	if *config == "" {
		flag.Usage()
		os.Exit(1)
	}

	InitManagementServer(*config)
}

func InitManagementServer(config string) {
	lib.InitModule(config, []string{"base", "mysql", "redis"})
	defer lib.Destroy()

	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
