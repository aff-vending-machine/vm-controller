package registry

import (
	"vm-controller/internal/boot/modules"
	"vm-controller/internal/layer/flow"
	"vm-controller/internal/layer/transport/rpc/machine"
	"vm-controller/internal/layer/transport/rpc/slot"
	"vm-controller/internal/layer/transport/rpc/transaction"
	"vm-controller/internal/layer/transport/websocket/frontend"
)

func NewTransport(uc modules.Usecase, fw flow.Jetts) modules.Transport {
	return modules.Transport{
		RPC: modules.RPCTransport{
			Machine:     machine.New(uc.Machine),
			Slot:        slot.New(uc.Slot),
			Transaction: transaction.New(uc.Transaction),
		},
		WebSocket: modules.WebSocketTransport{
			Frontend: frontend.New(fw),
		},
	}
}
