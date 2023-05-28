package payment_channel

import (
	"context"

	"vm-controller/internal/layer/usecase/payment_channel/request"
	"vm-controller/internal/layer/usecase/payment_channel/response"
	"vm-controller/pkg/helpers/db"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) GetOne(ctx context.Context, req *request.GetOne) (*response.PaymentChannel, error) {
	if v := validate.Struct(req); !v.Validate() {
		return nil, errors.Wrap(v.Errors.OneError(), "invalid request")
	}

	entity, err := uc.paymentChannelRepo.FindOne(ctx, db.NewQuery().AddWhere("channel = ?", req.Channel))
	if err != nil {
		return nil, err
	}

	return response.ToModel(entity), nil
}
