package fiber

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/transport/http"
	"github.com/gofiber/fiber/v2"
)

func routeSlot(api fiber.Router, endpoint http.Slot) {
	api.Get("slots", endpoint.Read)
	api.Get("slots/:id", endpoint.ReadOne)
	api.Post("slots/:id/set-stock", endpoint.SetStock)
	api.Patch("slots/:id", endpoint.Update)
}
