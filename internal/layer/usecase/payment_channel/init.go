package payment_channel

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository"
)

type usecaseImpl struct {
	paymentChannelRepo repository.PaymentChannel
}

func New(pr repository.PaymentChannel) *usecaseImpl {
	return &usecaseImpl{pr}
}
