package http

import "github.com/gofiber/fiber/v2"

type Slot interface {
	Read(*fiber.Ctx) error // GET {slots}
	// Count(*fiber.Ctx) error    // GET {slots/count}
	ReadOne(*fiber.Ctx) error // GET {slots/:id}
	// Create(*fiber.Ctx) error   // POST {slots}
	Update(*fiber.Ctx) error // PATCH {slots/:id}
	// Delete(*fiber.Ctx) error   // DELETE {slots/:id}
	SetStock(*fiber.Ctx) error // POST {slots/:id/set-stock}
}
