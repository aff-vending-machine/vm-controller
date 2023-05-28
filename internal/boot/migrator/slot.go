package migrator

import (
	"context"
	"fmt"

	"github.com/aff-vending-machine/vm-controller/internal/core/interface/slot"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/slot/model"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/slot/request"
	"github.com/aff-vending-machine/vm-controller/pkg/boot"
	"github.com/aff-vending-machine/vm-controller/pkg/helpers/errs"
	"github.com/rs/zerolog/log"
)

func InitMachineSlot(uc slot.Usecase) {
	ctx := context.TODO()

	slots, err := uc.List(ctx, &request.Filter{})
	if errs.Not(err, "record not found") {
		boot.TerminateWhenError(err)
		return
	}
	if len(slots) != 0 {
		log.Info().Msg("slot is already exist")
		return
	}

	setItem(uc, 10, 11, 10, 10, 140)
	setItem(uc, 12, 14, 10, 10, 99)
	setItem(uc, 15, 16, 10, 10, 89)
	setItem(uc, 17, 19, 10, 10, 99)
	setItem(uc, 20, 24, 10, 10, 20)
	setItem(uc, 25, 28, 10, 10, 25)
	setItem(uc, 30, 34, 10, 10, 10)
	setItem(uc, 35, 38, 10, 10, 15)
	setItem(uc, 40, 48, 10, 10, 75)
	setItem(uc, 50, 59, 10, 10, 49)

	log.Info().Msg("create default slots")
}

func setItem(uc slot.Usecase, codeFrom int, codeEnd int, stock int, capacity int, price float64) {
	ctx := context.TODO()

	for i := codeFrom; i <= codeEnd; i++ {
		data := &request.Create{
			Code:     fmt.Sprintf("%03d", i),
			Stock:    stock,
			Capacity: capacity,
			Product: &model.Product{
				SKU:      fmt.Sprintf("P%05.0f", price),
				Name:     fmt.Sprintf("Slot %03d", i),
				ImageURL: "https://image-placeholder.com/images/actual-size/75x100.png",
				Price:    price,
			},
			IsEnable: true,
		}

		err := uc.Create(ctx, data)
		if err != nil {
			log.Error().Err(err).Str("code", data.Code).Msg("unable to create slot")
		}
	}
}
