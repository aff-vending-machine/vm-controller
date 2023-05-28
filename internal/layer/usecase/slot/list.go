package slot

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/slot/request"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/slot/response"
	"github.com/aff-vending-machine/vm-controller/pkg/helpers/db"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) List(ctx context.Context, req *request.Filter) ([]response.Slot, error) {
	if v := validate.Struct(req); !v.Validate() {
		return nil, errors.Wrap(v.Errors.OneError(), "invalid request")
	}

	slots, err := uc.slotRepo.FindMany(ctx, db.NewQuery().PtrLimit(req.Limit).PtrOffset(req.Offset).PtrOrder(req.SortBy))
	if err != nil {
		return nil, errors.Wrap(err, "unable to find slots")
	}

	return response.ToSlotList(slots), nil
}
