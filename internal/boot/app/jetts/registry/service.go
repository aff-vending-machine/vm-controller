package registry

import (
	"vm-controller/internal/boot/modules"
	"vm-controller/internal/layer/service/api/ksher"
	"vm-controller/internal/layer/service/api/link2500"
	"vm-controller/internal/layer/service/api/topic"
	"vm-controller/internal/layer/service/hardware/queue"
	"vm-controller/internal/layer/service/repository/machine"
	"vm-controller/internal/layer/service/repository/payment_channel"
	"vm-controller/internal/layer/service/repository/slot"
	"vm-controller/internal/layer/service/repository/transaction"
	"vm-controller/internal/layer/service/websocket/frontend"
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
