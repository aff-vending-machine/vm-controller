package machine

import (
	"context"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/machine/request"
)

type Usecase interface {
	Healthy(context.Context) error
	StartUp(context.Context, *request.StartUp) (string, error)
	Reset(context.Context) error
	Emergency(context.Context) error
	OpenGate(context.Context) error
}
