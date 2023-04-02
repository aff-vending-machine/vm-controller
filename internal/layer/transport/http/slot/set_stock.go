package slot_http

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/slot/request"
	"github.com/aff-vending-machine/vm-controller/pkg/module/fiber/rest"
	"github.com/aff-vending-machine/vm-controller/pkg/trace"
	"github.com/gofiber/fiber/v2"
)

func (t *restImpl) SetStock(c *fiber.Ctx) error {
	ctx, span := trace.Start(c.Context())
	defer span.End()

	req, err := makeSetStockRequest(c)
	if err != nil {
		trace.RecordError(span, err)
		return rest.BadRequest(c, err)
	}

	// usecase execution
	err = t.usecase.SetStock(ctx, req)
	if err != nil {
		trace.RecordError(span, err)
		return rest.UsecaseError(c, err)
	}

	return rest.NoContent(c)
}

func makeSetStockRequest(c *fiber.Ctx) (*request.SetStock, error) {
	var req request.SetStock
	if err := c.BodyParser(&req); err != nil {
		return nil, err
	}
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return nil, err
	}
	req.ID = uint(id)

	return &req, nil
}
