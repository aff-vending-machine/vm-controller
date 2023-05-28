package payment_channel

import (
	"context"

	"vm-controller/internal/layer/usecase/payment_channel/request"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) Create(ctx context.Context, req *request.Create) error {
	if v := validate.Struct(req); !v.Validate() {
		return errors.Wrap(v.Errors.OneError(), "invalid request")
	}

	_, err := uc.paymentChannelRepo.Create(ctx, req.ToEntity())
	if err != nil {
		return errors.Wrap(err, "unable to insert payment channel")
	}

	return nil
}
