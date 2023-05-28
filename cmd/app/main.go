package main

import (
	"github.com/aff-vending-machine/vm-controller/config"
	"github.com/aff-vending-machine/vm-controller/internal/boot/registry/app"
	"github.com/aff-vending-machine/vm-controller/pkg/boot"
	"github.com/aff-vending-machine/vm-controller/pkg/log"
)

func init() {
	log.New()
}

func main() {
	// Create boot with configuration
	conf := config.Init("env/jetts")
	boot.Init(conf)
	defer boot.Serve()

	initLog(conf)

	// Run main application
	app.Run(conf)
}

func initLog(conf config.BootConfig) {
	log.SetOutput(log.ColorConsole())
	log.SetLogLevel(conf.App.LogLevel)
}
