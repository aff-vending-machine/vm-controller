package slot

import (
	"context"

	"vm-controller/internal/layer/usecase/slot/request"
	"vm-controller/internal/layer/usecase/slot/response"
	"vm-controller/pkg/helpers/db"

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
