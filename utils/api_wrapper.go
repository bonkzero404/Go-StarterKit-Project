package utils

import (
	"go-boilerplate-clean-arch/domain/models"

	"github.com/gofiber/fiber/v2"
)

func ApiWrapper(ctx *fiber.Ctx, message string, code int, status string, data interface{}) error {
	meta := models.Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	responseJson := models.Response{
		Meta: meta,
		Data: data,
	}

	return ctx.Status(code).JSON(responseJson)
}

func ApiErrorValidation(ctx *fiber.Ctx, message string, data interface{}) error {
	return ApiWrapper(ctx, message, 422, "error_validation", data)
}

func ApiUnprocessableEntity(ctx *fiber.Ctx, message string, data interface{}) error {
	return ApiWrapper(ctx, message, 422, "error_unprocessable_entity", data)
}

func ApiCreated(ctx *fiber.Ctx, message string, data interface{}) error {
	return ApiWrapper(ctx, message, 201, "success_created", data)
}

func ApiOk(ctx *fiber.Ctx, message string, data interface{}) error {
	return ApiWrapper(ctx, message, 200, "success_ok", data)
}
