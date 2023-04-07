package console

import (
	"github.com/aff-vending-machine/vm-controller/internal/boot/registry"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/api/topic"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/asset/images"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/console/fonts"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/console/ksher"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/console/lcd2k"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/console/link2500"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/console/queue"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository/machine"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository/payment_channel"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository/slot"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository/transaction"
)

func NewService(module Module) registry.Service {
	return registry.Service{
		API: registry.APIService{
			Ksher:    ksher.New(),
			Link2500: link2500.New(),
			Topic:    topic.New(module.RabbitMQ.Connection),
		},
		Asset: registry.AssetService{
			Fonts:  fonts.New(),
			Images: images.New(module.Config.App.Asset),
		},
		Display: registry.DisplayService{
			LCD: lcd2k.New(),
		},
		Hardware: registry.HardwareService{
			Queue: queue.New(),
		},
		Repository: registry.RepositoryService{
			Machine:        machine.New(module.SQLite.DB),
			PaymentChannel: payment_channel.New(module.SQLite.DB),
			Slot:           slot.New(module.SQLite.DB),
			Transaction:    transaction.New(module.SQLite.DB),
		},
	}
}
