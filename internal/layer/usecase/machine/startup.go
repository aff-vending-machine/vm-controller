package machine

import (
	"context"

	"vm-controller/internal/layer/usecase/machine/request"
	"vm-controller/internal/layer/usecase/machine/response"
	"vm-controller/pkg/helpers/db"
	"vm-controller/pkg/helpers/errs"

	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) StartUp(ctx context.Context, req *request.StartUp) (*response.Machine, error) {
	machine, err := uc.machineRepo.FindOne(ctx, db.NewQuery().AddWhere("id = ?", 1))
	if errs.Not(err, "not found") {
		return nil, err
	}
	if err != nil {
		machine = req.ToEntity()
		_, err = uc.machineRepo.Create(ctx, machine)
		if err != nil {
			return nil, err
		}
	}

	result := response.ToMachine(machine)

	uc.topic.RegisterMachine(ctx, machine, result)
	log.Debug().Str("serial_number", machine.SerialNumber).Msg("this machine is registered")

	return result, nil
}
