package main

import (
	"github.com/aff-vending-machine/vm-controller/config"
	"github.com/aff-vending-machine/vm-controller/internal/app"
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
	// initTrace(conf)

	// Run main application
	app.Run(conf)
}

func initLog(conf config.BootConfig) {
	if conf.App.ENV == "local" {
		log.SetOutput(log.ColorConsole())
	}
	log.SetLogLevel(conf.App.LogLevel)
}

// func initTrace(conf config.BootConfig) {
// 	endpoint := "http://localhost:14268/api/traces"
// 	provider, err := trace.Jaeger(endpoint, "raspi-ctrl", conf.App.ENV)
// 	boot.TerminateWhenError(err)

// 	trace.SetProvider(provider)
// }
