package console

import (
	"github.com/aff-vending-machine/vm-controller/config"
	"github.com/aff-vending-machine/vm-controller/internal/core/module/console"
	"github.com/aff-vending-machine/vm-controller/internal/core/module/fiber"
	"github.com/aff-vending-machine/vm-controller/internal/core/module/rabbitmq"
	"github.com/aff-vending-machine/vm-controller/internal/core/module/sqlite"
)

// Infrastructure
type Module struct {
	Config   config.BootConfig
	Console  *console.Wrapper
	Fiber    *fiber.Wrapper
	RabbitMQ *rabbitmq.Wrapper
	SQLite   *sqlite.Wrapper
}

func NewModule(cfg config.BootConfig) Module {
	return Module{
		Config:   cfg,
		Console:  console.New(),
		Fiber:    fiber.New(cfg.Fiber),
		RabbitMQ: rabbitmq.New(cfg.RabbitMQ),
		SQLite:   sqlite.New(cfg.SQLite),
	}
}
