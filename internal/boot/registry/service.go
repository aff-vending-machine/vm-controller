package registry

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/api"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/asset"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/display"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/hardware"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository"
)

// Interface Adapter layers (driven)
type Service struct {
	API        APIService
	Asset      AssetService
	Display    DisplayService
	Hardware   HardwareService
	Repository RepositoryService
}

type APIService struct {
	Ksher    api.Ksher
	Link2500 api.Link2500
	Topic    api.Topic
}

type AssetService struct {
	Fonts  asset.Fonts
	Images asset.Images
}

type DisplayService struct {
	LCD display.LCD
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
