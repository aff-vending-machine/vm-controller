package summary

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase"
)

type stageImpl struct {
	displayUc       usecase.Display
	transactionRepo repository.Transaction
}

func New(du usecase.Display, tr repository.Transaction) *stageImpl {
	return &stageImpl{du, tr}
}
