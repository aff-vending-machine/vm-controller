package slot

import (
	"context"

	"vm-controller/internal/layer/usecase/slot/request"
	"vm-controller/internal/layer/usecase/slot/response"
)

type Usecase interface {
	List(context.Context, *request.Filter) ([]response.Slot, error)
	Create(context.Context, *request.Create) error
	Get(context.Context) ([]response.Slot, error)
	Set(context.Context, *request.Set) error
}
