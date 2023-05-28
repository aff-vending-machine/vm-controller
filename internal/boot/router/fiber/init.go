package fiber

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/infra/fiber"
)

type server struct {
	*fiber.Wrapper
}

func New(app *fiber.Wrapper) *server {
	return &server{
		app,
	}
}
