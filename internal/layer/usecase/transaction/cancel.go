package transaction

import (
	"context"

	"vm-controller/internal/layer/usecase/transaction/request"
)

func (uc *usecaseImpl) Cancel(ctx context.Context, req *request.Cancel) error {
	return nil
}
