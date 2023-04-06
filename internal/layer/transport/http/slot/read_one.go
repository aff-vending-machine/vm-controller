package slot_http

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/module/fiber/http"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/slot/request"
	"github.com/aff-vending-machine/vm-controller/pkg/trace"
	"github.com/gofiber/fiber/v2"
)

func (t *httpImpl) ReadOne(c *fiber.Ctx) error {
	ctx, span := trace.Start(c.Context())
	defer span.End()

	req, err := makeGetRequest(c)
	if err != nil {
		trace.RecordError(span, err)
		return http.BadRequest(c, err)
	}

	// usecase execution
	res, err := t.usecase.GetOne(ctx, req.ID)
	if err != nil {
		trace.RecordError(span, err)
		return http.UsecaseError(c, err)
	}

	return http.OK(c, res)
}

func makeGetRequest(c *fiber.Ctx) (*request.Get, error) {
	req := request.Get{}
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return nil, err
	}
	req.ID = uint(id)

	return &req, nil
}
