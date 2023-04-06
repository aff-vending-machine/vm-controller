package rpc

import (
	"context"
	"encoding/json"

	"github.com/rabbitmq/amqp091-go"
)

type Ctx struct {
	UserContext context.Context
	context.CancelFunc
	amqp091.Delivery
	amqp091.Publishing
}

type Handler func(*Ctx) error

func NewContext(delivery amqp091.Delivery) *Ctx {
	ctx, cancel := context.WithCancel(context.Background())
	return &Ctx{
		UserContext: ctx,
		CancelFunc:  cancel,
		Delivery:    delivery,
		Publishing: amqp091.Publishing{
			CorrelationId: delivery.CorrelationId,
		},
	}
}

func (c *Ctx) Ok(data interface{}) error {
	res, _ := json.Marshal(map[string]interface{}{
		"code":    200,
		"data":    data,
		"message": "ok",
	})

	c.Publishing.Body = res
	return nil
}

func (c *Ctx) BadRequest(err error) error {
	res, _ := json.Marshal(map[string]interface{}{
		"code":    400,
		"message": "bad request",
		"error":   err.Error(),
	})

	c.Publishing.Body = res
	return nil
}

func (c *Ctx) InternalServer(err error) error {
	res, _ := json.Marshal(map[string]interface{}{
		"code":    500,
		"message": "internal server error",
		"error":   err.Error(),
	})

	c.Publishing.Body = res
	return nil
}
