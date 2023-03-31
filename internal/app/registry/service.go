package registry

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/config"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/infra/postgresql"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/infra/sqlite"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/api"
	lugentpay_api "github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/api/lugentpay"
	mail_api "github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/api/mail"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/api/smartedcv2"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/hardware"
	queue_hardware "github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/hardware/queue"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/repository"
	customer_repository "github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/repository/customer"
	machine_repository "github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/repository/machine"
	payment_channel_repository "github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/repository/payment_channel"
	slot_repository "github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/repository/slot"
	transaction_repository "github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/repository/transaction"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/serial"
	// smartedc_serial "github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/serial/smartedc"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/ws"
	ui_ws "github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/ws/ui"
	"gorm.io/gorm"
)

// Interface Adapter layers (driven)
type AppDriven struct {
	API        APIDriven
	WebSocket  WebSocketDriven
	Hardware   HardwareDriven
	Repository RepositoryDriven
	Serial     SerialDriven
}

type APIDriven struct {
	LugentPay api.LugentPay
	Mail      api.Mail
}

type WebSocketDriven struct {
	UI ws.UI
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

type SerialDriven struct {
	SmartEDC serial.SmartEDC
}

func NewAppDriven(cfg config.BootConfig) AppDriven {
	var db *gorm.DB
	if cfg.PostgreSQL.Enable {
		db = postgresql.New(cfg.PostgreSQL)
	} else {
		db = sqlite.New(cfg.SQLite)
	}

	return AppDriven{
		APIDriven{
			lugentpay_api.New(),
			mail_api.New(cfg.Mail),
			// smartedcv2.New(),
		},
		WebSocketDriven{
			ui_ws.New(),
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
		SerialDriven{
			// smartedc_serial.New(cfg.SmartEDC),
			smartedcv2.New(),
		},
	}
}
