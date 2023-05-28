package response

import (
	"vm-controller/internal/core/domain/entity"
	"vm-controller/internal/layer/usecase/payment_channel/model"
)

type PaymentChannel = model.PaymentChannel

func ToModel(channel *entity.PaymentChannel) *PaymentChannel {
	return &PaymentChannel{
		Name:         channel.Name,
		Channel:      channel.Channel,
		Vendor:       channel.Vendor,
		IsEnable:     channel.IsEnable,
		Host:         channel.Host,
		MerchantID:   channel.MerchantID,
		MerchantName: channel.MerchantName,
		BillerCode:   channel.BillerCode,
		BillerID:     channel.BillerID,
		StoreID:      channel.StoreID,
		TerminalID:   channel.TerminalID,
	}
}

func ToPaymentChannelList(entities []entity.PaymentChannel) []model.PaymentChannel {
	results := make([]model.PaymentChannel, len(entities))
	for i, e := range entities {
		results[i] = *model.ToPaymentChannel(&e)
	}

	return results
}
