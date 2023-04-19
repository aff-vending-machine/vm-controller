package ws

import (
	"github.com/aff-vending-machine/vm-controller/internal/boot/registry"
	"github.com/aff-vending-machine/vm-controller/internal/layer/transport/rpc/machine"
	"github.com/aff-vending-machine/vm-controller/internal/layer/transport/rpc/slot"
	"github.com/aff-vending-machine/vm-controller/internal/layer/transport/rpc/transaction"
	"github.com/aff-vending-machine/vm-controller/internal/layer/transport/websocket/frontend"
)

func NewTransport(uc registry.Usecase, fw registry.Flow) registry.Transport {
	return registry.Transport{
		RPC: registry.RPCTransport{
			Machine:     machine.New(uc.Machine),
			Slot:        slot.New(uc.Slot),
			Transaction: transaction.New(uc.Transaction),
		},
		WebSocket: registry.WebSocketTransport{
			Frontend: frontend.New(fw.Jetts),
		},
	}
}
