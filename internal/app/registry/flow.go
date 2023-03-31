package registry

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/flow"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/flow/thaitropica"
)

// Usecase layers
type AppFlow struct {
	ThaiTropica interface{ flow.ThaiTropica }
}

func NewAppFlow(adapter AppDriven) AppFlow {
	return AppFlow{
		thaitropica.New(
			adapter.API.LugentPay,
			adapter.API.Mail,
			adapter.Hardware.Queue,
			adapter.Repository.Customer,
			adapter.Repository.Machine,
			adapter.Repository.PaymentChannel,
			adapter.Repository.Slot,
			adapter.Repository.Transaction,
			adapter.Serial.SmartEDC,
			adapter.WebSocket.UI,
		),
	}
}
