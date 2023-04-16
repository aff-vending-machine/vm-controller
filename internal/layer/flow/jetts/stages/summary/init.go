package summary

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/websocket"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase"
)

type stageImpl struct {
	displayUc       usecase.Screen
	transactionRepo repository.Transaction
	frontendWs      websocket.Frontend
}

func New(du usecase.Screen, tr repository.Transaction, fw websocket.Frontend) *stageImpl {
	return &stageImpl{du, tr, fw}
}
