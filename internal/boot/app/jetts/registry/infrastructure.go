package registry

import (
	"github.com/aff-vending-machine/vm-controller/configs"
	"github.com/aff-vending-machine/vm-controller/internal/boot/modules"
	"github.com/aff-vending-machine/vm-controller/internal/core/infra/network/fiber"
	"github.com/aff-vending-machine/vm-controller/internal/core/infra/network/http"
	"github.com/aff-vending-machine/vm-controller/internal/core/infra/network/rabbitmq"
	"github.com/aff-vending-machine/vm-controller/internal/core/infra/network/redis"
	"github.com/aff-vending-machine/vm-controller/internal/core/infra/network/websocket"
	"github.com/aff-vending-machine/vm-controller/internal/core/infra/storage/sqlite"
)

func NewInfrastructure(cfg configs.Config) modules.Infrastructure {
	return modules.Infrastructure{
		Config:    cfg,
		Fiber:     fiber.New(cfg.Fiber),
		HTTP:      http.New(cfg.HTTP),
		RabbitMQ:  rabbitmq.New(cfg.RabbitMQ),
		Redis:     redis.New(cfg.Redis),
		SQLite:    sqlite.New(cfg.SQLite),
		WebSocket: websocket.New(cfg.WebSocket),
	}
}
