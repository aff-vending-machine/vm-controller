package payment_channel

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase"
)

type stageImpl struct {
	displayUc          usecase.Screen
	paymentChannelRepo repository.PaymentChannel
	transactionRepo    repository.Transaction
	channels           []entity.PaymentChannel
}

func New(dg usecase.Screen, pr repository.PaymentChannel, tr repository.Transaction) *stageImpl {
	return &stageImpl{dg, pr, tr, make([]entity.PaymentChannel, 0)}
}
