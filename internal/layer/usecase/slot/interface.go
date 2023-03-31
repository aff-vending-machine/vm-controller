package slot

import (
	"context"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/slot/request"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/slot/response"
)

type Usecase interface {
	Get(context.Context, []string) ([]response.Slot, error)
	Set(context.Context, *request.Set) error
	Clear(context.Context, []string) error
	SetStock(context.Context, *request.SetStock) error
	GetOne(context.Context, uint) (*response.Slot, error)
	Update(context.Context, *request.Update) error
}
