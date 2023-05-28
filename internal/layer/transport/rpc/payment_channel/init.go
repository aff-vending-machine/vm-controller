package payment_channel

import "vm-controller/internal/core/interface/payment_channel"

type rpcImpl struct {
	usecase payment_channel.Usecase
}

func New(uc payment_channel.Usecase) *rpcImpl {
	return &rpcImpl{uc}
}
