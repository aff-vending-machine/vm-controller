package registry

import (
	"github.com/aff-vending-machine/vm-controller/internal/boot/modules"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/api/ksher"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/api/link2500"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/api/topic"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/hardware/queue"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository/machine"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository/payment_channel"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository/slot"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository/transaction"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/websocket/frontend"
)

func NewService(infra modules.Infrastructure) modules.Service {
	return modules.Service{
		API: modules.APIService{
			Ksher:    ksher.New(infra.HTTP.Client),
			Link2500: link2500.New(infra.HTTP.Client),
			Topic:    topic.New(infra.RabbitMQ.Connection),
		},
		Hardware: modules.HardwareService{
			Queue: queue.New(infra.Redis.Client),
		},
		Repository: modules.RepositoryService{
			Machine:        machine.New(infra.SQLite.DB),
			PaymentChannel: payment_channel.New(infra.SQLite.DB),
			Slot:           slot.New(infra.SQLite.DB),
			Transaction:    transaction.New(infra.SQLite.DB),
		},
		WebSocket: modules.WebSocketService{
			Frontend: frontend.New(),
		},
	}
}
