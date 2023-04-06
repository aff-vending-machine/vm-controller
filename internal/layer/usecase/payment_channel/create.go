package payment_channel

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/payment_channel/request"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) Create(ctx context.Context, req *request.Create) error {
	if v := validate.Struct(req); !v.Validate() {
		return errors.Wrap(v.Errors.OneError(), "invalid request")
	}

	err := uc.paymentChannelRepo.InsertOne(ctx, req.ToEntity())
	if err != nil {
		return errors.Wrap(err, "unable to insert device")
	}

	return nil
}
