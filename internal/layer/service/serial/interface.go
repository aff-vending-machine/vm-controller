package serial

import (
	"context"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/smartedc"
)

type SmartEDC interface {
	Sale(ctx context.Context, req *smartedc.SaleRequest) (*smartedc.SaleResult, error)
	Void(ctx context.Context, req *smartedc.VoidRequest) (*smartedc.VoidResult, error)
}
