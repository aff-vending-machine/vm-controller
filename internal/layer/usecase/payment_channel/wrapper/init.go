package payment_channel_wrapper

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/payment_channel"
)

type wrapperImpl struct {
	usecase payment_channel.Usecase
}

func New(uc payment_channel.Usecase) *wrapperImpl {
	return &wrapperImpl{uc}
}
