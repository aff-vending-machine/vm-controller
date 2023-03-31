package rest

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// 200 - OK
func OK(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "done",
		"data":   data,
	})
}

// 204 - NoContent
func NoContent(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":   fiber.StatusNoContent,
		"status": "done",
	})
}

// 201 - Created
func Created(ctx *fiber.Ctx, id string) error {
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":   fiber.StatusCreated,
		"id":     id,
		"status": "done",
	})
}

func UsecaseError(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"code":    fiber.StatusInternalServerError,
		"status":  "error",
		"message": fmt.Sprintf("unexpected: (%s)", err.Error()),
	})
}

// 400 - Bad Request
func BadRequest(ctx *fiber.Ctx, cause error) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"code":    fiber.StatusBadRequest,
		"status":  "error",
		"message": cause.Error(),
	})
}

// 401 - Unauthorized
func Unauthorized(ctx *fiber.Ctx, cause error) error {
	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"code":    fiber.StatusUnauthorized,
		"status":  "error",
		"message": cause.Error(),
	})
}

// 403 - Forbidden
func Forbidden(ctx *fiber.Ctx, cause error) error {
	return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
		"code":    fiber.StatusForbidden,
		"status":  "error",
		"message": cause.Error(),
	})
}
