package app

import (
	"github.com/aff-vending-machine/vm-controller/configs"
	"github.com/aff-vending-machine/vm-controller/internal/core/infra/network/fiber"
	"github.com/aff-vending-machine/vm-controller/internal/core/infra/network/http"
	"github.com/aff-vending-machine/vm-controller/internal/core/infra/network/rabbitmq"
	"github.com/aff-vending-machine/vm-controller/internal/core/infra/network/redis"
	"github.com/aff-vending-machine/vm-controller/internal/core/infra/network/websocket"
	"github.com/aff-vending-machine/vm-controller/internal/core/infra/storage/sqlite"
)

// Infrastructure
type Module struct {
	Config    configs.Config
	Fiber     *fiber.Wrapper
	HTTP      *http.Wrapper
	RabbitMQ  *rabbitmq.Wrapper
	Redis     *redis.Wrapper
	SQLite    *sqlite.Client
	WebSocket *websocket.Wrapper
}

func NewInfrastructure(cfg configs.Config) Module {
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
