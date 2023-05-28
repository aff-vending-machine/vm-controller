package main

import (
	"vm-controller/configs"
	"vm-controller/internal/boot/app/jetts"
	"vm-controller/pkg/boot"
	"vm-controller/pkg/log"
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
	jetts.Run(cfg)
}

func initLog(cfg configs.Config) {
	log.SetOutput(log.ColorConsole())
	log.SetLogLevel(cfg.App.LogLevel)
	cfg.Preview()
}
