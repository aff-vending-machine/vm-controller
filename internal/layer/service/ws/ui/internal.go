package ui_ws

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/hardware"
	"github.com/gorilla/websocket"
)

type PayloadModel struct {
	Code  int         `json:"code"`
	Stage string      `json:"stage"`
	Data  interface{} `json:"data"`
	Error *string     `json:"error,omitempty"`
}

type Total struct {
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Received int     `json:"received"`
}

func calculateTotal(cart []hardware.Item) Total {
	totalQuantity := 0
	totalReceived := 0
	totalPrice := 0.0

	for _, item := range cart {
		totalQuantity += item.Quantity
		totalReceived += item.Received
		totalPrice += float64(item.Quantity) * item.Price
	}

	return Total{
		Price:    totalPrice,
		Quantity: totalQuantity,
		Received: totalReceived,
	}
}

func checkConnection(c *websocket.Conn) error {
	if c == nil {
		return websocket.ErrCloseSent
	}

	return nil
}
