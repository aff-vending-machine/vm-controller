package registry

import (
	"github.com/aff-vending-machine/vm-controller/config"
	"github.com/aff-vending-machine/vm-controller/internal/core/infra/http"
	"github.com/aff-vending-machine/vm-controller/internal/core/infra/sqlite"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/api"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/api/ksher"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/api/link2500"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/asset"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/asset/fonts"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/asset/images"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/display"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/display/lcd2k"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/hardware"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/hardware/queue"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository/machine"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository/payment_channel"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository/slot"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository/transaction"
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
	Ksher    api.Ksher
	Link2500 api.Link2500
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
			ksher.New(rest.Client),
			link2500.New(rest.Client),
		},
		AssetDriven{
			fonts.New(cfg.App.Asset),
			images.New(cfg.App.Asset),
		},
		DisplayDriven{
			lcd2k.New(cfg.RasPi),
		},
		HardwareDriven{
			queue.New(cfg.Redis),
		},
		RepositoryDriven{
			machine.New(db),
			payment_channel.New(db),
			slot.New(db),
			transaction.New(db),
		},
	}
}
