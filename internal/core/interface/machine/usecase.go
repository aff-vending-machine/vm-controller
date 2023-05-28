package machine

import (
	"context"

	"vm-controller/internal/layer/usecase/machine/request"
	"vm-controller/internal/layer/usecase/machine/response"
)

type Usecase interface {
	Get(context.Context) (*response.Machine, error)
	Set(context.Context, *request.Set) error
	Healthy(context.Context) error
	StartUp(context.Context, *request.StartUp) (*response.Machine, error)
	Reset(context.Context) error
	Emergency(context.Context) error
	OpenGate(context.Context) error
}
