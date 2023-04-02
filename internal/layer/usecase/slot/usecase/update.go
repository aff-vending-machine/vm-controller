package slot_usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/slot/request"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) Update(ctx context.Context, req *request.Update) error {
	if v := validate.Struct(req); !v.Validate() {
		return errors.Wrap(v.Errors.OneError(), "invalid request")
	}

	_, err := uc.slotRepo.UpdateMany(ctx, req.ToFilter(), req.ToJson())
	if err != nil {
		return errors.Wrap(err, "unable to replace role")
	}

	return nil
}
