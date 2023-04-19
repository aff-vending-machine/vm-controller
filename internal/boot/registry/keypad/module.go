package keypad

import (
	"github.com/aff-vending-machine/vm-controller/config"
	"github.com/aff-vending-machine/vm-controller/internal/core/module/fiber"
	"github.com/aff-vending-machine/vm-controller/internal/core/module/http"
	"github.com/aff-vending-machine/vm-controller/internal/core/module/keypad"
	"github.com/aff-vending-machine/vm-controller/internal/core/module/rabbitmq"
	"github.com/aff-vending-machine/vm-controller/internal/core/module/redis"
	"github.com/aff-vending-machine/vm-controller/internal/core/module/sqlite"
	"github.com/aff-vending-machine/vm-controller/internal/core/module/websocket"
)

// Infrastructure
type Module struct {
	Config    config.BootConfig
	Keypad    *keypad.Wrapper
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
		Keypad:    keypad.New(cfg.Board),
		Fiber:     fiber.New(cfg.Fiber),
		HTTP:      http.New(cfg.HTTP),
		RabbitMQ:  rabbitmq.New(cfg.RabbitMQ),
		Redis:     redis.New(cfg.Redis),
		SQLite:    sqlite.New(cfg.SQLite),
		WebSocket: websocket.New(cfg.WebSocket),
	}
}
