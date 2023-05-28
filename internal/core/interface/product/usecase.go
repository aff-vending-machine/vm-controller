package product

import (
	"context"

	"vm-controller/internal/layer/usecase/product/request"
)

type Usecase interface {
	Set(context.Context, *request.Set) error
}
