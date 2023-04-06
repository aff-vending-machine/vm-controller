package fiber

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/module/fiber"
)

type server struct {
	*fiber.Wrapper
}

func New(client *fiber.Wrapper) *server {
	return &server{
		client,
	}
}
