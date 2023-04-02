package registry

import (
	"github.com/aff-vending-machine/vm-controller/config"
	"github.com/aff-vending-machine/vm-controller/internal/core/infra/http"
	"github.com/aff-vending-machine/vm-controller/internal/core/infra/sqlite"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/api"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/hardware"
	queue_hardware "github.com/aff-vending-machine/vm-controller/internal/layer/service/hardware/queue"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository"
	customer_repository "github.com/aff-vending-machine/vm-controller/internal/layer/service/repository/customer"
	machine_repository "github.com/aff-vending-machine/vm-controller/internal/layer/service/repository/machine"
	payment_channel_repository "github.com/aff-vending-machine/vm-controller/internal/layer/service/repository/payment_channel"
	slot_repository "github.com/aff-vending-machine/vm-controller/internal/layer/service/repository/slot"
	transaction_repository "github.com/aff-vending-machine/vm-controller/internal/layer/service/repository/transaction"
)

// Interface Adapter layers (driven)
type AppDriven struct {
	API        APIDriven
	Asset      AssetDriven
	Display    DisplayDriven
	Hardware   HardwareDriven
	Repository RepositoryDriven
}

type APIDriven struct {
	Ksher api.Ksher
	EDC   api.Link2500
}

type AssetDriven struct {
	Fonts  asset.Fonts
	Images asset.Images
}

type DisplayDriven struct {
	LCD display.LCD
}

type HardwareDriven struct {
	Queue hardware.Queue
}

type RepositoryDriven struct {
	Customer       repository.Customer
	Machine        repository.Machine
	PaymentChannel repository.PaymentChannel
	Slot           repository.Slot
	Transaction    repository.Transaction
}

func NewAppDriven(cfg config.BootConfig) AppDriven {
	rest := http.New(cfg.HTTP)
	db := sqlite.New(cfg.SQLite)

	return AppDriven{
		APIDriven{
			ksher.New(rest),
			link2500.New(rest),
		},
		AssetDriven{
			fonts.New(cfg.App.Asset),
			images.New(cfg.App.Asset),
		},
		DisplayDriven{
			lcd_display.New(cfg.RasPi),
		},
		HardwareDriven{
			queue_hardware.New(cfg.Redis),
		},
		RepositoryDriven{
			customer_repository.New(db),
			machine_repository.New(db),
			payment_channel_repository.New(db),
			slot_repository.New(db),
			transaction_repository.New(db),
		},
	}
}
