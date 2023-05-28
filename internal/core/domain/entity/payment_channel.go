package entity

import (
	"errors"
	"time"
)

type PaymentChannel struct {
	ID           uint      `json:"id" gorm:"primarykey"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string    `json:"name"`
	Channel      string    `json:"channel"`
	Vendor       string    `json:"vendor"`
	IsEnable     bool      `json:"is_enable"`
	Host         string    `json:"host"`
	MerchantID   string    `json:"merchant_id"`
	MerchantName string    `json:"merchant_name"`
	BillerCode   string    `json:"biller_code"`
	BillerID     string    `json:"biller_id"`
	Token        string    `json:"token"`
	StoreID      string    `json:"store_id"`
	TerminalID   string    `json:"terminal_id"`
}

func (c *PaymentChannel) TableName() string {
	return "payment_channels"
}

func (p *PaymentChannel) Validate() error {
	if p.Name == "" {
		return errors.New("name is required")
	}
	if p.Channel == "" {
		return errors.New("channel is required")
	}
	if p.Vendor == "" {
		return errors.New("vendor is required")
	}

	return nil
}
