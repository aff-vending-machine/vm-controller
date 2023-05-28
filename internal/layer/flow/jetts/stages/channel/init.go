package channel

import (
	"vm-controller/internal/core/domain/entity"
	"vm-controller/internal/core/interface/payment_channel"
	"vm-controller/internal/core/interface/transaction"
	"vm-controller/internal/layer/service/websocket"
)

type stageImpl struct {
	paymentChannelRepo payment_channel.Repository
	transactionRepo    transaction.Repository
	frontendWs         websocket.Frontend
	channels           []entity.PaymentChannel
}

func New(pr payment_channel.Repository, tr transaction.Repository, fw websocket.Frontend) *stageImpl {
	return &stageImpl{pr, tr, fw, make([]entity.PaymentChannel, 0)}
}
