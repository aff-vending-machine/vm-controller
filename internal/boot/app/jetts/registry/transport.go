package registry

import (
	"github.com/aff-vending-machine/vm-controller/internal/boot/modules"
	"github.com/aff-vending-machine/vm-controller/internal/layer/flow"
	"github.com/aff-vending-machine/vm-controller/internal/layer/transport/rpc/machine"
	"github.com/aff-vending-machine/vm-controller/internal/layer/transport/rpc/slot"
	"github.com/aff-vending-machine/vm-controller/internal/layer/transport/rpc/transaction"
	"github.com/aff-vending-machine/vm-controller/internal/layer/transport/websocket/frontend"
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
