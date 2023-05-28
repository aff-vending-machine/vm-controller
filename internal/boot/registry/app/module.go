package app

import (
	"github.com/aff-vending-machine/vm-controller/config"
	"github.com/aff-vending-machine/vm-controller/internal/core/infra/fiber"
	"github.com/aff-vending-machine/vm-controller/internal/core/infra/http"
	"github.com/aff-vending-machine/vm-controller/internal/core/infra/rabbitmq"
	"github.com/aff-vending-machine/vm-controller/internal/core/infra/redis"
	"github.com/aff-vending-machine/vm-controller/internal/core/infra/sqlite"
	"github.com/aff-vending-machine/vm-controller/internal/core/infra/websocket"
)

// Infrastructure
type Module struct {
	Config    config.BootConfig
	Fiber     *fiber.Wrapper
	HTTP      *http.Wrapper
	RabbitMQ  *rabbitmq.Wrapper
	Redis     *redis.Wrapper
	SQLite    *sqlite.Wrapper
	WebSocket *websocket.Wrapper
}

func NewModule(cfg config.BootConfig) Module {
	return Module{
		Config:    cfg,
		Fiber:     fiber.New(cfg.Fiber),
		HTTP:      http.New(cfg.HTTP),
		RabbitMQ:  rabbitmq.New(cfg.RabbitMQ),
		Redis:     redis.New(cfg.Redis),
		SQLite:    sqlite.New(cfg.SQLite),
		WebSocket: websocket.New(cfg.WebSocket),
	}
}
