package preload

import (
	"context"
	"fmt"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/entity"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/slot"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/slot/request"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/boot"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/utils/errs"
	"github.com/rs/zerolog/log"
)

func InitMachineSlot(uc slot.Usecase) {
	ctx := context.TODO()

	slot, err := uc.Get(ctx, []string{"id:=:1"})
	if errs.Not(err, "record not found") {
		boot.TerminateWhenError(err)
		return
	}
	if slot != nil {
		log.Info().Msg("slot is already exist")
		return
	}

	for i := 0; i < 10; i++ {
		uc.Set(ctx, &request.Set{
			Code:     fmt.Sprintf("0%d", i+10),
			IsEnable: true,
			Product: &entity.Product{
				SKU:      fmt.Sprintf("Product#%d", i+1),
				Name:     fmt.Sprintf("Product#%d", i+1),
				ImageURL: "https://image-placeholder.com/images/actual-size/75x100.png",
				Price:    0,
			},
			Capacity: 10,
			Stock:    10,
		},
		)
	}

	for i := 10; i < 50; i++ {
		uc.Set(ctx, &request.Set{
			Code:     fmt.Sprintf("0%d", i+10),
			IsEnable: true,
			Product: &entity.Product{
				SKU:      fmt.Sprintf("Product#%d", i+1),
				Name:     fmt.Sprintf("Product#%d", i+1),
				ImageURL: "https://image-placeholder.com/images/actual-size/75x100.png",
				Price:    10,
			},
			Capacity: 10,
			Stock:    10,
		},
		)
	}
}
