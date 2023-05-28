package slot

import (
	"context"

	"vm-controller/internal/layer/usecase/slot/response"
	"vm-controller/pkg/helpers/db"
)

func (uc *usecaseImpl) Get(ctx context.Context) ([]response.Slot, error) {
	slots, err := uc.slotRepo.FindMany(ctx, db.NewQuery())
	if err != nil {
		return nil, err
	}

	return response.ToSlotList(slots), nil
}
