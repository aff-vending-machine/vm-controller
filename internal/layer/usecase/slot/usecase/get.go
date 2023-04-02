package slot_usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/slot/response"
)

func (uc *usecaseImpl) Get(ctx context.Context, filter []string) ([]response.Slot, error) {
	slots, err := uc.slotRepo.FindMany(ctx, filter)
	if err != nil {
		return nil, err
	}

	result := toResponseList(slots)

	return result, nil
}
