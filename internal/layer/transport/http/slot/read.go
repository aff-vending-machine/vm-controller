package slot_http

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/fiber/rest"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/trace"
	"github.com/gofiber/fiber/v2"
)

func (t *restImpl) Read(c *fiber.Ctx) error {
	ctx, span := trace.Start(c.Context())
	defer span.End()

	// usecase execution
	res, err := t.usecase.Get(ctx, []string{})
	if err != nil {
		trace.RecordError(span, err)
		return rest.UsecaseError(c, err)
	}

	return rest.OK(c, res)
}
