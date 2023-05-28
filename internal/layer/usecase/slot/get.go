package slot

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/slot/response"
)

func (uc *usecaseImpl) Get(ctx context.Context) ([]response.Slot, error) {
	slots, err := uc.slotRepo.FindMany(ctx, []string{})
	if err != nil {
		return nil, err
	}

	return response.ToSlotList(slots), nil
}
