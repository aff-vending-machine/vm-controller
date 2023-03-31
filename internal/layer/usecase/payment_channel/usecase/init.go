package payment_channel_usecase

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/repository"
)

type usecaseImpl struct {
	paymentChannelRepo repository.PaymentChannel
}

func New(pr repository.PaymentChannel) *usecaseImpl {
	return &usecaseImpl{pr}
}
