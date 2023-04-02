package slot_usecase

import (
	"context"
	"strings"

	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/slot/request"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) Set(ctx context.Context, req *request.Set) error {
	if v := validate.Struct(req); !v.Validate() {
		return errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	_, err := uc.slotRepo.FindOne(ctx, req.ToFilter())
	if err != nil && strings.Contains(err.Error(), "not found") {
		slot := req.ToEntity()
		err := uc.slotRepo.InsertOne(ctx, &slot)
		if err != nil {
			return errors.Wrapf(err, "insert slot %s failed", req.Code)
		}

		return nil
	}
	if err != nil {
		return errors.Wrapf(err, "find slot %s failed", req.Code)
	}

	data := req.ToJson()
	_, err = uc.slotRepo.UpdateMany(ctx, req.ToFilter(), data)
	if err != nil {
		return errors.Wrapf(err, "update slot %s failed", req.Code)
	}

	return nil
}
