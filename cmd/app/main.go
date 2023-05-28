package main

import (
	"github.com/aff-vending-machine/vm-controller/configs"
	"github.com/aff-vending-machine/vm-controller/internal/boot/registry/app"
	"github.com/aff-vending-machine/vm-controller/pkg/boot"
	"github.com/aff-vending-machine/vm-controller/pkg/log"
)

func init() {
	log.New()
}

func main() {
	// Create boot with configuration
	cfg := configs.Init("env/jetts")
	boot.Init(cfg)
	defer boot.Serve()

	initLog(cfg)

	// Run main application
	app.Run(cfg)
}

func initLog(cfg configs.Config) {
	log.SetOutput(log.ColorConsole())
	log.SetLogLevel(cfg.App.LogLevel)
	cfg.Preview()
}
