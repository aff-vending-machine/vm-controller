package modules

import (
	"vm-controller/internal/core/interface/machine"
	"vm-controller/internal/core/interface/payment_channel"
	"vm-controller/internal/core/interface/product"
	"vm-controller/internal/core/interface/slot"
	"vm-controller/internal/core/interface/transaction"
)

// Usecase layers
type Usecase struct {
	Machine        machine.Usecase
	PaymentChannel payment_channel.Usecase
	Product        product.Usecase
	Slot           slot.Usecase
	Transaction    transaction.Usecase
}
