package machine

import (
	"context"

	"vm-controller/internal/layer/usecase/machine/response"
	"vm-controller/pkg/helpers/db"
)

func (uc *usecaseImpl) Get(ctx context.Context) (*response.Machine, error) {
	machine, err := uc.machineRepo.FindOne(ctx, db.NewQuery().AddWhere("id = ?", 1))
	if err != nil {
		return nil, err
	}

	return response.ToMachine(machine), nil
}
