package modules

import (
	"vm-controller/internal/core/interface/machine"
	"vm-controller/internal/core/interface/payment_channel"
	"vm-controller/internal/core/interface/slot"
	"vm-controller/internal/core/interface/transaction"
	"vm-controller/internal/layer/service/api"
	"vm-controller/internal/layer/service/hardware"
	"vm-controller/internal/layer/service/websocket"
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
