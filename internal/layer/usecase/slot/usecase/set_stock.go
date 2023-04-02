package slot_usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/slot/request"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) SetStock(ctx context.Context, req *request.SetStock) error {
	if v := validate.Struct(req); !v.Validate() {
		return errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	_, err := uc.slotRepo.UpdateMany(ctx, req.ToFilter(), req.ToJson())
	if err != nil {
		return errors.Wrapf(err, "update slot %s failed", req.ID)
	}

	return nil
}
