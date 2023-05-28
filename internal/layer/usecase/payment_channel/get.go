package payment_channel

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/payment_channel/request"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/payment_channel/response"
	"github.com/aff-vending-machine/vm-controller/pkg/helpers/db"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) Get(ctx context.Context, req *request.Get) (*response.PaymentChannel, error) {
	if v := validate.Struct(req); !v.Validate() {
		return nil, errors.Wrap(v.Errors.OneError(), "invalid request")
	}

	entity, err := uc.paymentChannelRepo.FindOne(ctx, db.NewQuery().AddWhere("channel = ?", req.Channel))
	if err != nil {
		return nil, err
	}

	return response.ToModel(entity), nil
}
