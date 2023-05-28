package modules

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/interface/machine"
	"github.com/aff-vending-machine/vm-controller/internal/core/interface/payment_channel"
	"github.com/aff-vending-machine/vm-controller/internal/core/interface/slot"
	"github.com/aff-vending-machine/vm-controller/internal/core/interface/transaction"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/api"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/hardware"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/websocket"
)

// Interface Adapter layers (driven)
type Service struct {
	API        APIService
	Hardware   HardwareService
	Repository RepositoryService
	WebSocket  WebSocketService
}

type APIService struct {
	Ksher    api.Ksher
	Link2500 api.Link2500
	Topic    api.Topic
}
type HardwareService struct {
	Queue hardware.Queue
}

type RepositoryService struct {
	Machine        machine.Repository
	PaymentChannel payment_channel.Repository
	Slot           slot.Repository
	Transaction    transaction.Repository
}

type WebSocketService struct {
	Frontend websocket.Frontend
}
