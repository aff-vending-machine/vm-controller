package topic

import (
	"context"
	"encoding/json"

	"vm-controller/internal/core/domain/entity"
	"vm-controller/internal/layer/usecase/machine/model"
)

const REGISTER_MACHINE_KEY = "center.machine.register"

type RegisterMachineRequest struct {
	Data model.Machine `json:"data"`
}

func (a *apiImpl) RegisterMachine(ctx context.Context, machine *entity.Machine, data *model.Machine) error {
	req := RegisterMachineRequest{Data: *data}
	breq, _ := json.Marshal(req)

	err := a.Client.EmitTopic(ctx, machine.SerialNumber, machine.Center, REGISTER_MACHINE_KEY, breq)
	if err != nil {
		return err
	}

	return nil
}
