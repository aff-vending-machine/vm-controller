package machine_usecase

import (
	"context"
	"strings"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/machine/request"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) StartUp(ctx context.Context, req *request.StartUp) (string, error) {
	machine, err := uc.machineRepo.FindOne(ctx, []string{"id:=:1"})
	if err != nil && !strings.Contains(err.Error(), "not found") {
		return "", err
	}
	if err != nil {
		machine = req.ToEntity()
		err = uc.machineRepo.InsertOne(ctx, machine)
		if err != nil {
			return "", err
		}
	}

	log.Debug().Str("serial_number", machine.SerialNumber).Msg("this machine is registered")

	return machine.SerialNumber, nil
}
