package registry

import (
	"vm-controller/configs"
	"vm-controller/internal/boot/modules"
	"vm-controller/internal/core/infra/network/fiber"
	"vm-controller/internal/core/infra/network/http"
	"vm-controller/internal/core/infra/network/rabbitmq"
	"vm-controller/internal/core/infra/network/redis"
	"vm-controller/internal/core/infra/network/websocket"
	"vm-controller/internal/core/infra/storage/sqlite"
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
