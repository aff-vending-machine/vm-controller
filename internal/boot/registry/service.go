package registry

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/api"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/hardware"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository"
	websocket "github.com/aff-vending-machine/vm-controller/internal/layer/service/websocket"
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
	Machine        repository.Machine
	PaymentChannel repository.PaymentChannel
	Slot           repository.Slot
	Transaction    repository.Transaction
}

type WebSocketService struct {
	Frontend websocket.Frontend
}
