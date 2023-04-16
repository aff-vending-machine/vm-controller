package websocket

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
	"github.com/gorilla/websocket"
)

type Frontend interface {
	SetConnection(context.Context, *websocket.Conn)
	SendSlots(context.Context, []entity.Slot) error
	SendCart(context.Context, []hardware.Item) error
	SendPaymentChannel(context.Context, []entity.PaymentChannel) error
	SendQRCode(ctx context.Context, orderID string, qrcode string, qty int, price float64) error
	SendPaid(ctx context.Context, orderID string, qty int, price float64) error
	SendReceivedItem(ctx context.Context, orderID string, cart []hardware.Item, item hardware.Item) error
	SendDone(ctx context.Context, orderID string, cart []hardware.Item) error
	SendToReceive(ctx context.Context, orderID string, cart []hardware.Item) error
	SendToIdle(ctx context.Context) error
	SendError(ctx context.Context, stage string, message string) error
	SendGrabItem(ctx context.Context, stage string, message string) error
	SendEmergency(ctx context.Context, err error) error
}
