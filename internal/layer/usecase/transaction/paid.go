package transaction

import (
	"context"

	"vm-controller/internal/layer/usecase/transaction/request"
)

func (uc *usecaseImpl) Paid(ctx context.Context, req *request.Paid) error {
	return nil
}
