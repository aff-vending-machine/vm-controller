package modules

import (
	"vm-controller/configs"
	"vm-controller/internal/core/infra/network/fiber"
	"vm-controller/internal/core/infra/network/http"
	"vm-controller/internal/core/infra/network/rabbitmq"
	"vm-controller/internal/core/infra/network/redis"
	"vm-controller/internal/core/infra/network/websocket"
	"vm-controller/internal/core/infra/storage/sqlite"
)

// Infrastructure
type Infrastructure struct {
	Config    configs.Config
	Fiber     *fiber.Wrapper
	HTTP      *http.Wrapper
	RabbitMQ  *rabbitmq.Wrapper
	Redis     *redis.Wrapper
	SQLite    *sqlite.Client
	WebSocket *websocket.Wrapper
}
