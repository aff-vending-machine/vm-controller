package slot

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/slot/request"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) Create(ctx context.Context, req *request.Create) error {
	if v := validate.Struct(req); !v.Validate() {
		return errors.Wrap(v.Errors.OneError(), "invalid request")
	}

	_, err := uc.slotRepo.Create(ctx, req.ToEntity())
	if err != nil {
		return errors.Wrap(err, "unable to insert slot")
	}

	return nil
}
