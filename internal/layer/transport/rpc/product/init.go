package product

import "vm-controller/internal/core/interface/product"

type rpcImpl struct {
	usecase product.Usecase
}

func New(uc product.Usecase) *rpcImpl {
	return &rpcImpl{uc}
}
