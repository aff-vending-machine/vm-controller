package registry

import (
	"vm-controller/internal/boot/modules"
	"vm-controller/internal/layer/flow"
	"vm-controller/internal/layer/transport/rpc/machine"
	"vm-controller/internal/layer/transport/rpc/payment_channel"
	"vm-controller/internal/layer/transport/rpc/product"
	"vm-controller/internal/layer/transport/rpc/slot"
	"vm-controller/internal/layer/transport/rpc/transaction"
	"vm-controller/internal/layer/transport/websocket/frontend"
)

func NewTransport(uc modules.Usecase, fw flow.Jetts) modules.Transport {
	return modules.Transport{
		RPC: modules.RPCTransport{
			Machine:        machine.New(uc.Machine),
			PaymentChannel: payment_channel.New(uc.PaymentChannel),
			Product:        product.New(uc.Product),
			Slot:           slot.New(uc.Slot),
			Transaction:    transaction.New(uc.Transaction),
		},
		WebSocket: modules.WebSocketTransport{
			Frontend: frontend.New(fw),
		},
	}
}
