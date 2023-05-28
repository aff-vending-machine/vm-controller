package registry

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/interface/machine"
	"github.com/aff-vending-machine/vm-controller/internal/core/interface/payment_channel"
	"github.com/aff-vending-machine/vm-controller/internal/core/interface/slot"
	"github.com/aff-vending-machine/vm-controller/internal/core/interface/transaction"
)

// Usecase layers
type Usecase struct {
	Machine        machine.Usecase
	PaymentChannel payment_channel.Usecase
	Slot           slot.Usecase
	Transaction    transaction.Usecase
}
