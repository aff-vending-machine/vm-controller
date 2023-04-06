package machine

import (
	"context"
	"strings"

	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/machine/request"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/machine/response"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) StartUp(ctx context.Context, req *request.StartUp) (*response.Machine, error) {
	machine, err := uc.machineRepo.FindOne(ctx, []string{"id:=:1"})
	if err != nil && !strings.Contains(err.Error(), "not found") {
		return nil, err
	}
	if err != nil {
		machine = req.ToEntity()
		err = uc.machineRepo.InsertOne(ctx, machine)
		if err != nil {
			return nil, err
		}
	}

	log.Debug().Str("serial_number", machine.SerialNumber).Msg("this machine is registered")

	return response.ToMachine(machine), nil
}
