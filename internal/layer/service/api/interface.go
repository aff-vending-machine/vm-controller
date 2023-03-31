package api

import (
	"context"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/entity"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/lugentpay"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/mail"
)

type Mail interface {
	Send(ctx context.Context, data *mail.Message) error
}

type LugentPay interface {
	AliPay(ctx context.Context, channel *entity.PaymentChannel, req *lugentpay.QRCodeGenerateRequest) (*lugentpay.QRCodeGenerateResponse, error)
	ThaiQR(ctx context.Context, channel *entity.PaymentChannel, req *lugentpay.QRCodeGenerateRequest) (*lugentpay.QRCodeGenerateResponse, error)
	WechatPay(ctx context.Context, channel *entity.PaymentChannel, req *lugentpay.QRCodeGenerateRequest) (*lugentpay.QRCodeGenerateResponse, error)
	Inquiry(ctx context.Context, channel *entity.PaymentChannel, req *lugentpay.InquiryBody) (*lugentpay.InquiryResult, error)
}
