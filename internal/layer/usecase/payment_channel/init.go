package payment_channel

import (
	"vm-controller/internal/core/interface/payment_channel"
)

type usecaseImpl struct {
	paymentChannelRepo payment_channel.Repository
}

func New(pr payment_channel.Repository) *usecaseImpl {
	return &usecaseImpl{pr}
}
