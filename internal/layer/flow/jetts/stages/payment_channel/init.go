package payment_channel

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/websocket"
)

type stageImpl struct {
	paymentChannelRepo repository.PaymentChannel
	transactionRepo    repository.Transaction
	frontendWs         websocket.Frontend
	channels           []entity.PaymentChannel
}

func New(pr repository.PaymentChannel, tr repository.Transaction, fw websocket.Frontend) *stageImpl {
	return &stageImpl{pr, tr, fw, make([]entity.PaymentChannel, 0)}
}
