package payment_channel

import (
	"context"

	"vm-controller/internal/layer/usecase/payment_channel/response"
	"vm-controller/pkg/helpers/db"
)

func (uc *usecaseImpl) Get(ctx context.Context) ([]response.PaymentChannel, error) {
	entity, err := uc.paymentChannelRepo.FindMany(ctx, db.NewQuery())
	if err != nil {
		return nil, err
	}

	return response.ToPaymentChannelList(entity), nil
}
