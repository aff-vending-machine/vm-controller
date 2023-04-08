package machine

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/machine/response"
)

func (uc *usecaseImpl) Get(ctx context.Context) (*response.Machine, error) {
	machine, err := uc.machineRepo.FindOne(ctx, []string{"id:=:1"})
	if err != nil {
		return nil, err
	}

	return response.ToMachine(machine), nil
}