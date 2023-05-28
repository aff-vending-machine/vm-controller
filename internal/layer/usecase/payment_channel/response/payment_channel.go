package response

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
)

type PaymentChannel struct {
	ID           uint   `json:"id" gorm:"primarykey"`
	Name         string `json:"name"`
	Channel      string `json:"channel"`
	Vendor       string `json:"vendor"`
	Active       bool   `json:"active"`
	Host         string `json:"host"`
	MerchantID   string `json:"merchant_id"`
	MerchantName string `json:"merchant_name"`
	BillerCode   string `json:"biller_code"`
	BillerID     string `json:"biller_id"`
	StoreID      string `json:"store_id"`
	TerminalID   string `json:"terminal_id"`
}

func ToModel(channel *entity.PaymentChannel) *PaymentChannel {
	return &PaymentChannel{
		ID:           channel.ID,
		Name:         channel.Name,
		Channel:      channel.Channel,
		Vendor:       channel.Vendor,
		Active:       channel.Active,
		Host:         channel.Host,
		MerchantID:   channel.MerchantID,
		MerchantName: channel.MerchantName,
		BillerCode:   channel.BillerCode,
		BillerID:     channel.BillerID,
		StoreID:      channel.StoreID,
		TerminalID:   channel.TerminalID,
	}
}
