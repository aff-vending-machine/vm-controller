package slot_http

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/module/fiber/http"
	"github.com/aff-vending-machine/vm-controller/pkg/trace"
	"github.com/gofiber/fiber/v2"
)

func (t *httpImpl) Read(c *fiber.Ctx) error {
	ctx, span := trace.Start(c.Context())
	defer span.End()

	// usecase execution
	res, err := t.usecase.Get(ctx, []string{})
	if err != nil {
		trace.RecordError(span, err)
		return http.UsecaseError(c, err)
	}

	return http.OK(c, res)
}
