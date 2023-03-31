package ws

import (
	"context"
	"time"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/entity"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/hardware"
	"github.com/gorilla/websocket"
)

type UI interface {
	SetConnection(context.Context, *websocket.Conn)
	SendSlots(context.Context, []entity.Slot) error
	SendCart(context.Context, []hardware.Item) error
	SendPaymentChannel(context.Context, []entity.PaymentChannel) error
	SendQRCode(ctx context.Context, orderID string, qrcode string, qty int, price float64) error
	SendPaid(ctx context.Context, orderID string, qty int, price float64) error
	SendIdentified(ctx context.Context, orderID string, qty int, price float64) error
	SendMailRequest(ctx context.Context, orderID string) error
	SendMailIsUsed(ctx context.Context, orderID string, mail string) error
	SendOTPRequest(ctx context.Context, orderID string, mail string, reference string, timestamp time.Time) error
	SendReceivedItem(ctx context.Context, orderID string, cart []hardware.Item, item hardware.Item) error
	SendDone(ctx context.Context, orderID string, cart []hardware.Item) error
	SendToReceive(ctx context.Context, orderID string, cart []hardware.Item) error
	SendToIdle(ctx context.Context) error
	SendError(ctx context.Context, stage string, message string) error
	SendGrabItem(ctx context.Context, stage string, message string) error
	SendEmergency(ctx context.Context, err error) error
}
