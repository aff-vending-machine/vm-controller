package slot_http

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/slot/request"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/fiber/rest"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/trace"
	"github.com/gofiber/fiber/v2"
)

func (t *restImpl) Update(c *fiber.Ctx) error {
	ctx, span := trace.Start(c.Context())
	defer span.End()

	req, err := makeUpdateRequest(c)
	if err != nil {
		trace.RecordError(span, err)
		return rest.BadRequest(c, err)
	}

	// usecase execution
	err = t.usecase.Update(ctx, req)
	if err != nil {
		trace.RecordError(span, err)
		return rest.UsecaseError(c, err)
	}

	return rest.NoContent(c)
}

func makeUpdateRequest(c *fiber.Ctx) (*request.Update, error) {
	var req request.Update
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
