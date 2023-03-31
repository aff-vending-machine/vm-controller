package payment_channel

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/entity"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/repository"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/ws"
)

type Stage struct {
	paymentChannelRepo repository.PaymentChannel
	ui                 ws.UI
	channels           []entity.PaymentChannel
}

func New(pr repository.PaymentChannel, u ws.UI) *Stage {
	return &Stage{pr, u, []entity.PaymentChannel{}}
}
