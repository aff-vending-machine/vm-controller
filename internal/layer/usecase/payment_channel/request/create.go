package request

import (
	"vm-controller/internal/core/domain/entity"
)

type Create struct {
	Name         string `json:"name"`
	Channel      string `json:"channel"`
	Vendor       string `json:"vendor"`
	Host         string `json:"host"`
	MerchantID   string `json:"merchant_id"`
	MerchantName string `json:"merchant_name"`
	BillerCode   string `json:"biller_code"`
	BillerID     string `json:"biller_id"`
	StoreID      string `json:"store_id"`
	TerminalID   string `json:"terminal_id"`
}

func (r *Create) ToEntity() *entity.PaymentChannel {
	return &entity.PaymentChannel{
		Name:         r.Name,
		Channel:      r.Channel,
		Vendor:       r.Vendor,
		IsEnable:     true,
		Host:         r.Host,
		MerchantID:   r.MerchantID,
		MerchantName: r.MerchantName,
		BillerCode:   r.BillerCode,
		BillerID:     r.BillerID,
		StoreID:      r.StoreID,
		TerminalID:   r.TerminalID,
	}
}
