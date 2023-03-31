package slot_usecase

import (
	"context"
	"fmt"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/slot/response"
)

func (uc *usecaseImpl) GetOne(ctx context.Context, id uint) (*response.Slot, error) {
	slot, err := uc.slotRepo.FindOne(ctx, []string{fmt.Sprintf("id:=:%d", id)})
	if err != nil {
		return nil, err
	}

	return response.ToModel(slot), nil
}
