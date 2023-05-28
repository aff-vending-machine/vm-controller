package model

import "vm-controller/internal/core/domain/entity"

type PaymentChannel struct {
	Channel      string `json:"channel"` // primary key
	Name         string `json:"name"`
	Vendor       string `json:"vendor"`
	IsEnable     bool   `json:"is_enable"`
	Host         string `json:"host"`
	MerchantID   string `json:"merchant_id"`
	MerchantName string `json:"merchant_name"`
	BillerCode   string `json:"biller_code"`
	BillerID     string `json:"biller_id"`
	Token        string `json:"token"`
	StoreID      string `json:"store_id"`
	TerminalID   string `json:"terminal_id"`
}

func ToPaymentChannel(e *entity.PaymentChannel) *PaymentChannel {
	return &PaymentChannel{
		Channel:      e.Channel,
		Name:         e.Name,
		Vendor:       e.Vendor,
		IsEnable:     e.IsEnable,
		Host:         e.Host,
		MerchantID:   e.MerchantID,
		MerchantName: e.MerchantName,
		BillerCode:   e.BillerCode,
		BillerID:     e.BillerID,
		Token:        e.Token,
		StoreID:      e.StoreID,
		TerminalID:   e.TerminalID,
	}
}
