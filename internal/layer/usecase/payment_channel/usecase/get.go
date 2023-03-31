package payment_channel_usecase

import (
	"context"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/payment_channel/response"
)

func (uc *usecaseImpl) Get(ctx context.Context, filter []string) (*response.PaymentChannel, error) {
	entity, err := uc.paymentChannelRepo.FindOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return response.ToModel(entity), nil
}
