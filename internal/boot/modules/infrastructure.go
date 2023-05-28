package modules

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
type Infrastructure struct {
	Config    configs.Config
	Fiber     *fiber.Wrapper
	HTTP      *http.Wrapper
	RabbitMQ  *rabbitmq.Wrapper
	Redis     *redis.Wrapper
	SQLite    *sqlite.Client
	WebSocket *websocket.Wrapper
}
