package ws

import (
	"github.com/aff-vending-machine/vm-controller/internal/boot/registry"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/api/ksher"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/api/link2500"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/api/topic"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/asset/fonts"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/asset/images"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/hardware/queue"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository/machine"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository/payment_channel"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository/slot"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository/transaction"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/websocket/frontend"
)

func NewService(module Module) registry.Service {
	return registry.Service{
		API: registry.APIService{
			Ksher:    ksher.New(module.HTTP.Client),
			Link2500: link2500.New(module.HTTP.Client),
			Topic:    topic.New(module.RabbitMQ.Connection),
		},
		Asset: registry.AssetService{
			Fonts:  fonts.New(module.Config.App.Asset),
			Images: images.New(module.Config.App.Asset),
		},
		Hardware: registry.HardwareService{
			Queue: queue.New(module.Redis.Client),
		},
		Repository: registry.RepositoryService{
			Machine:        machine.New(module.SQLite.DB),
			PaymentChannel: payment_channel.New(module.SQLite.DB),
			Slot:           slot.New(module.SQLite.DB),
			Transaction:    transaction.New(module.SQLite.DB),
		},
		WebSocket: registry.WebSocketService{
			Frontend: frontend.New(),
		},
	}
}
